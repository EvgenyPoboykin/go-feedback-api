package connection

import "database/sql"

func (c *Connection) ConnectToDB() (*DBConnection, error) {
	DB, err := sql.Open(c.DriverName, c.DSN)
	if err != nil {
		return nil, err
	}

	if err = c.TestDBConnection(DB); err != nil {
		return nil, err
	}

	return &DBConnection{
		DB: DB,
	}, nil
}
