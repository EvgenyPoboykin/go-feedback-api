package server

import (
	"fmt"
	"net/http"

	"github.com/eugenepoboykin/go-feedback-api/internal/connection"
	"github.com/eugenepoboykin/go-feedback-api/internal/env"
	"github.com/eugenepoboykin/go-feedback-api/internal/mw"
	"github.com/eugenepoboykin/go-feedback-api/internal/pool"
	"github.com/eugenepoboykin/go-feedback-api/router"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Start() {

	enviroment := env.NewEnv()

	conn := connection.NewDBConnection("postgres", enviroment.DSN)
	DB := conn.DBConnection()

	pool.SetPool(DB)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(mw.Auth)

	router.Support(r)

	http.ListenAndServe(enviroment.AppPort, r)

	fmt.Printf("Start Server on port : '%s'!", enviroment.AppPort)
}
