package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/eugenepoboykin/go-feedback-api/internal/domain/storage"
	"github.com/eugenepoboykin/go-feedback-api/pkg/attempts"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresClient(ctx context.Context, dsn string) (*storage.Repository, error) {
	var pool *pgxpool.Pool

	err := attempts.Attempts(func() error {
		_ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		defer cancel()

		dbPool, err := pgxpool.New(_ctx, dsn)
		if err != nil {
			return err
		}

		pool = dbPool

		return nil
	}, 5, 5*time.Second)
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected!")

	return &storage.Repository{
		Pool: pool,
	}, nil
}
