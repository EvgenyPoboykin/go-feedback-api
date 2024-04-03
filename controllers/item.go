package controllers

import (
	"net/http"

	"github.com/eugenepoboykin/go-feedback-api/constant"
	"github.com/eugenepoboykin/go-feedback-api/helpers"
	"github.com/eugenepoboykin/go-feedback-api/schema"
	"github.com/eugenepoboykin/go-feedback-api/services"

	"github.com/go-chi/chi/v5"
)

func Item(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	isseu, err := services.GetIsseu(id)
	if err != nil {
		helpers.ErrorResponse(w, http.StatusBadRequest, constant.BAD_REQUEST, constant.ResponseMessage_NotFoundIsseu+id)

		return
	}

	helpers.Response[schema.Issue](w, isseu)
}
