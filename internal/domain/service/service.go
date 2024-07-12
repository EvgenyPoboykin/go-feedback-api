package service

import (
	"context"
	globalErrors "errors"
	"io"

	"github.com/eugenepoboykin/go-feedback-api/internal/converter"
	"github.com/eugenepoboykin/go-feedback-api/internal/domain/env"
	"github.com/eugenepoboykin/go-feedback-api/internal/domain/models"
	"github.com/eugenepoboykin/go-feedback-api/internal/errors"
	"github.com/eugenepoboykin/go-feedback-api/internal/lib/pagination"
	"github.com/eugenepoboykin/go-feedback-api/internal/validator"
)

type IssueStorage interface {
	ListAdmin(ctx context.Context) (*[]models.StorageIssue, error)
	ListEmployee(ctx context.Context, clientId *string) (*[]models.StorageIssue, error)
	IssueById(ctx context.Context, issueId string) (*models.StorageIssue, error)
	Create(ctx context.Context, params models.StorageAddIssueDTO) (*models.StorageIssue, error)
	Update(ctx context.Context, params models.StorageUpdateIssueDTO) (*models.StorageIssue, error)
	IssueByParams(ctx context.Context, params models.StorageAddIssueDTO) (*models.StorageIssue, error)
	Options(ctx context.Context) (*[]models.StorageOption, error)
	OptionByValue(ctx context.Context, value string) (*models.StorageOption, error)
}

type Service struct {
	storage IssueStorage
}

func NewService(issueRepository IssueStorage) (*Service, error) {
	if issueRepository == nil {
		return nil, globalErrors.New("no db connection")
	}

	return &Service{
		storage: issueRepository,
	}, nil
}

func (s Service) ListAdmin(ctx context.Context, role string, body io.ReadCloser) (*models.ServiceIssuesList, *errors.ErrorResponse) {
	if role == env.Environment.EmployeeRole || role == "" {
		return nil, errors.Error(NoCredential, ResponseMessage_AccessDenied)
	}

	validation := validator.NewValidator(body)
	params, validationError := validation.CheckListArgs()
	if validationError != nil {
		return nil, errors.Error(validationError.Type, validationError.Description)
	}

	if params.Status != nil {
		if _, errorStatus := s.storage.OptionByValue(ctx, *params.Status); errorStatus != nil {
			return nil, errors.Error(ServiceValidate, ResponseMessage_UpdateStatusError)
		}
	}

	var page models.ServiceIssuesList
	page.Page = params.Page
	page.PageSize = params.PageSize
	page.Status = params.Status

	issues, err := s.storage.ListAdmin(ctx)
	if err != nil {
		return nil, errors.Error(ServiceReturn, ResponseMessage_ListError)
	}

	result, err := converter.FromDbToServiceMap[[]models.StorageIssue, []models.ServiceIssue](*issues)
	if err != nil {
		return nil, errors.Error(ServiceReturn, ResponseMessage_ListError)
	}

	issuesSlice := pagination.IssuePerPage(*result, converter.ToListArgsFromService(*params))
	page.TotalCount = len(*issues)
	page.Issues = issuesSlice

	return &page, nil
}

func (s Service) ListEmployee(ctx context.Context, role string, clientId string, body io.ReadCloser) (*models.ServiceIssuesList, *errors.ErrorResponse) {
	if clientId == "" {
		return nil, errors.Error(ServiceReturn, ResponseMessage_ListError)
	}

	if role == env.Environment.AdminRole || role == "" {
		return nil, errors.Error(NoCredential, ResponseMessage_AccessDenied)
	}

	validation := validator.NewValidator(body)
	params, validationError := validation.CheckListArgs()
	if validationError != nil {
		return nil, errors.Error(validationError.Type, validationError.Description)
	}

	if params.Status != nil {
		if _, errorStatus := s.storage.OptionByValue(ctx, *params.Status); errorStatus != nil {
			return nil, errors.Error(ServiceValidate, ResponseMessage_UpdateStatusError)
		}
	}

	var page models.ServiceIssuesList
	page.Page = params.Page
	page.PageSize = params.PageSize
	page.Status = params.Status

	issues, err := s.storage.ListEmployee(ctx, &clientId)
	if err != nil {
		return nil, errors.Error(ServiceReturn, ResponseMessage_ListError)
	}

	result, err := converter.FromDbToServiceMap[[]models.StorageIssue, []models.ServiceIssue](*issues)
	if err != nil {
		return nil, errors.Error(ServiceReturn, ResponseMessage_ListError)
	}

	issuesSlice := pagination.IssuePerPage(*result, converter.ToListArgsFromService(*params))
	page.TotalCount = len(*issues)
	page.Issues = issuesSlice

	return &page, nil
}

