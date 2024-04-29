package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/eugenepoboykin/go-feedback-api/internal/lib/response"
	"github.com/eugenepoboykin/go-feedback-api/internal/validator"
)

func (as ApiSettings) Create(w http.ResponseWriter, r *http.Request) {
	c, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	role := r.Context().Value("oauth.role").(string)
	if role == Admin {
		response.ErrorResponse(w, http.StatusConflict, NO_CREDENTIAL, ResponseMessage_AccessDenied)

		return
	}

	validatoion := validator.NewValidtor(r.Body)
	body, validationError := validatoion.CheckAddAgrs()
	if validationError != nil {
		response.ErrorResponse(w, validationError.Status, validationError.Type, validationError.Description)

		return
	}

	res, err := as.conn.AddIssueItem(c, *body)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadGateway, SERVICE_CREATE_ISSEU, ResponseMessage_NotCreateIsseu)

		return
	}

	response.Response(w, res)

}
