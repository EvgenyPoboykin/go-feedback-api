package database

import (
	"context"

	"github.com/eugenepoboykin/go-feedback-api/internal/models"
)

func (r *IssueRepository) GetOptions(ctx context.Context) (*[]models.Option, error) {
	var options []models.Option

	stmt, errStmt := r.DB.Prepare(queryGetOptions)

	if errStmt != nil {
		return nil, errStmt
	}

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var option models.Option

		err := rows.Scan(
			&option.Value,
			&option.Label,
		)
		if err != nil {
			return nil, err
		}

		options = append(options, option)
	}

	return &options, nil
}
