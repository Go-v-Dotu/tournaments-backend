package main

import (
	"context"
	"os"

	"tournaments_backend/internal/tournament_management/application"
	"tournaments_backend/internal/tournament_management/infrastructure/api"
	"tournaments_backend/internal/tournament_management/infrastructure/mongodb"
)

func main() {
	ctx := context.TODO()
	app, err := application.NewApp(ctx, mongodb.Config{
		IP:         "127.0.0.1",
		Port:       27017,
		User:       "admin",
		Password:   "admin",
		AuthSource: "admin",
	})
	if err != nil {
		os.Exit(1)
	}

	srv := api.NewServer(app)
	srv.Start()
}
