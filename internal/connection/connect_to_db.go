package connection

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func (c *Connection) ConnectToDB() (*DBConnection, error) {

	db, err := sql.Open(c.DriverName, c.DSN)

	if err != nil {
		return nil, err
	}

	if err = c.TestDBConnection(db); err != nil {
		return nil, err
	}

	return &DBConnection{
		DB: db,
	}, nil
}
