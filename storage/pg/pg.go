package pg

import (
	"database/sql"

	"github.com/eugenepoboykin/go-feedback-api/internal/database"
)

func NewPostgresRepo(db *sql.DB) *database.IssueRepository {
	return &database.IssueRepository{
		DB: db,
	}
}
