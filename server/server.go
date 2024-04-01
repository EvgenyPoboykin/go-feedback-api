package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/eugenepoboykin/go-feedback-api/authmiddleware"
	"github.com/eugenepoboykin/go-feedback-api/constant"
	"github.com/eugenepoboykin/go-feedback-api/db"
	"github.com/eugenepoboykin/go-feedback-api/helpers"
	"github.com/eugenepoboykin/go-feedback-api/router"
	"github.com/eugenepoboykin/go-feedback-api/services"
	"github.com/eugenepoboykin/go-feedback-api/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/spf13/viper"
)

func Start() {
	postgres, err := db.Connect(utils.GetEnv("SUPPORT_SERVICE_CONNECT_DATABASE_URL", constant.DefaultDsn))

	if err != nil {
		helpers.Log.ErrorLog.Fatal("Cannot connect database!")

		return
	}

	defer postgres.DB.Close(context.Background())

	services.New(postgres.DB)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(authmiddleware.Auth)

	router.Support(r)

	http.ListenAndServe(viper.GetString("port"), r)

	fmt.Printf("Start Server")
}
