package models

import "context"

type IssueRepository interface {
	GetListByAdmin(ctx context.Context) (*[]Issue, error)
	GetListByEmployee(ctx context.Context, userId string) (*[]Issue, error)
	GetById(ctx context.Context, issueId string) (*Issue, error)
	AddIssueItem(ctx context.Context, params AddIssueArgs) (*Issue, error)
	UpdateIssue(ctx context.Context, params UpdateIssueArgs) (*Issue, error)
	FindIssueByParams(ctx context.Context, params AddIssueArgs) (*Issue, error)
	GetOptions(ctx context.Context) (*[]Option, error)
	GetOptionByValue(ctx context.Context, value string) (*Option, error)
}
