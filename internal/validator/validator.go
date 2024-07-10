package validator

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/eugenepoboykin/go-feedback-api/internal/domain/models"

	"github.com/eugenepoboykin/go-feedback-api/internal/model"
)

type Validator struct {
	Body io.ReadCloser
}

func NewValidator(body io.ReadCloser) *Validator {
	return &Validator{
		Body: body,
	}
}

func (v *Validator) CheckAddArgs() (*models.StorageAddIssueDTO, *model.ErrorMessage) {
	var body models.StorageAddIssueDTO
	err := json.NewDecoder(v.Body).Decode(&body)
	if err != nil {
		return nil, &model.ErrorMessage{
			Status:      http.StatusServiceUnavailable,
			Type:        SERVICE_PARSE,
			Description: ResponseMessage_ServiceBodyParser,
		}
	}

	if body.Description == "" {
		return nil, &model.ErrorMessage{
			Status:      http.StatusBadRequest,
			Type:        FIELD_DESCRIPTION,
			Description: ResponseMessage_DescriptionError,
		}
	}

	if body.Uri == "" {
		return nil, &model.ErrorMessage{
			Status:      http.StatusBadRequest,
			Type:        FIELD_URI,
			Description: ResponseMessage_UriError,
		}
	}

	return &body, nil
}

func (v *Validator) CheckUpdateArgs(issueId string) (*models.StorageUpdateIssueDTO, *model.ErrorMessage) {
	var body models.StorageUpdateIssueDTO
	err := json.NewDecoder(v.Body).Decode(&body)
	if err != nil {
		return nil, &model.ErrorMessage{
			Status:      http.StatusServiceUnavailable,
			Type:        SERVICE_PARSE,
			Description: ResponseMessage_ServiceBodyParser,
		}
	}

	if issueId == "" {
		return nil, &model.ErrorMessage{
			Status:      http.StatusServiceUnavailable,
			Type:        SERVICE_PARSE,
			Description: ResponseMessage_ParamsIdError,
		}
	}

	if body.Status == nil {
		return nil, &model.ErrorMessage{
			Status:      http.StatusServiceUnavailable,
			Type:        SERVICE_PARSE,
			Description: ResponseMessage_StatusError,
		}
	}

	body.Id = issueId

	return &body, nil
}

func (v *Validator) CheckListArgs() (*models.StorageListDTO, *model.ErrorMessage) {
	var body models.StorageListDTO
	err := json.NewDecoder(v.Body).Decode(&body)
	if err != nil {
		return nil, &model.ErrorMessage{
			Status:      http.StatusServiceUnavailable,
			Type:        SERVICE_PARSE,
			Description: ResponseMessage_ServiceBodyParser,
		}
	}

	if body.Page == 0 {
		return nil, &model.ErrorMessage{
			Status:      http.StatusServiceUnavailable,
			Type:        SERVICE_PARSE,
			Description: ResponseMessage_StatusError,
		}
	}

	if body.PageSize == 0 {
		return nil, &model.ErrorMessage{
			Status:      http.StatusServiceUnavailable,
			Type:        SERVICE_PARSE,
			Description: ResponseMessage_StatusError,
		}
	}

	return &body, nil
}
