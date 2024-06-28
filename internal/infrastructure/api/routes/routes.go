package routes

import (
	"github.com/sirupsen/logrus"
	"tournaments_backend/internal/application"

	"github.com/labstack/echo/v4"
)

func Make(e *echo.Group, app *application.App, logger *logrus.Logger) {
	makeTournamentRoutes(e, app.UseCases, logger)
	makeUserRoutes(e, app.UseCases)
}
