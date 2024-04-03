package services

import (
	"context"
	"time"

	"github.com/eugenepoboykin/go-feedback-api/constant"
	"github.com/eugenepoboykin/go-feedback-api/helpers"
	"github.com/eugenepoboykin/go-feedback-api/schema"
)

func CreateIsseu(uri string, image string, description string, clientId string, clientName string) (*schema.Issue, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := /* sql */ `INSERT INTO isseus (isseu_target_uri, isseu_image64, isseu_description, created_at, updated_at, client_id, client_name) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := db.Exec(ctx, query,
		uri,
		image,
		description,
		time.Now(),
		time.Now(),
		clientId,
		clientName,
	)

	if err != nil {
		helpers.Log.ErrorLog.Printf(constant.Log_ErrorInsert, err)
		return nil, err
	}

	isseu, err := CreateIsseuReturn(uri, image, description)

	if err != nil {
		helpers.Log.ErrorLog.Printf(constant.Log_ErrorInsert, err)
		return nil, err
	}

	return isseu, nil
}
