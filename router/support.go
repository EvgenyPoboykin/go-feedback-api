package router

import (
	"fmt"

	"github.com/eugenepoboykin/go-feedback-api/internal/endpoints"
	"github.com/eugenepoboykin/go-feedback-api/internal/handlers"
	"github.com/eugenepoboykin/go-feedback-api/internal/pool"
	"github.com/eugenepoboykin/go-feedback-api/storage/pg"

	"github.com/go-chi/chi/v5"
)

func Support(r chi.Router) {
	storage := pg.NewPostgresRepo(pool.GetPool())

	endpoints := endpoints.NewApiVersion("v1")

	h := handlers.NewApi(storage)
	fmt.Print(h)

	r.Post(endpoints.ListAdmin(), h.ListAdmin)
	r.Post(endpoints.ListEmployee(), h.ListEmployee)
	r.Get(endpoints.Item(), h.Item)
	r.Put(endpoints.Item(), h.Update)
	r.Post(endpoints.Create(), h.Create)
	r.Get(endpoints.Options(), h.Options)
}
