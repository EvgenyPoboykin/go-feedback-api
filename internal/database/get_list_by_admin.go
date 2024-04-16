package database

import (
	"context"

	"github.com/eugenepoboykin/go-feedback-api/internal/models"
)

func (r *IssueRepository) GetListByAdmin(ctx context.Context) (*[]models.Issue, error) {
	var isseus []models.Issue

	stmt, errStmt := r.DB.Prepare(queryListByAdmin)

	if errStmt != nil {
		r.Log.Printf(Log_ErrorInsert, errStmt)
		return nil, errStmt
	}

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		r.Log.Printf(Log_ErrorSelect, err)
		return nil, err
	}

	for rows.Next() {
		var isseu models.Issue

		err := rows.Scan(
			&isseu.Id,
			&isseu.Uri,
			&isseu.Image64,
			&isseu.Description,
			&isseu.Comment,
			&isseu.Status,
			&isseu.Created,
			&isseu.Updated,
			&isseu.ClientId,
			&isseu.ClientName,
		)
		if err != nil {
			return nil, err
		}

		isseus = append(isseus, isseu)
	}

	return &isseus, nil
}
