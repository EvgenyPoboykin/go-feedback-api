package router

import (
	"github.com/eugenepoboykin/go-feedback-api/internal/endpoints"
	"github.com/eugenepoboykin/go-feedback-api/internal/handlers"
	"github.com/eugenepoboykin/go-feedback-api/internal/pool"
	"github.com/eugenepoboykin/go-feedback-api/storage/pg"

	"github.com/go-chi/chi/v5"
)

func Support(r chi.Router) {
	db := pg.NewPostgresRepo(pool.GetBDPool())

	endpoints := endpoints.NewApiVersion("v1")

	r.Post(endpoints.ListAdmin(), handlers.ListAdmin(db))
	r.Post(endpoints.ListEmployee(), handlers.ListEmployee(db))
	r.Get(endpoints.Item(), handlers.Item(db))
	r.Put(endpoints.Item(), handlers.Update(db))
	r.Post(endpoints.Create(), handlers.Create(db))
	r.Get(endpoints.Options(), handlers.Options(db))
}
