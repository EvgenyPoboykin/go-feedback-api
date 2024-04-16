package pool

import (
	"database/sql"
	"time"
)

var dbPool *sql.DB

func SetBDPool(db *sql.DB) {
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	dbPool = db
}

func GetBDPool() *sql.DB {
	return dbPool
}
