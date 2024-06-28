package commands

import (
	"context"
	"fmt"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"

	"tournaments_backend/internal/domain"
)

type CreateHostHandler struct {
	eventBus *cqrs.EventBus
	hostRepo domain.HostRepository
}

func NewCreateHostHandler(eventBus *cqrs.EventBus, hostRepo domain.HostRepository) *CreateHostHandler {
	return &CreateHostHandler{eventBus: eventBus, hostRepo: hostRepo}
}

func (h *CreateHostHandler) Execute(ctx context.Context, userID string, username string) (string, error) {
	id := h.hostRepo.NextID(ctx)

	host := domain.NewHost(id, userID, username)

	if err := h.hostRepo.Save(ctx, host); err != nil {
		return "", fmt.Errorf("error creating tournament: %w", err)
	}

	h.eventBus.Publish(ctx, &domain.HostCreatedEvent{
		ID:       host.ID,
		UserID:   host.UserID,
		Username: host.Username,
	})

	return id, nil
}
