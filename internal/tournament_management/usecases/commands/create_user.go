package commands

import (
	"context"
	"errors"

	"tournaments_backend/internal/tournament_management/domain"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
)

type CreateUserHandler struct {
	eventBus *cqrs.EventBus
}

func NewCreateUserHandler(eventBus *cqrs.EventBus) *CreateUserHandler {
	return &CreateUserHandler{eventBus: eventBus}
}

func (h *CreateUserHandler) Execute(ctx context.Context, id string, username string) error {
	err := h.eventBus.Publish(ctx, &domain.UserRegisteredEvent{
		ID:       id,
		Username: username,
	})
	if err != nil {
		return errors.New("")
	}

	return nil
}
