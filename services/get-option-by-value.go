package services

import (
	"github.com/eugenepoboykin/go-feedback-api/constant"
	"github.com/eugenepoboykin/go-feedback-api/helpers"
	"github.com/eugenepoboykin/go-feedback-api/schema"
	"github.com/eugenepoboykin/go-feedback-api/utils"
)

func GetOptionByValue(value string) (*schema.IssuesStatus, error) {

	ctx := utils.Ctx()

	var status schema.IssuesStatus

	query := /* sql */ `SELECT value, label FROM statuses WHERE value = $1`

	res := db.QueryRow(ctx, query, value)

	err := res.Scan(&status.Value, &status.Label)

	if err != nil {
		helpers.Log.ErrorLog.Printf(constant.Log_ErrorSelect, err)
		return nil, err
	}

	return &status, nil
}