func (s Service) IssueById(ctx context.Context, role string, id string) (*models.ServiceIssue, *errors.ErrorResponse) {
	if role != env.Environment.EmployeeRole && role != env.Environment.AdminRole {
		return nil, errors.Error(NoCredential, ResponseMessage_AccessDenied)
	}

	res, err := s.storage.IssueById(ctx, id)
	if err != nil {
		return nil, errors.Error(BadRequest, ResponseMessage_NotFoundIssue+id)
	}

	issue, err := converter.FromDbToServiceMap[models.StorageIssue, models.ServiceIssue](*res)
	if err != nil {
		return nil, errors.Error(BadRequest, ResponseMessage_NotFoundIssue+id)
	}

	return issue, nil
}

func (s Service) Create(ctx context.Context, role string, body io.ReadCloser) (*models.ServiceIssue, *errors.ErrorResponse) {
	if role == env.Environment.AdminRole {
		return nil, errors.Error(NoCredential, ResponseMessage_AccessDenied)
	}

	validation := validator.NewValidator(body)
	params, validationError := validation.CheckAddArgs()
	if validationError != nil {
		return nil, errors.Error(validationError.Type, validationError.Description)
	}

	res, err := s.storage.Create(ctx, *params)
	if err != nil {
		return nil, errors.Error(ServiceCreateIssue, ResponseMessage_NotCreateIssue)
	}

	issue, err := converter.FromDbToServiceMap[models.StorageIssue, models.ServiceIssue](*res)
	if err != nil {
		return nil, errors.Error(ServiceCreateIssue, ResponseMessage_NotCreateIssue)
	}

	return issue, nil
}

func (s Service) Update(ctx context.Context, role string, id string, body io.ReadCloser) (*models.ServiceIssue, *errors.ErrorResponse) {
	if role == env.Environment.EmployeeRole {
		return nil, errors.Error(NoCredential, ResponseMessage_AccessDenied)
	}

	validation := validator.NewValidator(body)
	params, bodyError := validation.CheckUpdateArgs(id)
	if bodyError != nil {
		return nil, errors.Error(bodyError.Type, bodyError.Description)
	}

	if params.Status != nil {
		if _, errorStatus := s.storage.OptionByValue(ctx, *params.Status); errorStatus != nil {
			return nil, errors.Error(ServiceValidate, ResponseMessage_UpdateStatusError)
		}
	}

	issue, err := s.storage.Update(ctx, *params)
	if err != nil {
		return nil, errors.Error(ServiceCreateIssue, ResponseMessage_ServiceUpdateError)
	}

	result, err := converter.FromDbToServiceMap[models.StorageIssue, models.ServiceIssue](*issue)
	if err != nil {
		return nil, errors.Error(ServiceReturn, ResponseMessage_ListError)
	}

	return result, nil
}

func (s Service) Options(ctx context.Context, role string) (*[]models.ServiceOption, *errors.ErrorResponse) {
	if role != env.Environment.EmployeeRole && role != env.Environment.AdminRole {
		return nil, errors.Error(NoCredential, ResponseMessage_AccessDenied)
	}

	options, err := s.storage.Options(ctx)
	if err != nil {
		return nil, errors.Error(ServiceReturn, ResponseMessage_OptionsError)
	}

	result, err := converter.FromDbToServiceMap[[]models.StorageOption, []models.ServiceOption](*options)
	if err != nil {
		return nil, errors.Error(ServiceReturn, ResponseMessage_ListError)
	}

	return result, nil
}
