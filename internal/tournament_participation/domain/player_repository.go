package domain

import "context"

type PlayerRepository interface {
	Get(ctx context.Context, id string) (*Player, error)
	Save(ctx context.Context, p *Player) error
	Delete(ctx context.Context, p *Player) error
	NextID(ctx context.Context) string
}
