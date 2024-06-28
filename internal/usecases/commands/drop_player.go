package commands

import (
	"context"
	"fmt"

	"tournaments_backend/internal/domain"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
)

type DropPlayerHandler struct {
	eventBus       *cqrs.EventBus
	hostRepo       domain.HostRepository
	playerRepo     domain.PlayerRepository
	tournamentRepo domain.TournamentRepository
}

func NewDropPlayerHandler(
	eventBus *cqrs.EventBus,
	hostRepo domain.HostRepository,
	playerRepo domain.PlayerRepository,
	tournamentRepo domain.TournamentRepository,
) *DropPlayerHandler {
	return &DropPlayerHandler{
		eventBus:       eventBus,
		hostRepo:       hostRepo,
		playerRepo:     playerRepo,
		tournamentRepo: tournamentRepo,
	}
}

func (h *DropPlayerHandler) Execute(
	ctx context.Context,
	hostUserID string,
	playerID string,
	tournamentID string,
) error {
	tournament, err := h.tournamentRepo.Get(ctx, tournamentID)
	if err != nil {
		return fmt.Errorf("error dropping player: %w", err)
	}

	host, err := h.hostRepo.GetByUserID(ctx, hostUserID)
	if err != nil {
		return fmt.Errorf("error dropping player: %w", err)
	}

	player, err := h.playerRepo.Get(ctx, playerID)
	if err != nil {
		return fmt.Errorf("error dropping player: %w", err)
	}

	if err := tournament.DropPlayer(host, player); err != nil {
		return fmt.Errorf("error dropping player: %w", err)
	}

	if err := h.tournamentRepo.Save(ctx, tournament); err != nil {
		return fmt.Errorf("error dropping player: %w", err)
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
