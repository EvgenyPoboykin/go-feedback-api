package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type DB struct {
	DB *pgx.Conn
}

var DbConnect = &DB{}

func Connect(dsn string) (*DB, error) {
	connect, err := pgx.Connect(context.Background(), dsn)

	if err != nil {
		return nil, err
	}

	if err = TestDB(connect); err != nil {
		return nil, err
	}

	DbConnect.DB = connect

	return DbConnect, nil
}
