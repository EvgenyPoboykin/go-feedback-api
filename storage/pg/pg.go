package pg

import (
	"database/sql"

	"github.com/eugenepoboykin/go-feedback-api/internal/database"
	"github.com/eugenepoboykin/go-feedback-api/internal/lib/logger"
)

func NewPostgresRepo(db *sql.DB) *database.IssueRepository {
	return &database.IssueRepository{
		DB:  db,
		Log: logger.Log.ErrorLog,
	}
}
