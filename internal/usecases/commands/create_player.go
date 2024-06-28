package commands

import (
	"context"
	"fmt"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"

	"tournaments_backend/internal/domain"
)

type CreatePlayerHandler struct {
	eventBus   *cqrs.EventBus
	playerRepo domain.PlayerRepository
}

func NewCreatePlayerHandler(eventBus *cqrs.EventBus, playerRepo domain.PlayerRepository) *CreatePlayerHandler {
	return &CreatePlayerHandler{eventBus: eventBus, playerRepo: playerRepo}
}

func (h *CreatePlayerHandler) Execute(ctx context.Context, userID string, username string) (string, error) {
	id := h.playerRepo.NextID(ctx)

	player := domain.NewPlayer(id, userID, username)

	if err := h.playerRepo.Save(ctx, player); err != nil {
		return "", fmt.Errorf("error creating player: %w", err)
	}

	h.eventBus.Publish(ctx, &domain.PlayerCreatedEvent{
		ID:       player.ID,
		UserID:   player.UserID,
		Username: player.Username,
	})

	return id, nil
}
