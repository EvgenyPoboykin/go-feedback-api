package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/eugenepoboykin/go-feedback-api/internal/lib/response"
)

func (as ApiSettings) Options(w http.ResponseWriter, r *http.Request) {

	c, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	role := r.Context().Value("oauth.role").(string)
	if role != Employee && role != Admin {
		response.ErrorResponse(w, http.StatusConflict, NO_CREDENTIAL, ResponseMessage_AccessDenied)

		return
	}

	options, err := as.conn.GetOptions(c)

	fmt.Print(options, err)
	if err != nil {
		response.ErrorResponse(w, http.StatusServiceUnavailable, SERVICE_RETURN, ResponseMessage_OptionsError)

		return
	}

	response.Response(w, options)

}
