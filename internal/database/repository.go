package database

import (
	"database/sql"
	"log"
)

type IssueRepository struct {
	DB  *sql.DB
	Log *log.Logger
}
