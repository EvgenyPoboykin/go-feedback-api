package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/eugenepoboykin/go-feedback-api/constant"
	"github.com/eugenepoboykin/go-feedback-api/helpers"
	"github.com/eugenepoboykin/go-feedback-api/schema"
	"github.com/eugenepoboykin/go-feedback-api/services"

	"github.com/go-chi/chi/v5"
)

func Update(w http.ResponseWriter, r *http.Request) {
	role := r.Context().Value("oauth.role").(string)

	if role == constant.Employee_server {
		helpers.ErrorResponse(w, http.StatusConflict, constant.NO_CREDENTIAL, constant.ResponseMessage_AccessDenied)

		return
	}

	var body schema.IssueUpdateArgs

	id := chi.URLParam(r, "id")

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		helpers.ErrorResponse(w, http.StatusServiceUnavailable, constant.SERVICE_PARSE, constant.ResponseMessage_ServiceBodyParser)

		return
	}

	if string(body.Status) != `` {
		_, errorStatus := services.GetOptionByValue(body.Status)

		if errorStatus != nil {
			helpers.ErrorResponse(w, http.StatusServiceUnavailable, constant.SERVICE_VAliDATE, constant.ResponseMessage_UpdateStatusError)

			return
		}
	}

	isseu, err := services.UpdateIsseu(
		id,
		body.Comment,
		body.Status,
	)

	if err != nil {
		helpers.ErrorResponse(w, http.StatusBadGateway, constant.SERVICE_CREATE_ISSEU, constant.ResponseMessage_ServiceUpdateError)

		return
	}

	helpers.Response[schema.Issue](w, isseu)
}
