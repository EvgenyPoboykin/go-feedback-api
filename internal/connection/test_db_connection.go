package connection

import (
	"context"
	"database/sql"
	"fmt"

	"log"

	_ "github.com/lib/pq"
)

const (
	Log_Ping = "*** Ping database successfully! ***"
)

func (c *Connection) TestDBConnection(db *sql.DB) error {

	if err := db.PingContext(context.Background()); err != nil {
		log.Fatalf("can't connect to db: %s", err)
	}

	fmt.Println(Log_Ping)

	return nil
}
