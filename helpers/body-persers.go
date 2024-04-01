package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/eugenepoboykin/go-feedback-api/constant"
	"github.com/eugenepoboykin/go-feedback-api/schema"
)

func CreateBodyParse(w http.ResponseWriter, r *http.Request, body *schema.IssueCreateArgs) *schema.ErrorMessage {

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		return &schema.ErrorMessage{
			Status:      http.StatusServiceUnavailable,
			Type:        constant.SERVICE_PARSE,
			Description: constant.ResponseMessage_ServiceBodyParser,
		}
	}

	if body.Description == "" {
		return &schema.ErrorMessage{
			Status:      http.StatusBadRequest,
			Type:        constant.FIELD_DESCRIPTION,
			Description: constant.ResponseMessage_DescriptionError,
		}
	}

	if body.Uri == "" {
		return &schema.ErrorMessage{
			Status:      http.StatusBadRequest,
			Type:        constant.FIELD_URI,
			Description: constant.ResponseMessage_UriError,
		}
	}

	return nil
}
