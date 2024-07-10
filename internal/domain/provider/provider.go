package provider

import (
	"context"
	"github.com/eugenepoboykin/go-feedback-api/internal/domain/service"
	"github.com/eugenepoboykin/go-feedback-api/internal/domain/storage"
	"github.com/eugenepoboykin/go-feedback-api/pkg/client"
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

	db, err := client.NewPostgresRepository("postgres", PostgresDsn())
	if err != nil {
		return nil, err
	}

	repo := storage.NewRepository(db.DB)

	s, err := service.NewService(repo)
	if err != nil {
		return nil, err
	}

	return s, nil

}
