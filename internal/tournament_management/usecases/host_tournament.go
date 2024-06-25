package usecases

import (
	"context"

	"tournaments_backend/internal/tournament_management/domain"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
)

type HostTournamentHandler struct {
	eventBus       *cqrs.EventBus
	tournamentRepo domain.TournamentRepository
}

func (h *HostTournamentHandler) Execute(ctx context.Context) {

}
