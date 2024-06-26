package usecases

import (
	"context"

	"tournaments_backend/internal/tournament_participation/domain"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
)

type DropTournamentHandler struct {
	eventBus       *cqrs.EventBus
	tournamentRepo domain.TournamentRepository
	playerRepo     domain.PlayerRepository
}

func (h *DropTournamentHandler) Execute(ctx context.Context, playerID string, tournamentID string) {
	//tournament := h.tournamentRepo.Get(ctx, tournamentID)
	//player := h.playerRepo.Get(ctx, playerID)

}
