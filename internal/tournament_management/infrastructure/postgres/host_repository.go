package postgres

import (
	"context"
	"errors"

	"tournaments_backend/internal/tournament_management/domain"
	"tournaments_backend/internal/tournament_management/infrastructure/postgres/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type hostRepository struct {
	client Client
}

var (
	_ domain.HostRepository = (*hostRepository)(nil)
)

func NewHostRepository(client Client) domain.HostRepository {
	return &hostRepository{client: client}
}

func (r *hostRepository) GetByUserID(ctx context.Context, userID string) (*domain.Host, error) {
	query := "SELECT id, user_id, username FROM hosts WHERE user_id = $1"

	var hostModel models.Host
	if err := r.client.QueryRow(ctx, query, userID).Scan(
		&hostModel.ID,
		&hostModel.UserID,
		&hostModel.Username,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, errors.New("")
	}

	return hostModel.ToEntity(), nil
}

func (r *hostRepository) Save(ctx context.Context, host *domain.Host) error {
	query := "INSERT INTO hosts (id, user_id, username) VALUES ($1, $2, $3)"

	hostModel := models.NewHost(host)
	if _, err := r.client.Exec(ctx, query, hostModel.ID, hostModel.UserID, host.Username); err != nil {
		return errors.New("")
	}

	return nil
}

func (r *hostRepository) Delete(ctx context.Context, host *domain.Host) error {
	return nil
}

func (r *hostRepository) NextID(_ context.Context) string {
	return uuid.New().String()
}
