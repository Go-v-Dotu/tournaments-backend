package postgres

import (
	"context"
	"errors"

	"tournaments_backend/internal/domain"
	"tournaments_backend/internal/infrastructure/postgres/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type tournamentRepository struct {
	client Client
}

var (
	_ domain.TournamentRepository = (*tournamentRepository)(nil)
)

func NewTournamentRepository(client Client) domain.TournamentRepository {
	return &tournamentRepository{client: client}
}

func (r *tournamentRepository) Get(ctx context.Context, id string) (*domain.Tournament, error) {
	query := "SELECT id, host_id, title, date, state FROM tournaments WHERE id = $1"

	var tournamentModel models.Tournament
	if err := r.client.QueryRow(ctx, query, id).Scan(
		&tournamentModel.ID,
		&tournamentModel.HostID,
		&tournamentModel.Title,
		&tournamentModel.Date,
		&tournamentModel.State,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, errors.New("")
	}

	query = "SELECT max_players FROM tournament_settings WHERE tournament_id = $1"
	var settingsModel models.Settings

	query = "SELECT player_id, dropped, has_commander FROM enrolled_players WHERE tournament_id = $1"

	return hostModel.ToEntity(), nil
}

func (r *tournamentRepository) Save(ctx context.Context, tournament *domain.Tournament) error {
	query := "INSERT INTO tournaments (id, host_id, title, date, state) VALUES ($1, $2, $3, $4, $5)"

	tournamentModel := models.New(host)
	if _, err := r.client.Exec(ctx, query, hostModel.ID, hostModel.UserID, host.Username); err != nil {
		return errors.New("")
	}

	return nil
}

func (r *tournamentRepository) Delete(ctx context.Context, tournament *domain.Tournament) error {
	return nil
}

func (r *tournamentRepository) NextID(_ context.Context) string {
	return uuid.New().String()
}
