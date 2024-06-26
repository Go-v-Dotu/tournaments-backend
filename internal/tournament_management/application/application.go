package application

import (
	"context"
	"fmt"

	"tournaments_backend/internal/tournament_management/infrastructure/mongodb"
	"tournaments_backend/internal/tournament_management/usecases"
)

type App struct {
	UseCases *usecases.UseCases
}

func NewApp(ctx context.Context, mongoConfig mongodb.Config) (*App, error) {
	dbClient, err := mongodb.NewClient(ctx, mongoConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create Mongo client: %v", err)
	}

	hostRepo := mongodb.NewHostRepository(*dbClient)
	playerRepo := mongodb.NewPlayerRepository(*dbClient)
	tournamentRepo := mongodb.NewTournamentRepository(*dbClient)
	tournamentQueryService := mongodb.NewTournamentQueryService(*dbClient)

	app := App{
		UseCases: usecases.NewUseCases(hostRepo, playerRepo, tournamentRepo, tournamentQueryService),
	}

	return &app, nil
}
