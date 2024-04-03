package services

import (
	"context"

	"github.com/eugenepoboykin/go-feedback-api/constant"
	"github.com/eugenepoboykin/go-feedback-api/helpers"
	"github.com/eugenepoboykin/go-feedback-api/schema"
)

func GetListAll() ([]schema.Issue, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var isseus []schema.Issue

	query := /* sql */ `SELECT id, isseu_target_uri, isseu_image64, isseu_description, comment, isseu_status, created_at, updated_at, client_id, client_name  FROM isseus ORDER BY created_at DESC`

	rows, err := db.Query(ctx, query)

	if err != nil {
		helpers.Log.ErrorLog.Printf(constant.Log_ErrorSelect, err)
		return nil, err
	}

	for rows.Next() {
		var isseu schema.Issue

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

	return isseus, nil
}
