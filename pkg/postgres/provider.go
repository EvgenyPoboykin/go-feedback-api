package postgres

import (
	"context"

	"github.com/eugenepoboykin/go-feedback-api/internal/domain/service"
	"github.com/eugenepoboykin/go-feedback-api/internal/domain/storage"
)

type Provider struct {
	ctx context.Context
}

func NewProvider(ctx context.Context) *Provider {
	return &Provider{
		ctx: ctx,
	}
}

func (p *Provider) Registry() (*service.Service, error) {
	db, err := NewPostgresClient(p.ctx, PostgresDsn())
	if err != nil {
		return nil, err
	}

	repo, err := storage.NewRepository(db.Pool)
	if err != nil {
		return nil, err
	}

	s, err := service.NewService(repo)
	if err != nil {
		return nil, err
	}

	return s, nil
}
