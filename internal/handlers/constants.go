package handlers

import (
	"github.com/eugenepoboykin/go-feedback-api/internal/database"
)

const (
	Admin    = "admin"
	Employee = "employee"
)

type ApiSettings struct {
	conn *database.IssueRepository
}

func NewApi(conn *database.IssueRepository) *ApiSettings {
	return &ApiSettings{
		conn: conn,
	}
}
