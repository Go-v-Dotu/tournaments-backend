package commands

import (
	"context"
	"errors"

	"tournaments_backend/internal/tournament_management/domain"
)

type CreateHostHandler struct {
	hostRepo domain.HostRepository
}

func NewCreateHostHandler(hostRepo domain.HostRepository) *CreateHostHandler {
	return &CreateHostHandler{hostRepo: hostRepo}
}

func (h *CreateHostHandler) Execute(ctx context.Context, userID string, username string) (string, error) {
	id := h.hostRepo.NextID(ctx)

	host := domain.NewHost(id, userID, username)

	if err := h.hostRepo.Save(ctx, host); err != nil {
		return "", errors.New("")
	}

	return id, nil
}
