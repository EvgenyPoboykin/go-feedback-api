package router

import (
	"github.com/eugenepoboykin/go-feedback-api/controllers"
	"github.com/eugenepoboykin/go-feedback-api/utils"

	"github.com/go-chi/chi/v5"
)

func Support(r chi.Router) {
	r.Post(utils.VersionApiUrl("/list"), controllers.List)
	r.Get(utils.VersionApiUrl("/list/{id}"), controllers.Item)
	r.Put(utils.VersionApiUrl("/list/{id}"), controllers.Update)
	r.Post(utils.VersionApiUrl("/list/create"), controllers.Create)
	r.Get(utils.VersionApiUrl("/options"), controllers.Options)
}
