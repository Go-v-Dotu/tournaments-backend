package routes

import (
	"net/http"

	"tournaments_backend/internal/tournament_management/infrastructure/api/controllers"
	"tournaments_backend/internal/tournament_management/infrastructure/api/views"
	"tournaments_backend/internal/tournament_management/usecases"

	"github.com/labstack/echo/v4"
)

type tournamentHandler struct {
	uc *usecases.UseCases
}

// HostedTournaments godoc
//
//	@Summary		Hosted Tournaments
//	@Description	get all tournaments hosted by authorized user
//	@Tags			tournaments
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	true	"Authorization info"
//	@Success		200				{object}	views.HostedTournamentsResponse
//	@Router			/user/tournaments [get]
func (h *tournamentHandler) HostedTournaments(c echo.Context) error {
	ctx := c.Request().Context()

	userID := c.Request().Header.Get("Authorization")
	if userID == "" {
		return c.JSON(http.StatusUnauthorized, views.ErrorResponse{})
	}

	tournaments, err := h.uc.Queries.HostedTournamentsHandler.Execute(ctx, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, views.ErrorResponse{})
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

	resp := &views.HostedTournamentsResponse{Tournaments: tt}

	return c.JSON(http.StatusOK, resp)
}

// HostTournament godoc
//
//	@Summary		Host Tournament
//	@Description	host tournament by authorized player
//	@Tags			tournaments
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string						true	"Authorization info"
//	@Param			tournament_info	body		controllers.TournamentInfo	true	"Tournament info"
//	@Success		200				{object}	views.HostTournamentResponse
//	@Router			/tournaments [post]
func (h *tournamentHandler) HostTournament(c echo.Context) error {
	ctx := c.Request().Context()

	userID := c.Request().Header.Get("Authorization")
	if userID == "" {
		return c.JSON(http.StatusUnauthorized, views.ErrorResponse{})
	}

	var req controllers.HostTournamentRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, views.ErrorResponse{})
	}

	id, err := h.uc.Commands.HostTournamentHandler.Execute(ctx, userID, req.Title, req.Date)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, views.ErrorResponse{})
	}

	resp := &views.HostTournamentResponse{ID: id}

	return c.JSON(http.StatusOK, resp)
}

// GetTournament godoc
//
//	@Summary		Get Tournament
//	@Description	get tournament
//	@Tags			tournaments
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	true	"Authorization info"
//	@Param			id				path		string	true	"ID of the tournament"
//	@Success		200				{object}	views.GetTournamentResponse
//	@Router			/tournaments/{id} [get]
func (h *tournamentHandler) GetTournament(c echo.Context) error {
	ctx := c.Request().Context()

	userID := c.Request().Header.Get("Authorization")
	if userID == "" {
		return c.JSON(http.StatusUnauthorized, views.ErrorResponse{})
	}

	var req controllers.GetTournamentRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, views.ErrorResponse{})
	}

	tournament, err := h.uc.Queries.TournamentByIDHandler.Execute(ctx, userID, req.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, views.ErrorResponse{})
	}

	tt := &views.Tournament{
		ID:    tournament.ID,
		Title: tournament.Title,
		Date:  tournament.Date,
	}

	resp := &views.GetTournamentResponse{Tournament: tt}

	return c.JSON(http.StatusOK, resp)
}

// GetPlayers godoc
//
//	@Summary		Get Players
//	@Description	get players for tournament
//	@Tags			tournaments,players
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	true	"Authorization info"
//	@Param			id				path		string	true	"ID of the tournament"
//	@Success		200				{object}	views.GetPlayersResponse
//	@Router			/tournaments/{id}/players [get]
func (h *tournamentHandler) GetPlayers(c echo.Context) error {
	ctx := c.Request().Context()

	userID := c.Request().Header.Get("Authorization")
	if userID == "" {
		return c.JSON(http.StatusUnauthorized, views.ErrorResponse{})
	}

	var req controllers.GetPlayersRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, views.ErrorResponse{})
	}

	enrolledPlayers, err := h.uc.Queries.EnrolledPlayersHandler.Execute(ctx, userID, req.TournamentID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, views.ErrorResponse{})
	}

	pp := make([]*views.Player, 0, len(enrolledPlayers))
	for _, enrolledPlayer := range enrolledPlayers {
		pp = append(pp, &views.Player{
			ID:       enrolledPlayer.ID,
			UserID:   enrolledPlayer.UserID,
			Username: enrolledPlayer.Username,
			Dropped:  enrolledPlayer.Dropped,
		})
	}

	resp := &views.GetPlayersResponse{Players: pp}

	return c.JSON(http.StatusOK, resp)
}

// EnrollGuestPlayer godoc
//
//	@Summary		Enroll Guest Player
//	@Description	enroll a player that isn't a registered user
//	@Tags			tournaments,players
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string						true	"Authorization info"
//	@Param			id				path		string						true	"ID of the tournament"
//	@Param			user_info		body		controllers.GuestUserInfo	true	"Guest info"
//	@Success		200				{object}	views.EnrollGuestPlayerResponse
//	@Router			/tournaments/{id}/players [get]
func (h *tournamentHandler) EnrollGuestPlayer(c echo.Context) error {
	ctx := c.Request().Context()

	userID := c.Request().Header.Get("Authorization")
	if userID == "" {
		return c.JSON(http.StatusUnauthorized, views.ErrorResponse{})
	}

	var req controllers.EnrollGuestPlayerRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, views.ErrorResponse{})
	}

	id, err := h.uc.Commands.EnrollGuestPlayerHandler.Execute(ctx, userID, req.TournamentID, req.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, views.ErrorResponse{})
	}

	resp := &views.EnrollGuestPlayerResponse{ID: id}

	return c.JSON(http.StatusOK, resp)
}

// EnrollPlayer godoc
//
//	@Summary		Enroll Player
//	@Description	enroll a player that is a registered user
//	@Tags			tournaments,players
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	true	"Authorization info"
//	@Param			id				path		string	true	"ID of the tournament"
//	@Param			user_id			path		string	true	"ID of the user to be added"
//	@Success		200				{object}	views.EnrollPlayerResponse
//	@Router			/tournaments/{id}/players/{user_id} [put]
func (h *tournamentHandler) EnrollPlayer(c echo.Context) error {
	ctx := c.Request().Context()

	userID := c.Request().Header.Get("Authorization")
	if userID == "" {
		return c.JSON(http.StatusUnauthorized, views.ErrorResponse{})
	}

	var req controllers.EnrollPlayerRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, views.ErrorResponse{})
	}

	if err := h.uc.Commands.EnrollPlayerHandler.Execute(ctx, userID, req.UserID, req.TournamentID); err != nil {
		return c.JSON(http.StatusInternalServerError, views.ErrorResponse{})
	}

	resp := &views.EnrollPlayerResponse{}

	return c.JSON(http.StatusOK, resp)
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

			{
				e := e.Group("/players")
				e.GET("", h.GetPlayers)
				e.POST("", h.EnrollGuestPlayer)
				e.PUT("/:userID", h.EnrollPlayer)
			}
		}
	}
}
