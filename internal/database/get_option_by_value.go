package database

import (
	"context"

	"github.com/eugenepoboykin/go-feedback-api/internal/models"
)

func (r *IssueRepository) GetOptionByValue(ctx context.Context, value string) (*models.Option, error) {
	var option models.Option

	stmt, errStmt := r.DB.Prepare(queryGetOptionByValue)

	if errStmt != nil {
		r.Log.Printf(Log_ErrorInsert, errStmt)
		return nil, errStmt
	}

	res := stmt.QueryRowContext(ctx, value)

	err := res.Scan(&option.Value, &option.Label)
	if err != nil {
		r.Log.Printf(Log_ErrorSelect, err)
		return nil, err
	}

	return &option, nil
}
