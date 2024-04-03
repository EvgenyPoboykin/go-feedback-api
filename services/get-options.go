package services

import (
	"github.com/eugenepoboykin/go-feedback-api/constant"
	"github.com/eugenepoboykin/go-feedback-api/helpers"
	"github.com/eugenepoboykin/go-feedback-api/schema"
	"github.com/eugenepoboykin/go-feedback-api/utils"
)

func GetOptions() ([]schema.IssuesStatus, error) {

	ctx := utils.Ctx()

	var statuses []schema.IssuesStatus

	query := /* sql */ `SELECT value, label FROM statuses`

	rows, err := db.Query(ctx, query)

	if err != nil {
		helpers.Log.ErrorLog.Printf(constant.Log_ErrorSelect, err)
		return nil, err
	}

	for rows.Next() {
		var status schema.IssuesStatus

		err := rows.Scan(
			&status.Value,
			&status.Label,
		)

		if err != nil {
			return nil, err
		}

		statuses = append(statuses, status)
	}

	return statuses, nil
}
