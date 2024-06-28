package commands

import (
	"context"
	"fmt"
	"time"

	"tournaments_backend/internal/domain"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
)

type HostTournamentHandler struct {
	eventBus       *cqrs.EventBus
	hostRepo       domain.HostRepository
	tournamentRepo domain.TournamentRepository
}

func NewHostTournamentHandler(
	eventBus *cqrs.EventBus,
	hostRepo domain.HostRepository,
	tournamentRepo domain.TournamentRepository,
) *HostTournamentHandler {
	return &HostTournamentHandler{
		eventBus:       eventBus,
		hostRepo:       hostRepo,
		tournamentRepo: tournamentRepo,
	}
}

func (h *HostTournamentHandler) Execute(
	ctx context.Context,
	hostUserID string,
	title string,
	date time.Time,
) (string, error) {
	id := h.tournamentRepo.NextID(ctx)
	host, err := h.hostRepo.GetByUserID(ctx, hostUserID)
	if err != nil {
		return "", fmt.Errorf("error hosting tournament: %w", err)
	}

	tournament := domain.CreateTournament(id, host, title, date)

	if err := h.tournamentRepo.Save(ctx, tournament); err != nil {
		return "", fmt.Errorf("error hosting tournament: %w", err)
	}

	h.eventBus.Publish(ctx, &domain.TournamentHostedEvent{
		ID:      tournament.ID,
		HostID:  tournament.HostID,
		Title:   tournament.Title,
		Players: tournament.Players,
		Date:    tournament.Date,
		State:   tournament.State,
	})

	return id, nil
}
