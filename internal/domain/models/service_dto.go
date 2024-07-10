package models

type ServiceAddIssueDTO struct {
	Uri         string `json:"uri"`
	Image64     string `json:"image64"`
	Description string `json:"description"`
	ClientId    string `json:"clientId"`
	ClientName  string `json:"clientName"`
}

type ServiceUpdateIssueDTO struct {
	Id      string  `json:"id"`
	Comment *string `json:"comment,omitempty"`
	Status  *string `json:"status,omitempty"`
}

type ServiceListDTO struct {
	Status   *string `json:"status,omitempty"`
	Page     int     `json:"page"`
	PageSize int     `json:"pageSize"`
}
