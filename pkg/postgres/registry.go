package postgres

import (
	"context"
	"fmt"
	"net/http"

	"github.com/eugenepoboykin/go-feedback-api/internal/middleware/auth"

	"github.com/eugenepoboykin/go-feedback-api/internal/domain/env"
	"github.com/eugenepoboykin/go-feedback-api/internal/domain/handlers"
	serviceIssue "github.com/eugenepoboykin/go-feedback-api/internal/domain/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Registry struct {
	ctx     context.Context
	service *serviceIssue.Service
}

func NewRegistry(ctx context.Context) (*Registry, error) {
	r := &Registry{}

	err := env.NewEnv()
	if err != nil {
		return nil, err
	}

	service, err := r.ServiceInit(ctx)
	if err != nil {
		return nil, err
	}

	r.service = service

	return r, nil
}

func (r *Registry) ServiceInit(ctx context.Context) (*serviceIssue.Service, error) {
	s := NewProvider(ctx)

	service, err := s.Registry()
	if err != nil {
		return nil, err
	}

	return service, nil
}

func (r *Registry) Server() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(auth.Auth)

	c := handlers.NewHandlers(r.ctx, r.service)
	c.RegisterRoutes(router)

	return router
}

func (r *Registry) Run() error {
	router := r.Server()

	err := http.ListenAndServe(env.Environment.AppPort, router)
	if err != nil {
		return err
	}

	fmt.Printf("Start Server on port : '%s'!", env.Environment.AppPort)

	return nil
}
