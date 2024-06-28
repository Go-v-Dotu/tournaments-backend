package routes

import (
	"net/http"

	"tournaments_backend/internal/infrastructure/api/controllers"
	"tournaments_backend/internal/infrastructure/api/views"
	"tournaments_backend/internal/usecases"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	uc *usecases.UseCases
}

// AddUser godoc
//
//	@Summary		Add User
//	@Description	notify that user was registered
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			user_info	body		controllers.UserInfo	true	"User info"
//	@Success		200			{object}	views.AddUserResponse
//	@Router			/users [post]
func (h *userHandler) AddUser(c echo.Context) error {
	ctx := c.Request().Context()

	var req controllers.AddUserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, views.ErrorResponse{})
	}

	if err := h.uc.Commands.CreateUserHandler.Execute(ctx, req.ID, req.Username); err != nil {
		return c.JSON(http.StatusInternalServerError, views.ErrorResponse{})
	}

	resp := &views.AddUserResponse{}

	return c.JSON(http.StatusOK, resp)
}

func makeUserRoutes(e *echo.Group, uc *usecases.UseCases) {
	h := userHandler{uc: uc}

	e.POST("/users", h.AddUser)
}
