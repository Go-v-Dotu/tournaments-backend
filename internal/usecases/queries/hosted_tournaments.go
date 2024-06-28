package queries

import (
	"context"
	"fmt"

	"tournaments_backend/internal/domain"
)

type HostedTournamentsHandler struct {
	hostRepo               domain.HostRepository
	tournamentQueryService TournamentQueryService
}

func NewHostedTournamentsHandler(
	hostRepo domain.HostRepository,
	tournamentQueryService TournamentQueryService,
) *HostedTournamentsHandler {
	return &HostedTournamentsHandler{
		hostRepo:               hostRepo,
		tournamentQueryService: tournamentQueryService,
	}
}

func (h *HostedTournamentsHandler) Execute(ctx context.Context, hostUserID string) ([]*Tournament, error) {
	host, err := h.hostRepo.GetByUserID(ctx, hostUserID)
	if err != nil {
		return nil, fmt.Errorf("error getting hosted tournaments: %w", err)
	}

	tournaments, err := h.tournamentQueryService.GetByHostID(ctx, host.ID)
	if err != nil {
		return nil, fmt.Errorf("error getting hosted tournaments: %w", err)
	}

	return tournaments, nil
}
