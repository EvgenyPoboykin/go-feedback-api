package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/eugenepoboykin/go-feedback-api/internal/lib/pagination"
	"github.com/eugenepoboykin/go-feedback-api/internal/lib/response"
	"github.com/eugenepoboykin/go-feedback-api/internal/models"
	"github.com/eugenepoboykin/go-feedback-api/internal/validator"
)

func (as ApiSettings) ListEmployee(w http.ResponseWriter, r *http.Request) {
	c, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	var empty = make([]models.Issue, 0)

	role := r.Context().Value("oauth.role").(string)
	if role == Admin || role == "" {
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

	if string(body.Status) != `` {
		_, errorStatus := as.conn.GetOptionByValue(c, string(body.Status))
		if errorStatus != nil {
			response.ErrorResponse(w, http.StatusServiceUnavailable, SERVICE_VAliDATE, ResponseMessage_UpdateStatusError)

			return
		}
	}

	clientId := r.Context().Value("oauth.clientId").(string)

	if clientId == "" {
		response.ErrorResponse(w, http.StatusServiceUnavailable, SERVICE_RETURN, ResponseMessage_ListError)

		return
	}

	isseus, err := as.conn.GetListByEmployee(c, clientId)
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
