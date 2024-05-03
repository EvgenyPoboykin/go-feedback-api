package connection

import (
	"database/sql"

	"github.com/eugenepoboykin/go-feedback-api/internal/lib/logger"
	_ "github.com/lib/pq"
)

func (c *Connection) DBConnection() *sql.DB {
	db, err := c.ConnectToDB()

	if err != nil {
		logger.Log.InfoLog.Printf("Error %s connection, DSN=%s", c.DriverName, c.DSN)
	}

	defer db.DB.Close()

	return db.DB
}
