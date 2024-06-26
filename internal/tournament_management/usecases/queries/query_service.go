package queries

import "context"

type TournamentQueryService interface {
	GetByHostID(ctx context.Context, hostID string) (Tournaments, error)
}
