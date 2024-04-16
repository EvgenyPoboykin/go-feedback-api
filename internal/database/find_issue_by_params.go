package database

import (
	"context"

	"github.com/eugenepoboykin/go-feedback-api/internal/models"
)

func (r *IssueRepository) FindIsseuByParams(ctx context.Context, params models.AddIssueArgs) (*models.Issue, error) {
	var isseu models.Issue

	stmt, errStmt := r.DB.Prepare(queryFindIssueByParams)

	if errStmt != nil {
		r.Log.Printf(Log_ErrorInsert, errStmt)
		return nil, errStmt
	}

	res := stmt.QueryRowContext(ctx, params.Uri, params.Image, params.Description, params.ClientId, params.ClientName)

	err := res.Scan(&isseu.Id, &isseu.Uri, &isseu.Image64, &isseu.Description, &isseu.Comment, &isseu.Status, &isseu.Created, &isseu.Updated, &isseu.ClientId, &isseu.ClientName)
	if err != nil {
		r.Log.Printf(Log_ErrorInsert, err)
		return nil, err
	}

	return &isseu, nil
}
