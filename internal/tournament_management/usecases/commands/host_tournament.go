package commands

import (
	"context"
	"errors"
	"time"

	"tournaments_backend/internal/tournament_management/domain"

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
) error {
	id := h.tournamentRepo.NextID(ctx)
	host, err := h.hostRepo.GetByUserID(ctx, hostUserID)
	if err != nil {
		return errors.New("")
	}

	t := domain.CreateTournament(id, host, title, date)

	if err := h.tournamentRepo.Save(ctx, t); err != nil {
		return err
	}

	return nil
}
