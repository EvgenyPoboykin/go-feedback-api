package database

import (
	"context"
	"time"

	"github.com/eugenepoboykin/go-feedback-api/internal/models"
)

func (r *IssueRepository) UpdateIssue(ctx context.Context, params models.UpdateIssueArgs) (*models.Issue, error) {
	stmt, errStmt := r.DB.Prepare(queryUpdateIssueItem)

	if errStmt != nil {
		r.Log.Printf(Log_ErrorInsert, errStmt)
		return nil, errStmt
	}

	_, err := stmt.ExecContext(ctx,
		params.Comment,
		params.Status,
		time.Now(),
		params.Id,
	)
	if err != nil {
		r.Log.Printf(Log_ErrorUpdate, err)
		return nil, err
	}

	isseu, err := r.GetById(ctx, params.Id)
	if err != nil {
		r.Log.Printf(Log_ErrorSelect, err)
		return nil, err
	}

	return isseu, nil
}
