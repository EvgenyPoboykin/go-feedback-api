package helpers

import (
	"github.com/eugenepoboykin/go-feedback-api/schema"
	"github.com/eugenepoboykin/go-feedback-api/utils"
)

func IssuePerPage(isseus []schema.Issue, body schema.IssuesListArgs) []schema.Issue {
	start, end := utils.Pagination(body.Page-1, body.PageSize, len(isseus))
	pagedSlice := isseus[start:end]

	return pagedSlice
}
