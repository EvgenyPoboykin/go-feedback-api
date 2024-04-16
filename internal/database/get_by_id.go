package database

import (
	"context"

	"github.com/eugenepoboykin/go-feedback-api/internal/models"
)

func (r *IssueRepository) GetById(ctx context.Context, issueId string) (*models.Issue, error) {
	var isseu models.Issue

	stmt, errStmt := r.DB.Prepare(queryGetById)

	if errStmt != nil {
		r.Log.Printf(Log_ErrorInsert, errStmt)
		return nil, errStmt
	}

	res := stmt.QueryRowContext(ctx, issueId)

	err := res.Scan(&isseu.Id, &isseu.Uri, &isseu.Image64, &isseu.Description, &isseu.Comment, &isseu.Status, &isseu.Created, &isseu.Updated, &isseu.ClientId, &isseu.ClientName)
	if err != nil {
		r.Log.Printf(Log_ErrorSelect, err)
		return nil, err
	}

	return &isseu, nil
}
