package database

import (
	"context"
	"time"

	"github.com/eugenepoboykin/go-feedback-api/internal/models"
)

func (r *IssueRepository) UpdateIssue(ctx context.Context, params models.UpdateIssueArgs) (*models.Issue, error) {
	stmt, errStmt := r.DB.Prepare(queryUpdateIssueItem)

	if errStmt != nil {
		return nil, errStmt
	}

	_, err := stmt.ExecContext(ctx,
		params.Comment,
		params.Status,
		time.Now(),
		params.Id,
	)
	if err != nil {
		return nil, err
	}

	issue, err := r.GetById(ctx, params.Id)
	if err != nil {
		return nil, err
	}

	return issue, nil
}
