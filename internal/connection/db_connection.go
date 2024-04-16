package connection

import "database/sql"

func (c *Connection) DBConnection() *sql.DB {
	db, err := c.ConnectToDB()

	if err != nil {
		panic("Error " + c.DriverName + " connection, DSN=" + c.DSN)
	}

	defer db.DB.Close()

	return db.DB
}
