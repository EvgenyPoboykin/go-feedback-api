package schema

import "time"

type IssueCreateArgs struct {
	Uri         string `json:"uri"`
	Image64     string `json:"image64"`
	Description string `json:"description"`
}

type IssueUpdateArgs struct {
	Comment string `json:"comment"`
	Status  string `json:"status"`
}

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

type IssuesListArgs struct {
	Status   string `json:"status"`
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
}

type IssuesList struct {
	Issues     []Issue `json:"issues"`
	Status     string  `json:"status"`
	Page       int     `json:"page"`
	PageSize   int     `json:"pageSize"`
	TotalCount int     `json:"totalCount"`
}

type IssuesStatus struct {
	Value string `json:"value"`
	Label string `json:"label"`
}
