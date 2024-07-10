package storage

var (
	QueryAddIssueItem = `INSERT INTO issues (issue_target_uri, issue_image64, issue_description, created_at, updated_at, client_id, client_name) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	QueryFindIssueByParams = `SELECT id, issue_target_uri, issue_image64, issue_description, comment, issue_status, created_at, updated_at, client_id, client_name FROM issues WHERE issue_target_uri = $1 AND issue_image64 = $2 AND issue_description = $3 AND client_id = $4 AND client_name = $5`

	QueryGetById = `SELECT id, issue_target_uri, issue_image64, issue_description, comment, issue_status, created_at, updated_at, client_id, client_name FROM issues WHERE id = $1`

	QueryListByAdmin = `SELECT id, issue_target_uri, issue_image64, issue_description, comment, issue_status, created_at, updated_at, client_id, client_name  FROM issues ORDER BY created_at DESC`

	QueryListByEmployee = `SELECT id, issue_target_uri, issue_image64, issue_description, comment, issue_status, created_at, updated_at, client_id, client_name  FROM issues WHERE client_id = $1 ORDER BY created_at DESC`

	QueryGetOptionByValue = `SELECT value, label FROM statuses WHERE value = $1`

	QueryGetOptions = `SELECT value, label FROM statuses`

	QueryUpdateIssueItem = `UPDATE issues SET comment=$1, issue_status=$2, updated_at=$3 WHERE id = $4`
)
