package services

import (
	"context"

	"github.com/eugenepoboykin/go-feedback-api/constant"
	"github.com/eugenepoboykin/go-feedback-api/helpers"
	"github.com/eugenepoboykin/go-feedback-api/schema"
)

func GetIsseu(id string) (*schema.Issue, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var isseu schema.Issue

	query := /* sql */ `SELECT id, isseu_target_uri, isseu_image64, isseu_description, comment, isseu_status, created_at, updated_at, client_id, client_name FROM isseus WHERE id = $1`

	res := db.QueryRow(ctx, query, id)

	err := res.Scan(&isseu.Id, &isseu.Uri, &isseu.Image64, &isseu.Description, &isseu.Comment, &isseu.Status, &isseu.Created, &isseu.Updated, &isseu.ClientId, &isseu.ClientName)

	if err != nil {
		helpers.Log.ErrorLog.Printf(constant.Log_ErrorSelect, err)
		return nil, err
	}

	return &isseu, nil
}
