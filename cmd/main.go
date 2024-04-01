package main

import (
	"github.com/eugenepoboykin/go-feedback-api/constant"
	"github.com/eugenepoboykin/go-feedback-api/helpers"
	"github.com/eugenepoboykin/go-feedback-api/server"
	"github.com/eugenepoboykin/go-feedback-api/utils"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		helpers.Log.ErrorLog.Print(constant.Log_ErrorNoEnv)
	}
}

func main() {
	if err := utils.InitConfig(); err != nil {
		helpers.Log.ErrorLog.Fatalf(constant.Log_ErrorInitialConfig, err.Error())
	}

	server.Start()
}
