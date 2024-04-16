package handlers

import (
	"context"
	"net/http"

	"github.com/eugenepoboykin/go-feedback-api/internal/lib/ctx"
	"github.com/eugenepoboykin/go-feedback-api/internal/lib/response"
	"github.com/eugenepoboykin/go-feedback-api/internal/models"
)

type OptionsHandler interface {
	GetOptions(ctx context.Context) (*[]models.Option, error)
}

func Options(db OptionsHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		role := r.Context().Value("oauth.role").(string)
		if role != Employee && role != Admin {
			response.ErrorResponse(w, http.StatusConflict, NO_CREDENTIAL, ResponseMessage_AccessDenied)

			return
		}

		c := ctx.Ctx()

		options, err := db.GetOptions(c)
		if err != nil {
			response.ErrorResponse(w, http.StatusServiceUnavailable, SERVICE_RETURN, ResponseMessage_OptionsError)

			return
		}

		response.Response(w, options)

	}
}
