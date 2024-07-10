package response

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/eugenepoboykin/go-feedback-api/internal/errors"
)

const (
	SERVICE_VAliDATE                  = "SERVICE_VAliDATE"
	ResponseMessage_ServerReturnError = "Something went wrong!"
)

var errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

func Response[T interface{}](w http.ResponseWriter, data *T) {

	out, errMarshall := json.Marshal(&data)

	if errMarshall != nil {
		ErrorResponse(w, http.StatusInsufficientStorage, *errors.Error(SERVICE_VAliDATE, ResponseMessage_ServerReturnError))

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(out)

}

func ErrorResponse(w http.ResponseWriter, status int, e errors.ErrorResponse) {
	payload, err := json.Marshal(e)

	if err != nil {
		errorLog.Println(ResponseMessage_ServerReturnError, err)
	}

	w.WriteHeader(status)
	w.Write(payload)

}
