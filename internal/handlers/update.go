package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/eugenepoboykin/go-feedback-api/internal/lib/response"
	"github.com/eugenepoboykin/go-feedback-api/internal/validator"

	"github.com/go-chi/chi/v5"
)

func (as ApiSettings) Update(w http.ResponseWriter, r *http.Request) {
	c, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

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

	if string(body.Status) != `` {
		_, errorStatus := as.conn.GetOptionByValue(c, body.Status)
		if errorStatus != nil {
			response.ErrorResponse(w, http.StatusServiceUnavailable, SERVICE_VAliDATE, ResponseMessage_UpdateStatusError)

			return
		}
	}

	isseu, err := as.conn.UpdateIssue(c, *body)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadGateway, SERVICE_CREATE_ISSEU, ResponseMessage_ServiceUpdateError)

		return
	}

	response.Response(w, isseu)
}
