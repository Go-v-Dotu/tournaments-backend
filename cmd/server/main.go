package main

import (
	"context"
	"os"

	"tournaments_backend/internal/application"
	"tournaments_backend/internal/infrastructure/api"
	"tournaments_backend/internal/infrastructure/events_handler"
	"tournaments_backend/internal/infrastructure/mongodb"

	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.TODO()

	logger := logrus.New()
	rabbitAddr := "amqp://user:password@127.0.0.1:5672/"

	app, err := application.NewApp(
		ctx,
		mongodb.Config{
			IP:         "127.0.0.1",
			Port:       27017,
			User:       "admin",
			Password:   "admin",
			AuthSource: "admin",
		},
		rabbitAddr,
		logger,
	)
	if err != nil {
		logger.Fatal("Error during initialization", err)
		os.Exit(1)
	}

	eventsHandler, err := events_handler.NewEventsHandler(app, rabbitAddr, logger)
	if err != nil {
		logger.Fatal("Error during initialization", err)
		return
	}

	srv := api.NewServer(app, logger)

	go eventsHandler.Start(ctx)

	srv.Start()
}
