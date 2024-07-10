package client

import (
	"fmt"

	"github.com/eugenepoboykin/go-feedback-api/internal/domain/storage"
	"github.com/jmoiron/sqlx"
)

func NewPostgresRepository(driverName string, dsn string) (*storage.Repository, error) {
	db, err := sqlx.Open(driverName, dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected!")

	return &storage.Repository{
		DB: db.DB,
	}, nil
}
