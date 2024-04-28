package pool

import (
	"database/sql"
	"time"
)

var dbPool *sql.DB

func SetPool(db *sql.DB) {
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	dbPool = db
}

func GetPool() *sql.DB {
	return dbPool
}
