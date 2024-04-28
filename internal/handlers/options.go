package handlers

import (
	"net/http"

	"github.com/eugenepoboykin/go-feedback-api/internal/lib/ctx"
	"github.com/eugenepoboykin/go-feedback-api/internal/lib/response"
)

func (as ApiSettings) Options(w http.ResponseWriter, r *http.Request) {

	role := r.Context().Value("oauth.role").(string)
	if role != Employee && role != Admin {
		response.ErrorResponse(w, http.StatusConflict, NO_CREDENTIAL, ResponseMessage_AccessDenied)

		return
	}

	c := ctx.Ctx()

	options, err := as.conn.GetOptions(c)
	if err != nil {
		response.ErrorResponse(w, http.StatusServiceUnavailable, SERVICE_RETURN, ResponseMessage_OptionsError)

		return
	}

	response.Response(w, options)

}
