package models

import "time"

type StorageIssue struct {
	Id          string    `db:"id"`
	Uri         string    `db:"issue_target_uri"`
	Image64     string    `db:"issue_image64"`
	Description string    `db:"issue_description"`
	Comment     *string   `db:"comment"`
	Status      *string   `db:"issue_status"`
	ClientId    *string   `db:"client_id"`
	ClientName  *string   `db:"client_name"`
	Created     time.Time `db:"created_at"`
	Updated     time.Time `db:"updated_at"`
}

type StorageOption struct {
	Value string `db:"value"`
	Label string `db:"label"`
}
