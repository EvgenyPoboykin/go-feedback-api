package controllers

import (
	"net/http"

	"github.com/eugenepoboykin/go-feedback-api/constant"
	"github.com/eugenepoboykin/go-feedback-api/helpers"
	"github.com/eugenepoboykin/go-feedback-api/schema"
	"github.com/eugenepoboykin/go-feedback-api/services"
)

func Create(w http.ResponseWriter, r *http.Request) {

	var body schema.IssueCreateArgs

	role := r.Context().Value("oauth.role").(string)
	fullName := r.Context().Value("oauth.fullName").(string)
	clientId := r.Context().Value("oauth.clientId").(string)

	if role != constant.Employee_server {
		helpers.ErrorResponse(w, http.StatusConflict, constant.NO_CREDENTIAL, constant.ResponseMessage_AccessDenied)

		helpers.Log.ErrorLog.Println(role + " | " + fullName + " | *********")

		return
	}

	errorMessage := helpers.CreateBodyParse(w, r, &body)

	if errorMessage != nil {
		helpers.ErrorResponse(w, errorMessage.Status, errorMessage.Type, errorMessage.Description)

		return
	}

	isseus, err := services.CreateIsseu(
		body.Uri,
		body.Image64,
		body.Description,
		clientId,
		fullName,
	)

	if err != nil {
		helpers.ErrorResponse(w, http.StatusBadGateway, constant.SERVICE_CREATE_ISSEU, constant.ResponseMessage_NotCreateIsseu)

		return
	}

	helpers.Response[schema.Issue](w, isseus)
}
