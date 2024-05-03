package connection

import (
	"database/sql"
)

type DBConnection struct {
	DB *sql.DB
}

type Connection struct {
	DriverName string
	DSN        string
}

func NewDBConnection(driverName string, dsn string) *Connection {
	return &Connection{
		DriverName: driverName,
		DSN:        dsn,
	}
}
