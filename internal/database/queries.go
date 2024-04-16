package database

var queryAddIssueItem = /* sql */ `INSERT INTO isseus (isseu_target_uri, isseu_image64, isseu_description, created_at, updated_at, client_id, client_name) VALUES ($1, $2, $3, $4, $5, $6, $7)`

var queryFindIssueByParams = /* sql */ `SELECT id, isseu_target_uri, isseu_image64, isseu_description, comment, isseu_status, created_at, updated_at, client_id, client_name FROM isseus WHERE isseu_target_uri = $1 AND isseu_image64 = $2 AND isseu_description = $3 AND client_id = $4 AND client_name = $5`

var queryGetById = /* sql */ `SELECT id, isseu_target_uri, isseu_image64, isseu_description, comment, isseu_status, created_at, updated_at, client_id, client_name FROM isseus WHERE id = $1`

var queryListByAdmin = /* sql */ `SELECT id, isseu_target_uri, isseu_image64, isseu_description, comment, isseu_status, created_at, updated_at, client_id, client_name  FROM isseus ORDER BY created_at DESC`

var queryListByEmployee = /* sql */ `SELECT id, isseu_target_uri, isseu_image64, isseu_description, comment, isseu_status, created_at, updated_at, client_id, client_name  FROM isseus WHERE client_id = $1 ORDER BY created_at DESC`

var queryGetOptionByValue = /* sql */ `SELECT value, label FROM statuses WHERE value = $1`

var queryGetOptions = /* sql */ `SELECT value, label FROM statuses`

var queryUpdateIssueItem = /* sql */ `UPDATE isseus SET comment=$1, isseu_status=$2, updated_at=$3 WHERE id = $4`
