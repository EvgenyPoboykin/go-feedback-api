package handlers

import (
	"context"
	"net/http"

	"github.com/eugenepoboykin/go-feedback-api/internal/lib/ctx"
	"github.com/eugenepoboykin/go-feedback-api/internal/lib/response"
	"github.com/eugenepoboykin/go-feedback-api/internal/models"
	"github.com/eugenepoboykin/go-feedback-api/internal/validator"
)

type CreateItemHandler interface {
	AddIssueItem(ctx context.Context, params models.AddIssueArgs) (*models.Issue, error)
}

func Create(db CreateItemHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

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

		c := ctx.Ctx()
		res, err := db.AddIssueItem(c, *body)
		if err != nil {
			response.ErrorResponse(w, http.StatusBadGateway, SERVICE_CREATE_ISSEU, ResponseMessage_NotCreateIsseu)

			return
		}

		response.Response(w, res)

	}
}