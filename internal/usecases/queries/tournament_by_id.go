package queries

import (
	"context"
	"fmt"

	"tournaments_backend/internal/domain"
)

type TournamentByIDHandler struct {
	hostRepo               domain.HostRepository
	tournamentRepo         domain.TournamentRepository
	tournamentQueryService TournamentQueryService
}

func NewTournamentByIDHandler(
	hostRepo domain.HostRepository,
	tournamentRepo domain.TournamentRepository,
	tournamentQueryService TournamentQueryService,
) *TournamentByIDHandler {
	return &TournamentByIDHandler{
		hostRepo:               hostRepo,
		tournamentRepo:         tournamentRepo,
		tournamentQueryService: tournamentQueryService,
	}
}

func (h *TournamentByIDHandler) Execute(ctx context.Context, hostUserID string, id string) (*Tournament, error) {
	host, err := h.hostRepo.GetByUserID(ctx, hostUserID)
	if err != nil {
		return nil, fmt.Errorf("error getting tournament by id: %w", err)
	}

	tournament, err := h.tournamentRepo.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error getting tournament by id: %w", err)
	}

	if !tournament.IsHostedBy(host) {
		return nil, fmt.Errorf("error getting tournament by id: you aren't host of the tournament")
	}

	tournamentResp, err := h.tournamentQueryService.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error getting tournament by id: %w", err)
	}

	return tournamentResp, nil
}
