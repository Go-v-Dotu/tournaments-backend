package queries

import (
	"context"
	"errors"

	"tournaments_backend/internal/tournament_management/domain"
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
		return nil, errors.New("")
	}

	tournament, err := h.tournamentRepo.Get(ctx, id)
	if err != nil {
		return nil, errors.New("")
	}

	if !tournament.IsHostedBy(host) {
		return nil, errors.New("")
	}

	tournamentResp, err := h.tournamentQueryService.GetByID(ctx, id)
	if err != nil {
		return nil, errors.New("")
	}

	return tournamentResp, nil
}
