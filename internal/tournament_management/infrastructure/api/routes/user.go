package routes

import (
	"net/http"

	"tournaments_backend/internal/tournament_management/infrastructure/api/controllers"
	"tournaments_backend/internal/tournament_management/infrastructure/api/views"
	"tournaments_backend/internal/tournament_management/usecases"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	uc *usecases.UseCases
}

// AddUser godoc
//
//	@Summary		Add User
//	@Description	host tournament by authorized player
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string						true	"Authorization info"
//	@Param			tournament_info	body		controllers.TournamentInfo	true	"Tournament info"
//	@Success		200				{object}	views.HostTournamentResponse
//	@Router			/users [post]
func (h *userHandler) AddUser(c echo.Context) error {
	ctx := c.Request().Context()

	var req controllers.AddUserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, views.ErrorResponse{})
	}

	if err := h.uc.Commands.CreateUser.Execute(ctx, req.ID, req.Username); err != nil {
		return c.JSON(http.StatusInternalServerError, views.ErrorResponse{})
	}

	resp := &views.AddUserResponse{}

	return c.JSON(http.StatusOK, resp)
}

func makeUserRoutes(e *echo.Group, uc *usecases.UseCases) {
	h := userHandler{uc: uc}

	e.POST("/users", h.AddUser)
}
