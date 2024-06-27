package queries

import (
	"context"
)

type TournamentQueryService interface {
	GetByID(ctx context.Context, id string) (*Tournament, error)
	GetByHostID(ctx context.Context, hostID string) ([]*Tournament, error)
}

type PlayerQueryService interface {
	GetByTournamentID(ctx context.Context, tournamentID string) ([]*Player, error)
}
