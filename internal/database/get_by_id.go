package database

import (
	"context"

	"github.com/eugenepoboykin/go-feedback-api/internal/models"
)

func (r *IssueRepository) GetById(ctx context.Context, issueId string) (*models.Issue, error) {
	var issue models.Issue

	stmt, errStmt := r.DB.Prepare(queryGetById)

	if errStmt != nil {
		return nil, errStmt
	}

	res := stmt.QueryRowContext(ctx, issueId)

	err := res.Scan(&issue.Id, &issue.Uri, &issue.Image64, &issue.Description, &issue.Comment, &issue.Status, &issue.Created, &issue.Updated, &issue.ClientId, &issue.ClientName)
	if err != nil {
		return nil, err
	}

	return &issue, nil
}
