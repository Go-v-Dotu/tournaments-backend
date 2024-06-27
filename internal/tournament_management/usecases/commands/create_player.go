package commands

import (
	"context"
	"errors"

	"tournaments_backend/internal/tournament_management/domain"
)

type CreatePlayerHandler struct {
	playerRepo domain.PlayerRepository
}

func NewCreatePlayerHandler(playerRepo domain.PlayerRepository) *CreatePlayerHandler {
	return &CreatePlayerHandler{playerRepo: playerRepo}
}

func (h *CreatePlayerHandler) Execute(ctx context.Context, userID string, username string) (string, error) {
	id := h.playerRepo.NextID(ctx)

	player := domain.NewPlayer(id, userID, username)

	if err := h.playerRepo.Save(ctx, player); err != nil {
		return "", errors.New("")
	}

	return id, nil
}
