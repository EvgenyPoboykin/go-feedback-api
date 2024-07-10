package models

import "time"

type ServiceIssue struct {
	Id          string    `json:"id"`
	Uri         string    `json:"uri"`
	Image64     string    `json:"image64"`
	Description string    `json:"description"`
	Comment     *string   `json:"comment.omitempty"`
	Status      *string   `json:"status.omitempty"`
	ClientId    string    `json:"clientId"`
	ClientName  string    `json:"clientName"`
	Created     time.Time `json:"created"`
	Updated     time.Time `json:"updated"`
}

type ServiceOption struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

type ServiceIssuesList struct {
	Issues     []ServiceIssue `json:"issues"`
	Status     *string        `json:"status.omitempty"`
	Page       int            `json:"page"`
	PageSize   int            `json:"pageSize"`
	TotalCount int            `json:"totalCount"`
}
