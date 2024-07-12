package main

import (
	"context"

	"github.com/eugenepoboykin/go-feedback-api/internal/lib/logger"
	"github.com/eugenepoboykin/go-feedback-api/pkg/postgres"
)

func main() {
	ctx := context.Background()

	container, err := postgres.NewRegistry(ctx)
	if err != nil {
		logger.Log.ErrorLog.Print(err)
	}

	err = container.Run()
	if err != nil {
		logger.Log.ErrorLog.Print(err)
	}
}
