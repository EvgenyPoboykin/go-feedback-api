package services

import (
	"time"

	"github.com/jackc/pgx/v5"
)

var db *pgx.Conn

const dbTimeout = time.Second * 3

func New(dbConnection *pgx.Conn) {
	db = dbConnection
}
