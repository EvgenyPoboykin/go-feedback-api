package models

type ErrorPayloadResponse struct {
	Description string `json:"description"`
}

type ErrorResponse struct {
	Type    string               `json:"type"`
	Payload ErrorPayloadResponse `json:"payload"`
}

type ErrorMessage struct {
	Status      int
	Type        string
	Description string
}
