package postgres

import (
	"context"
	"errors"

	"tournaments_backend/internal/tournament_management/domain"
	"tournaments_backend/internal/tournament_management/infrastructure/postgres/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type playerRepository struct {
	client Client
}

var (
	_ domain.PlayerRepository = (*playerRepository)(nil)
)

func NewPlayerRepository(client Client) domain.PlayerRepository {
	return &playerRepository{client: client}
}

func (r *playerRepository) GetByUserID(ctx context.Context, userID string) (*domain.Player, error) {
	query := "SELECT id, user_id, username FROM players WHERE user_id = $1"

	var playerModel models.Player
	if err := r.client.QueryRow(ctx, query, userID).Scan(
		&playerModel.ID,
		&playerModel.UserID,
		&playerModel.Username,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, errors.New("")
	}

	return playerModel.ToEntity(), nil
}

func (r *playerRepository) Save(ctx context.Context, player *domain.Player) error {
	query := "INSERT INTO players (id, user_id, username) VALUES ($1, $2, $3)"

	playerModel := models.NewPlayer(player)
	if _, err := r.client.Exec(ctx, query, playerModel.ID, playerModel.UserID, playerModel.Username); err != nil {
		return errors.New("")
	}

	return nil
}

func (r *playerRepository) Delete(ctx context.Context, player *domain.Player) error {
	return nil
}

func (r *playerRepository) NextID(_ context.Context) string {
	return uuid.New().String()
}
