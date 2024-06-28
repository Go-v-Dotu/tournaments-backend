package commands

import (
	"context"
	"fmt"

	"tournaments_backend/internal/domain"

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
		return fmt.Errorf("error enrolling player: %w", err)
	}

	host, err := h.hostRepo.GetByUserID(ctx, hostUserID)
	if err != nil {
		return fmt.Errorf("error enrolling player: %w", err)
	}

	player, err := h.playerRepo.GetByUserID(ctx, playerUserID)
	if err != nil {
		return fmt.Errorf("error enrolling player: %w", err)
	}

	if err := tournament.EnrollPlayer(host, player); err != nil {
		return fmt.Errorf("error enrolling player: %w", err)
	}

	if err := h.tournamentRepo.Save(ctx, tournament); err != nil {
		return fmt.Errorf("error enrolling player: %w", err)
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
