package commands

import (
	"context"
	"fmt"

	"tournaments_backend/internal/domain"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
)

type RecoverPlayerHandler struct {
	eventBus       *cqrs.EventBus
	hostRepo       domain.HostRepository
	playerRepo     domain.PlayerRepository
	tournamentRepo domain.TournamentRepository
}

func NewRecoverPlayerHandler(
	eventBus *cqrs.EventBus,
	hostRepo domain.HostRepository,
	playerRepo domain.PlayerRepository,
	tournamentRepo domain.TournamentRepository,
) *RecoverPlayerHandler {
	return &RecoverPlayerHandler{
		eventBus:       eventBus,
		hostRepo:       hostRepo,
		playerRepo:     playerRepo,
		tournamentRepo: tournamentRepo,
	}
}

func (h *RecoverPlayerHandler) Execute(
	ctx context.Context,
	hostUserID string,
	playerID string,
	tournamentID string,
) error {
	tournament, err := h.tournamentRepo.Get(ctx, tournamentID)
	if err != nil {
		return fmt.Errorf("error recovering player: %w", err)
	}

	host, err := h.hostRepo.GetByUserID(ctx, hostUserID)
	if err != nil {
		return fmt.Errorf("error recovering player: %w", err)
	}

	player, err := h.playerRepo.Get(ctx, playerID)
	if err != nil {
		return fmt.Errorf("error recovering player: %w", err)
	}

	if err := tournament.RecoverPlayer(host, player); err != nil {
		return fmt.Errorf("error recovering player: %w", err)
	}

	if err := h.tournamentRepo.Save(ctx, tournament); err != nil {
		return fmt.Errorf("error recovering player: %w", err)
	}

	h.eventBus.Publish(ctx, &domain.TournamentUpdatedEvent{
		ID:      tournament.ID,
		HostID:  tournament.HostID,
		Title:   tournament.Title,
		Players: tournament.Players,
		Date:    tournament.Date,
		State:   tournament.State,
	})

	return nil
}
