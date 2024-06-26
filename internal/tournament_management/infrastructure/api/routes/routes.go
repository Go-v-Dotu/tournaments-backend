package routes

import (
	"tournaments_backend/internal/tournament_management/application"

	"github.com/labstack/echo/v4"
)

func Make(e *echo.Group, app *application.App) {
	makeTournamentRoutes(e, app.UseCases)
}
