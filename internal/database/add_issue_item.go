package database

import (
	"context"
	"time"

	"github.com/eugenepoboykin/go-feedback-api/internal/models"
)

func (r *IssueRepository) AddIssueItem(ctx context.Context, params models.AddIssueArgs) (*models.Issue, error) {
	stmt, errStmt := r.DB.Prepare(queryAddIssueItem)

	if errStmt != nil {
		return nil, errStmt
	}

	_, err := stmt.ExecContext(ctx,
		params.Uri,
		params.Image,
		params.Description,
		time.Now(),
		time.Now(),
		params.ClientId,
		params.ClientName,
	)
	if err != nil {
		return nil, err
	}

	issue, err := r.FindIssueByParams(ctx, params)
	if err != nil {
		return nil, err
	}

	return issue, nil
}
