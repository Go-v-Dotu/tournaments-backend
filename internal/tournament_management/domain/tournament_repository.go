package domain

import "context"

type TournamentRepository interface {
	Get(ctx context.Context, id string) (*Tournament, error)
	Save(ctx context.Context, t *Tournament) error
	Delete(ctx context.Context, t *Tournament) error
	NextID(ctx context.Context) string
}
