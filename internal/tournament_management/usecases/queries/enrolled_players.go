package queries

import (
	"context"
	"errors"

	"tournaments_backend/internal/tournament_management/domain"
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
		return nil, errors.New("")
	}

	tournament, err := h.tournamentRepo.Get(ctx, tournamentID)
	if err != nil {
		return nil, errors.New("")
	}

	if !tournament.IsHostedBy(host) {
		return nil, errors.New("")
	}

	enrolledPlayers, err := h.playerQueryService.GetByTournamentID(ctx, tournament.ID)
	if err != nil {
		return nil, errors.New("")
	}

	return enrolledPlayers, nil
}
