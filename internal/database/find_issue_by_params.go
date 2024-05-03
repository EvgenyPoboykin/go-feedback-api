package database

import (
	"context"

	"github.com/eugenepoboykin/go-feedback-api/internal/models"
)

func (r *IssueRepository) FindIssueByParams(ctx context.Context, params models.AddIssueArgs) (*models.Issue, error) {
	var issue models.Issue

	stmt, errStmt := r.DB.Prepare(queryFindIssueByParams)

	if errStmt != nil {
		return nil, errStmt
	}

	res := stmt.QueryRowContext(ctx, params.Uri, params.Image, params.Description, params.ClientId, params.ClientName)

	err := res.Scan(&issue.Id, &issue.Uri, &issue.Image64, &issue.Description, &issue.Comment, &issue.Status, &issue.Created, &issue.Updated, &issue.ClientId, &issue.ClientName)
	if err != nil {
		return nil, err
	}

	return &issue, nil
}
