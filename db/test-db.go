package db

import (
	"context"
	"fmt"

	"github.com/eugenepoboykin/go-feedback-api/constant"

	"github.com/jackc/pgx/v5"
)

func TestDB(d *pgx.Conn) error {
	if err := d.Ping(context.Background()); err != nil {
		fmt.Println("Error", err)
	}

	fmt.Println(constant.Log_Ping)

	return nil
}
