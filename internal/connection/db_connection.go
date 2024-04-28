package connection

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func (c *Connection) DBConnection() *sql.DB {
	db, err := c.ConnectToDB()

	if err != nil {
		panic("Error " + c.DriverName + " connection, DSN=" + c.DSN)
	}

	defer db.DB.Close()

	return db.DB
}
