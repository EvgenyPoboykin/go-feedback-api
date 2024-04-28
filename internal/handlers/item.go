package handlers

import (
	"net/http"

	"github.com/eugenepoboykin/go-feedback-api/internal/lib/ctx"
	"github.com/eugenepoboykin/go-feedback-api/internal/lib/response"

	"github.com/go-chi/chi/v5"
)

func (as ApiSettings) Item(w http.ResponseWriter, r *http.Request) {

	role := r.Context().Value("oauth.role").(string)
	if role != Employee && role != Admin {
		response.ErrorResponse(w, http.StatusConflict, NO_CREDENTIAL, ResponseMessage_AccessDenied)

		return
	}

	id := chi.URLParam(r, "id")

	c := ctx.Ctx()
	isseu, err := as.conn.GetById(c, id)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, BAD_REQUEST, ResponseMessage_NotFoundIsseu+id)

		return
	}

	response.Response(w, isseu)
}
