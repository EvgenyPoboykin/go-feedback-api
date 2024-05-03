package validator

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/eugenepoboykin/go-feedback-api/internal/models"
)

type Validator struct {
	Body io.ReadCloser
}

func NewValidator(body io.ReadCloser) *Validator {
	return &Validator{
		Body: body,
	}
}

func (v *Validator) CheckAddArgs() (*models.AddIssueArgs, *models.ErrorMessage) {
	var body models.AddIssueArgs
	err := json.NewDecoder(v.Body).Decode(&body)

	if err != nil {
		return nil, &models.ErrorMessage{
			Status:      http.StatusServiceUnavailable,
			Type:        SERVICE_PARSE,
			Description: ResponseMessage_ServiceBodyParser,
		}
	}

	if body.Description == "" {
		return nil, &models.ErrorMessage{
			Status:      http.StatusBadRequest,
			Type:        FIELD_DESCRIPTION,
			Description: ResponseMessage_DescriptionError,
		}
	}

	if body.Uri == "" {
		return nil, &models.ErrorMessage{
			Status:      http.StatusBadRequest,
			Type:        FIELD_URI,
			Description: ResponseMessage_UriError,
		}
	}

	return &body, nil
}

func (v *Validator) CheckUpdateArgs(issueId string) (*models.UpdateIssueArgs, *models.ErrorMessage) {
	var body models.UpdateIssueArgs
	err := json.NewDecoder(v.Body).Decode(&body)

	if err != nil {
		return nil, &models.ErrorMessage{
			Status:      http.StatusServiceUnavailable,
			Type:        SERVICE_PARSE,
			Description: ResponseMessage_ServiceBodyParser,
		}
	}

	if issueId == "" {
		return nil, &models.ErrorMessage{
			Status:      http.StatusServiceUnavailable,
			Type:        SERVICE_PARSE,
			Description: ResponseMessage_ParamsIdError,
		}
	}

	if body.Status == "" {
		return nil, &models.ErrorMessage{
			Status:      http.StatusServiceUnavailable,
			Type:        SERVICE_PARSE,
			Description: ResponseMessage_StatusError,
		}
	}

	body.Id = issueId

	return &body, nil
}

func (v *Validator) CheckListArgs() (*models.ListArgs, *models.ErrorMessage) {
	var body models.ListArgs
	err := json.NewDecoder(v.Body).Decode(&body)

	if err != nil {
		return nil, &models.ErrorMessage{
			Status:      http.StatusServiceUnavailable,
			Type:        SERVICE_PARSE,
			Description: ResponseMessage_ServiceBodyParser,
		}
	}

	if body.Page == 0 {
		return nil, &models.ErrorMessage{
			Status:      http.StatusServiceUnavailable,
			Type:        SERVICE_PARSE,
			Description: ResponseMessage_StatusError,
		}
	}

	if body.PageSize == 0 {
		return nil, &models.ErrorMessage{
			Status:      http.StatusServiceUnavailable,
			Type:        SERVICE_PARSE,
			Description: ResponseMessage_StatusError,
		}
	}

	return &body, nil
}
