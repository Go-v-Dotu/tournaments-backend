package commands

import (
	"context"
	"fmt"

	"tournaments_backend/internal/domain"

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
		return "", fmt.Errorf("error enrolling guest player: %w", err)
	}

	host, err := h.hostRepo.GetByUserID(ctx, hostUserID)
	if err != nil {
		return "", fmt.Errorf("error enrolling guest player: %w", err)
	}

	playerID := h.playerRepo.NextID(ctx)
	player := domain.CreateGuestPlayer(playerID, username)
	if err := h.playerRepo.Save(ctx, player); err != nil {
		return "", fmt.Errorf("error enrolling guest player: %w", err)
	}

	h.eventBus.Publish(ctx, &domain.PlayerCreatedEvent{
		ID:       player.ID,
		UserID:   player.UserID,
		Username: player.Username,
	})

	if err := tournament.EnrollPlayer(host, player); err != nil {
		return "", fmt.Errorf("error enrolling guest player: %w", err)
	}

	if err := h.tournamentRepo.Save(ctx, tournament); err != nil {
		return "", fmt.Errorf("error enrolling guest player: %w", err)
	}

	h.eventBus.Publish(ctx, &domain.TournamentUpdatedEvent{
		ID:      tournament.ID,
		HostID:  tournament.HostID,
		Title:   tournament.Title,
		Players: tournament.Players,
		Date:    tournament.Date,
		State:   tournament.State,
	})

	return playerID, nil
}
