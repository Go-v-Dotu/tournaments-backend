package routes

import (
	"net/http"
	"tournaments_backend/internal/tournament_management/infrastructure/api/views"

	"tournaments_backend/internal/tournament_management/infrastructure/api/controllers"
	"tournaments_backend/internal/tournament_management/usecases"

	"github.com/labstack/echo/v4"
)

type tournamentHandler struct {
	uc *usecases.UseCases
}

func (h *tournamentHandler) HostedTournaments(c echo.Context) error {
	ctx := c.Request().Context()

	userID := c.Request().Header.Get("Authorization")
	if userID == "" {
		return c.JSON(http.StatusUnauthorized, "")
	}

	tournaments, err := h.uc.Queries.HostedTournamentsHandler.Execute(ctx, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "")
	}

	tt := make([]*views.TournamentPreview, 0, len(tournaments))
	for _, tot := range tournaments {
		tt = append(tt, &views.TournamentPreview{
			ID:           tot.ID,
			Title:        tot.Title,
			Date:         tot.Date,
			TotalPlayers: tot.TotalPlayers,
		})
	}

	return c.JSON(http.StatusOK, &views.HostedTournamentsResponse{Tournaments: tt})
}

func (h *tournamentHandler) HostTournament(c echo.Context) error {
	ctx := c.Request().Context()

	userID := c.Request().Header.Get("Authorization")
	if userID == "" {
		return c.JSON(http.StatusUnauthorized, "")
	}

	var req controllers.HostTournamentRequest
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	if err := h.uc.Commands.HostTournamentHandler.Execute(ctx, userID, req.Title, req.Date); err != nil {
		return c.JSON(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, "")
}

func (h *tournamentHandler) GetTournament(c echo.Context) error {
	return nil
}

func (h *tournamentHandler) EnrollPlayer(c echo.Context) error {
	ctx := c.Request().Context()

	userID := c.Request().Header.Get("Authorization")
	if userID == "" {
		return c.JSON(http.StatusUnauthorized, "")
	}

	var req controllers.EnrollPlayerRequest
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	if err := h.uc.Commands.EnrollPlayerHandler.Execute(ctx, userID, req.UserID, req.TournamentID); err != nil {
		return c.JSON(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, "")
}

func makeTournamentRoutes(e *echo.Group, uc *usecases.UseCases) {
	h := tournamentHandler{uc: uc}

	{
		e := e.Group("/user/tournaments")
		e.GET("", h.HostedTournaments)
	}

	{
		e := e.Group("/tournaments")
		e.POST("", h.HostTournament)

		{
			e := e.Group("/:id")
			e.GET("", h.GetTournament)

			e.PUT("/players/:userID", h.EnrollPlayer)

			//e.PUT("/settings", h.GetTournaments)
		}
	}
}
