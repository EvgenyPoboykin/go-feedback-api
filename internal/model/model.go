package model

import "time"

type Issue struct {
	Id          string    `json:"id"`
	Uri         string    `json:"uri"`
	Image64     string    `json:"image64"`
	Description string    `json:"description"`
	Comment     *string   `json:"comment"`
	Status      *string   `json:"status"`
	ClientId    *string   `json:"clientId"`
	ClientName  *string   `json:"clientName"`
	Created     time.Time `json:"created"`
	Updated     time.Time `json:"updated"`
}

type Option struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

type IssuesList struct {
	Issues     []Issue `json:"issues"`
	Status     string  `json:"status"`
	Page       int     `json:"page"`
	PageSize   int     `json:"pageSize"`
	TotalCount int     `json:"totalCount"`
}

type AddIssueArgs struct {
	Uri         string
	Image       string
	Description string
	ClientId    string
	ClientName  string
}

type UpdateIssueArgs struct {
	Id      string
	Comment *string
	Status  string
}

type ListArgs struct {
	Status   string
	Page     int
	PageSize int
}

type ErrorMessage struct {
	Status      int32
	Type        string
	Description string
}
