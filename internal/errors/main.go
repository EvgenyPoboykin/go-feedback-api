package errors

type ErrorPayloadResponse struct {
	Description string `json:"description"`
}

type ErrorResponse struct {
	Type    string               `json:"type"`
	Payload ErrorPayloadResponse `json:"payload"`
}

func Error(statusType string, description string) *ErrorResponse {
	return &ErrorResponse{
		Type: statusType,
		Payload: ErrorPayloadResponse{
			Description: description,
		},
	}
}
