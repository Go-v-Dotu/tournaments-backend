package commands

import (
	"context"
	"errors"

	"tournaments_backend/internal/tournament_management/domain"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
)

type EnrollGuestPlayerHandler struct {
	eventBus       *cqrs.EventBus
	hostRepo       domain.HostRepository
	playerRepo     domain.PlayerRepository
	tournamentRepo domain.TournamentRepository
}

func NewEnrollGuestPlayerHandler(
	eventBus *cqrs.EventBus,
	hostRepo domain.HostRepository,
	playerRepo domain.PlayerRepository,
	tournamentRepo domain.TournamentRepository,
) *EnrollGuestPlayerHandler {
	return &EnrollGuestPlayerHandler{
		eventBus:       eventBus,
		hostRepo:       hostRepo,
		playerRepo:     playerRepo,
		tournamentRepo: tournamentRepo,
	}
}

func (h *EnrollGuestPlayerHandler) Execute(
	ctx context.Context,
	hostUserID string,
	tournamentID string,
	username string,
) (string, error) {
	tournament, err := h.tournamentRepo.Get(ctx, tournamentID)
	if err != nil {
		return "", errors.New("")
	}

	host, err := h.hostRepo.GetByUserID(ctx, hostUserID)
	if err != nil {
		return "", errors.New("")
	}

	playerID := h.playerRepo.NextID(ctx)
	player := domain.CreateGuestPlayer(playerID, username)
	if err := h.playerRepo.Save(ctx, player); err != nil {
		return "", errors.New("")
	}

	if err := tournament.EnrollPlayer(host, player); err != nil {
		return "", errors.New("")
	}

	if err := h.tournamentRepo.Save(ctx, tournament); err != nil {
		return "", err
	}

	return playerID, nil
}
