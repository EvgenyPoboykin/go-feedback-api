package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/eugenepoboykin/go-feedback-api/constant"
	"github.com/eugenepoboykin/go-feedback-api/schema"
)

func Response[T interface{}](w http.ResponseWriter, data *T) {

	out, errMarshall := json.Marshal(&data)

	if errMarshall != nil {
		ErrorResponse(w, http.StatusInsufficientStorage, constant.SERVICE_VAliDATE, constant.ResponseMessage_ServerReturnError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(out)

}

func ErrorResponse(w http.ResponseWriter, status int, statusType string, description string) {
	payload, err := json.Marshal(schema.ErrorResponse{
		Type: statusType,
		Payload: schema.ErrorPayloadResponse{
			Description: description,
		},
	})

	if err != nil {
		Log.ErrorLog.Println(constant.ResponseMessage_ServerReturnError, err)
	}

	w.WriteHeader(status)
	w.Write(payload)

}
