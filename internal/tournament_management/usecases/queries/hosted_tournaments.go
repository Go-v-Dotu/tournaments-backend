package queries

import (
	"context"
	"errors"

	"tournaments_backend/internal/tournament_management/domain"
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

func (h *HostedTournamentsHandler) Execute(ctx context.Context, hostUserID string) (Tournaments, error) {
	host, err := h.hostRepo.GetByUserID(ctx, hostUserID)
	if err != nil {
		return nil, errors.New("")
	}

	tournaments, err := h.tournamentQueryService.GetByHostID(ctx, host.ID)
	if err != nil {
		return nil, errors.New("")
	}

	return tournaments, nil
}
