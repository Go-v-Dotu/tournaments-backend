package commands

import (
	"context"
	"errors"

	"tournaments_backend/internal/tournament_management/domain"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
)

type EnrollPlayerHandler struct {
	eventBus       *cqrs.EventBus
	hostRepo       domain.HostRepository
	playerRepo     domain.PlayerRepository
	tournamentRepo domain.TournamentRepository
}

func NewEnrollPlayerHandler(
	eventBus *cqrs.EventBus,
	hostRepo domain.HostRepository,
	playerRepo domain.PlayerRepository,
	tournamentRepo domain.TournamentRepository,
) *EnrollPlayerHandler {
	return &EnrollPlayerHandler{
		eventBus:       eventBus,
		hostRepo:       hostRepo,
		playerRepo:     playerRepo,
		tournamentRepo: tournamentRepo,
	}
}

func (h *EnrollPlayerHandler) Execute(
	ctx context.Context,
	hostUserID string,
	playerUserID string,
	tournamentID string,
) error {
	tournament, err := h.tournamentRepo.Get(ctx, tournamentID)
	if err != nil {
		return errors.New("")
	}

	host, err := h.hostRepo.GetByUserID(ctx, hostUserID)
	if err != nil {
		return errors.New("")
	}

	player, err := h.playerRepo.GetByUserID(ctx, playerUserID)
	if err != nil {
		return errors.New("")
	}

	if err := tournament.EnrollPlayer(host, player); err != nil {
		return errors.New("")
	}

	if err := h.tournamentRepo.Save(ctx, tournament); err != nil {
		return err
	}

	return nil
}
