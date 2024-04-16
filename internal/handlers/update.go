package handlers

import (
	"context"
	"net/http"

	"github.com/eugenepoboykin/go-feedback-api/internal/lib/ctx"
	"github.com/eugenepoboykin/go-feedback-api/internal/lib/response"
	"github.com/eugenepoboykin/go-feedback-api/internal/models"
	"github.com/eugenepoboykin/go-feedback-api/internal/validator"

	"github.com/go-chi/chi/v5"
)

type UpdateItemHandler interface {
	UpdateIssue(ctx context.Context, params models.UpdateIssueArgs) (*models.Issue, error)
	GetOptionByValue(ctx context.Context, value string) (*models.Option, error)
}

func Update(db UpdateItemHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		role := r.Context().Value("oauth.role").(string)
		if role == Employee {
			response.ErrorResponse(w, http.StatusConflict, NO_CREDENTIAL, ResponseMessage_AccessDenied)

			return
		}

		id := chi.URLParam(r, "id")

		validation := validator.NewValidtor(r.Body)
		body, bodyError := validation.CheckUpdateAgrs(id)
		if bodyError != nil {
			response.ErrorResponse(w, http.StatusServiceUnavailable, SERVICE_PARSE, ResponseMessage_ServiceBodyParser)

			return
		}

		c := ctx.Ctx()

		if string(body.Status) != `` {
			_, errorStatus := db.GetOptionByValue(c, body.Status)
			if errorStatus != nil {
				response.ErrorResponse(w, http.StatusServiceUnavailable, SERVICE_VAliDATE, ResponseMessage_UpdateStatusError)

				return
			}
		}

		isseu, err := db.UpdateIssue(c, *body)
		if err != nil {
			response.ErrorResponse(w, http.StatusBadGateway, SERVICE_CREATE_ISSEU, ResponseMessage_ServiceUpdateError)

			return
		}

		response.Response(w, isseu)
	}
}
