package database

import (
	"context"

	"github.com/eugenepoboykin/go-feedback-api/internal/models"
)

func (r *IssueRepository) GetListByAdmin(ctx context.Context) (*[]models.Issue, error) {
	var issues []models.Issue

	stmt, errStmt := r.DB.Prepare(queryListByAdmin)

	if errStmt != nil {
		return nil, errStmt
	}

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var issue models.Issue

		err := rows.Scan(
			&issue.Id,
			&issue.Uri,
			&issue.Image64,
			&issue.Description,
			&issue.Comment,
			&issue.Status,
			&issue.Created,
			&issue.Updated,
			&issue.ClientId,
			&issue.ClientName,
		)
		if err != nil {
			return nil, err
		}

		issues = append(issues, issue)
	}

	return &issues, nil
}
