package services

import (
	"time"

	"github.com/eugenepoboykin/go-feedback-api/constant"
	"github.com/eugenepoboykin/go-feedback-api/helpers"
	"github.com/eugenepoboykin/go-feedback-api/schema"
	"github.com/eugenepoboykin/go-feedback-api/utils"
)

func UpdateIsseu(id string, comment string, status string) (*schema.Issue, error) {

	ctx := utils.Ctx()

	query := /* sql */ `UPDATE isseus SET comment=$1, isseu_status=$2, updated_at=$3 WHERE id = $4`

	_, err := db.Exec(ctx, query,
		comment,
		status,
		time.Now(),
		id,
	)

	if err != nil {
		helpers.Log.ErrorLog.Printf(constant.Log_ErrorUpdate, err)
		return nil, err
	}

	isseu, err := GetIsseu(id)

	if err != nil {
		helpers.Log.ErrorLog.Printf(constant.Log_ErrorSelect, err)
		return nil, err
	}

	return isseu, nil
}
