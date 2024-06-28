package main

import (
	"context"
	"errors"
	"os"
	"strconv"
	"tournaments_backend/internal/application"
	"tournaments_backend/internal/infrastructure/api"
	"tournaments_backend/internal/infrastructure/events_handler"
	"tournaments_backend/internal/infrastructure/mongodb"

	"github.com/sirupsen/logrus"
)

func initMongoConfig() (*mongodb.Config, error) {
	ip := os.Getenv("MONGO_IP")
	if ip == "" {
		return nil, errors.New("mongo ip is empty")
	}
	port, err := strconv.Atoi(os.Getenv("MONGO_PORT"))
	if err != nil {
		return nil, err
	}
	user := os.Getenv("MONGO_USER")
	if user == "" {
		return nil, errors.New("mongo user is empty")
	}
	password := os.Getenv("MONGO_PASSWORD")
	if password == "" {
		return nil, errors.New("mongo password is empty")
	}
	authSource := os.Getenv("MONGO_AUTH_SOURCE")
	if authSource == "" {
		return nil, errors.New("mongo auth source is empty")
	}

	return &mongodb.Config{
		IP:         ip,
		Port:       port,
		User:       user,
		Password:   password,
		AuthSource: authSource,
	}, nil
}

func main() {
	ctx := context.TODO()

	logger := logrus.New()
	rabbitAddr := os.Getenv("RABBIT_ADDR")
	if rabbitAddr == "" {
		logger.Fatal("Error during initialization: RABBIT_ADDR env not set")
	}

	mongoConfig, err := initMongoConfig()
	if err != nil {
		logger.Fatalf("Error during initialization: %v", err)
	}

	app, err := application.NewApp(
		ctx,
		*mongoConfig,
		rabbitAddr,
		logger,
	)
	if err != nil {
		logger.Fatalf("Error during initialization: %v", err)
		os.Exit(1)
	}

	eventsHandler, err := events_handler.NewEventsHandler(app, rabbitAddr, logger)
	if err != nil {
		logger.Fatalf("Error during initialization: %v", err)
		return
	}

	srv := api.NewServer(app, logger)

	go eventsHandler.Start(ctx)

	srv.Start()
}
