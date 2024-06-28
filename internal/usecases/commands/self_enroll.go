package commands

import (
	"context"
	"fmt"

	"tournaments_backend/internal/domain"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
)

type SelfEnrollHandler struct {
	eventBus       *cqrs.EventBus
	playerRepo     domain.PlayerRepository
	tournamentRepo domain.TournamentRepository
}

func NewSelfEnrollHandler(
	eventBus *cqrs.EventBus,
	playerRepo domain.PlayerRepository,
	tournamentRepo domain.TournamentRepository,
) *SelfEnrollHandler {
	return &SelfEnrollHandler{
		eventBus:       eventBus,
		playerRepo:     playerRepo,
		tournamentRepo: tournamentRepo,
	}
}

func (h *SelfEnrollHandler) Execute(ctx context.Context, playerUserID string, tournamentID string) error {
	tournament, err := h.tournamentRepo.Get(ctx, tournamentID)
	if err != nil {
		return fmt.Errorf("error self enrolling: %w", err)
	}

	player, err := h.playerRepo.GetByUserID(ctx, playerUserID)
	if err != nil {
		return fmt.Errorf("error self enrolling: %w", err)
	}

	if err := tournament.SelfEnroll(player); err != nil {
		return fmt.Errorf("error self enrolling: %w", err)
	}

	if err := h.tournamentRepo.Save(ctx, tournament); err != nil {
		return fmt.Errorf("error self enrolling: %w", err)
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
