package services

import (
	"github.com/jackc/pgx/v5"
)

var db *pgx.Conn

func New(dbConnection *pgx.Conn) {
	db = dbConnection
}
