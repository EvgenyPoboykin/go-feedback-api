package models

type StorageAddIssueDTO struct {
	Uri         string `db:"issue_target_uri"`
	Image64     string `db:"issue_image64"`
	Description string `db:"issue_description"`
	ClientId    string `db:"client_id"`
	ClientName  string `db:"client_name"`
}

type StorageUpdateIssueDTO struct {
	Id      string  `db:"id"`
	Comment *string `db:"comment,omitempty"`
	Status  *string `db:"issue_status,omitempty"`
}

type StorageListDTO struct {
	Status   *string `db:"issue_status,omitempty"`
	Page     int     `db:"page"`
	PageSize int     `db:"page_size"`
}
