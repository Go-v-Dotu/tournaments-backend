package domain

import "context"

type PlayerRepository interface {
	Get(ctx context.Context, id string) (*Player, error)
	GetByUserID(ctx context.Context, userID string) (*Player, error)
	Save(ctx context.Context, player *Player) error
	Delete(ctx context.Context, player *Player) error
	NextID(ctx context.Context) string
}
