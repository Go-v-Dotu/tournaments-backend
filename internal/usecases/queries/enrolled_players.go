package queries

import (
	"context"
	"fmt"

	"tournaments_backend/internal/domain"
)

type EnrolledPlayersHandler struct {
	hostRepo           domain.HostRepository
	tournamentRepo     domain.TournamentRepository
	playerQueryService PlayerQueryService
}

func NewEnrolledPlayersHandler(
	hostRepo domain.HostRepository,
	tournamentRepo domain.TournamentRepository,
	playerQueryService PlayerQueryService,
) *EnrolledPlayersHandler {
	return &EnrolledPlayersHandler{
		hostRepo:           hostRepo,
		tournamentRepo:     tournamentRepo,
		playerQueryService: playerQueryService,
	}
}

func (h *EnrolledPlayersHandler) Execute(ctx context.Context, hostUserID string, tournamentID string) ([]*Player, error) {
	host, err := h.hostRepo.GetByUserID(ctx, hostUserID)
	if err != nil {
		return nil, fmt.Errorf("error getting enrolled players: %w", err)
	}

	tournament, err := h.tournamentRepo.Get(ctx, tournamentID)
	if err != nil {
		return nil, fmt.Errorf("error getting enrolled players: %w", err)
	}

	if !tournament.IsHostedBy(host) {
		return nil, fmt.Errorf("error getting enrolled players: you aren't host of the tournament")
	}

	enrolledPlayers, err := h.playerQueryService.GetByTournamentID(ctx, tournament.ID)
	if err != nil {
		return nil, fmt.Errorf("error getting enrolled players: %w", err)
	}

	return enrolledPlayers, nil
}
