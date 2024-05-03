package pagination

import (
	"github.com/eugenepoboykin/go-feedback-api/internal/models"
)

func Pagination(pageNum int, pageSize int, sliceLength int) (int, int) {
	start := pageNum * pageSize

	if start > sliceLength {
		start = sliceLength
	}

	end := start + pageSize
	if end > sliceLength {
		end = sliceLength
	}

	return start, end
}

func IssuePerPage(issues []models.Issue, body models.ListArgs) []models.Issue {
	start, end := Pagination(body.Page-1, body.PageSize, len(issues))
	pagedSlice := issues[start:end]

	return pagedSlice
}
