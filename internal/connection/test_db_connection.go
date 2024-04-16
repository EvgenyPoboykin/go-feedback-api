package connection

import (
	"context"
	"database/sql"
	"fmt"

	"log"
)

const (
	Log_Ping = "*** Ping database successfuly! ***"
)

func (c *Connection) TestDBConnection(db *sql.DB) error {

	if err := db.PingContext(context.Background()); err != nil {
		log.Fatalf("can't connect to db: %s", err)
	}

	fmt.Println(Log_Ping)

	return nil
}
