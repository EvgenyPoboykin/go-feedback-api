package main

import (
	"context"

	_ "github.com/eugenepoboykin/go-feedback-api/docs"
	"github.com/eugenepoboykin/go-feedback-api/internal/domain/registry"
	"github.com/eugenepoboykin/go-feedback-api/internal/lib/logger"
	_ "github.com/lib/pq"
)

func main() {

	ctx := context.Background()

	container, err := registry.NewRegistry(ctx)

	if err != nil {
		logger.Log.ErrorLog.Print(err)
	}

	err = container.Run()

	if err != nil {
		logger.Log.ErrorLog.Print(err)
	}
}
