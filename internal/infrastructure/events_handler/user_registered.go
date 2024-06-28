package events_handler

import (
	"context"

	"tournaments_backend/internal/domain"
	"tournaments_backend/internal/usecases"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
)

type UserRegisteredHandler struct {
	uc *usecases.UseCases
}

var (
	_ cqrs.GroupEventHandler = (*UserRegisteredHandler)(nil)
)

func NewUserRegisteredHandler(uc *usecases.UseCases) *UserRegisteredHandler {
	return &UserRegisteredHandler{uc: uc}
}

func (h *UserRegisteredHandler) NewEvent() any {
	return &domain.UserRegisteredEvent{}
}

func (h *UserRegisteredHandler) Handle(ctx context.Context, e any) error {
	event := e.(*domain.UserRegisteredEvent)

	h.uc.Commands.CreateHostHandler.Execute(ctx, event.ID, event.Username)
	h.uc.Commands.CreatePlayerHandler.Execute(ctx, event.ID, event.Username)
	return nil
}
