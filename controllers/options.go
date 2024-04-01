package controllers

import (
	"net/http"

	"github.com/eugenepoboykin/go-feedback-api/constant"
	"github.com/eugenepoboykin/go-feedback-api/helpers"
	"github.com/eugenepoboykin/go-feedback-api/schema"
	"github.com/eugenepoboykin/go-feedback-api/services"
)

func Options(w http.ResponseWriter, r *http.Request) {

	options, err := services.GetOptions()

	if err != nil {
		helpers.ErrorResponse(w, http.StatusServiceUnavailable, constant.SERVICE_RETURN, constant.ResponseMessage_OptionsError)

		return
	}

	helpers.Response[[]schema.IssuesStatus](w, &options)

}
