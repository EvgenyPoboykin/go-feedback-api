package handlers

import (
	"context"
	"net/http"

	"github.com/eugenepoboykin/go-feedback-api/internal/lib/ctx"
	"github.com/eugenepoboykin/go-feedback-api/internal/lib/pagination"
	"github.com/eugenepoboykin/go-feedback-api/internal/lib/response"
	"github.com/eugenepoboykin/go-feedback-api/internal/models"
	"github.com/eugenepoboykin/go-feedback-api/internal/validator"
)

type ListAdminHandler interface {
	GetListByAdmin(ctx context.Context) (*[]models.Issue, error)
	GetOptionByValue(ctx context.Context, value string) (*models.Option, error)
}

func ListAdmin(db ListAdminHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var empty = make([]models.Issue, 0)

		role := r.Context().Value("oauth.role").(string)
		if role == Employee {
			response.ErrorResponse(w, http.StatusConflict, NO_CREDENTIAL, ResponseMessage_AccessDenied)

			return
		}

		validatoion := validator.NewValidtor(r.Body)
		body, validationError := validatoion.CheckListAgrs()
		if validationError != nil {
			response.ErrorResponse(w, validationError.Status, validationError.Type, validationError.Description)

			return
		}

		var page models.IssuesList

		page.Page = body.Page
		page.PageSize = body.PageSize
		page.Status = body.Status

		c := ctx.Ctx()

		if string(body.Status) != "" {
			_, errorStatus := db.GetOptionByValue(c, string(body.Status))

			if errorStatus != nil {
				response.ErrorResponse(w, http.StatusServiceUnavailable, SERVICE_VAliDATE, ResponseMessage_UpdateStatusError)

				return
			}
		}

		isseus, err := db.GetListByAdmin(c)
		if err != nil {
			response.ErrorResponse(w, http.StatusServiceUnavailable, SERVICE_RETURN, ResponseMessage_ListError)

			return
		}

		if isseus != nil {
			isseusSlice := pagination.IssuePerPage(*isseus, *body)

			page.TotalCount = len(*isseus)
			page.Issues = isseusSlice
		} else {
			page.TotalCount = 0
			page.Issues = empty
		}

		response.Response(w, &page)
	}
}
