package database

import (
	"context"
	"time"

	"github.com/eugenepoboykin/go-feedback-api/internal/models"
)

func (r *IssueRepository) AddIssueItem(ctx context.Context, params models.AddIssueArgs) (*models.Issue, error) {
	stmt, errStmt := r.DB.Prepare(queryAddIssueItem)

	if errStmt != nil {
		r.Log.Printf(Log_ErrorInsert, errStmt)
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
		r.Log.Printf(Log_ErrorInsert, err)
		return nil, err
	}

	isseu, err := r.FindIsseuByParams(ctx, params)
	if err != nil {
		r.Log.Printf(Log_ErrorInsert, err)
		return nil, err
	}

	return isseu, nil
}
