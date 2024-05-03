package database

import (
	"database/sql"
)

type IssueRepository struct {
	DB *sql.DB
}
