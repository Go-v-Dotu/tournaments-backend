package routes

import (
	"github.com/sirupsen/logrus"
	"net/http"

	"tournaments_backend/internal/infrastructure/api/controllers"
	"tournaments_backend/internal/infrastructure/api/views"
	"tournaments_backend/internal/usecases"

	"github.com/labstack/echo/v4"
)

type tournamentHandler struct {
	uc     *usecases.UseCases
	logger *logrus.Logger
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
	errMsg := "Hosting tournaments got error"

	userID := c.Request().Header.Get("Authorization")
	if userID == "" {
		h.logger.Infof("%s: %s", errMsg, "user not authorized")
		return c.JSON(http.StatusUnauthorized, views.ErrorResponse{})
	}

	tournaments, err := h.uc.Queries.HostedTournamentsHandler.Execute(ctx, userID)
	if err != nil {
		h.logger.Infof("%s: %s", errMsg, err.Error())
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
	errMsg := "Hosting tournament got error"

	userID := c.Request().Header.Get("Authorization")
	if userID == "" {
		h.logger.Infof("%s: %s", errMsg, "user not authorized")
		return c.JSON(http.StatusUnauthorized, views.ErrorResponse{})
	}

	var req controllers.HostTournamentRequest
	if err := c.Bind(&req); err != nil {
		h.logger.Infof("%s: %s", errMsg, err.Error())
		return c.JSON(http.StatusBadRequest, views.ErrorResponse{})
	}

	id, err := h.uc.Commands.HostTournamentHandler.Execute(ctx, userID, req.Title, req.Date)
	if err != nil {
		h.logger.Infof("%s: %s", errMsg, err.Error())
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
	errMsg := "Getting tournament got error"

	userID := c.Request().Header.Get("Authorization")
	if userID == "" {
		h.logger.Infof("%s: %s", errMsg, "user not authorized")
		return c.JSON(http.StatusUnauthorized, views.ErrorResponse{})
	}

	var req controllers.GetTournamentRequest
	if err := c.Bind(&req); err != nil {
		h.logger.Infof("%s: %s", errMsg, err.Error())
		return c.JSON(http.StatusBadRequest, views.ErrorResponse{})
	}

	tournament, err := h.uc.Queries.TournamentByIDHandler.Execute(ctx, userID, req.ID)
	if err != nil {
		h.logger.Infof("%s: %s", errMsg, err.Error())
		return c.JSON(http.StatusInternalServerError, views.ErrorResponse{})
	}

	tt := &views.Tournament{
		ID:           tournament.ID,
		Title:        tournament.Title,
		Date:         tournament.Date,
		TotalPlayers: tournament.TotalPlayers,
	}

	resp := &views.GetTournamentResponse{Tournament: tt}

	return c.JSON(http.StatusOK, resp)
}

// SelfEnroll godoc
//
//	@Summary		Self Enroll
//	@Description	enroll self to the tournament
//	@Tags			tournaments
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	true	"Authorization info"
//	@Param			id				path		string	true	"ID of the tournament"
//	@Success		200				{object}	views.SelfEnrollResponse
//	@Router			/tournaments/{id}/enroll [post]
func (h *tournamentHandler) SelfEnroll(c echo.Context) error {
	ctx := c.Request().Context()
	errMsg := "Enrolling self got error"

	userID := c.Request().Header.Get("Authorization")
	if userID == "" {
		h.logger.Infof("%s: %s", errMsg, "user not authorized")
		return c.JSON(http.StatusUnauthorized, views.ErrorResponse{})
	}

	var req controllers.SelfEnrollRequest
	if err := c.Bind(&req); err != nil {
		h.logger.Infof("%s: %s", errMsg, err.Error())
		return c.JSON(http.StatusBadRequest, views.ErrorResponse{})
	}

	if err := h.uc.Commands.SelfEnrollHandler.Execute(ctx, userID, req.ID); err != nil {
		h.logger.Infof("%s: %s", errMsg, err.Error())
		return c.JSON(http.StatusInternalServerError, views.ErrorResponse{})
	}

	resp := &views.SelfEnrollResponse{}

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
	errMsg := "Getting players got error"

	userID := c.Request().Header.Get("Authorization")
	if userID == "" {
		h.logger.Infof("%s: %s", errMsg, "user not authorized")
		return c.JSON(http.StatusUnauthorized, views.ErrorResponse{})
	}

	var req controllers.GetPlayersRequest
	if err := c.Bind(&req); err != nil {
		h.logger.Infof("%s: %s", errMsg, err.Error())
		return c.JSON(http.StatusBadRequest, views.ErrorResponse{})
	}

	enrolledPlayers, err := h.uc.Queries.EnrolledPlayersHandler.Execute(ctx, userID, req.TournamentID)
	if err != nil {
		h.logger.Infof("%s: %s", errMsg, err.Error())
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
//	@Router			/tournaments/{id}/players [post]
func (h *tournamentHandler) EnrollGuestPlayer(c echo.Context) error {
	ctx := c.Request().Context()
	errMsg := "Enrolling guest player got error"

	userID := c.Request().Header.Get("Authorization")
	if userID == "" {
		h.logger.Infof("%s: %s", errMsg, "user not authorized")
		return c.JSON(http.StatusUnauthorized, views.ErrorResponse{})
	}

	var req controllers.EnrollGuestPlayerRequest
	if err := c.Bind(&req); err != nil {
		h.logger.Infof("%s: %s", errMsg, err.Error())
		return c.JSON(http.StatusBadRequest, views.ErrorResponse{})
	}

	id, err := h.uc.Commands.EnrollGuestPlayerHandler.Execute(ctx, userID, req.TournamentID, req.Username)
	if err != nil {
		h.logger.Infof("%s: %s", errMsg, err.Error())
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
	errMsg := "Enrolling player got error"

	userID := c.Request().Header.Get("Authorization")
	if userID == "" {
		h.logger.Infof("%s: %s", errMsg, "user not authorized")
		return c.JSON(http.StatusUnauthorized, views.ErrorResponse{})
	}

	var req controllers.EnrollPlayerRequest
	if err := c.Bind(&req); err != nil {
		h.logger.Infof("%s: %s", errMsg, err.Error())
		return c.JSON(http.StatusBadRequest, views.ErrorResponse{})
	}

	if err := h.uc.Commands.EnrollPlayerHandler.Execute(ctx, userID, req.UserID, req.TournamentID); err != nil {
		h.logger.Infof("%s: %s", errMsg, err.Error())
		return c.JSON(http.StatusInternalServerError, views.ErrorResponse{})
	}

	resp := &views.EnrollPlayerResponse{}

	return c.JSON(http.StatusOK, resp)
}

// DropPlayer godoc
//
//	@Summary		Drop Player
//	@Description	drop a player
//	@Tags			tournaments,players
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	true	"Authorization info"
//	@Param			id				path		string	true	"ID of the tournament"
//	@Param			player_id		path		string	true	"ID of the player"
//	@Success		200				{object}	views.DropPlayerResponse
//	@Router			/tournaments/{id}/players/{player_id}/drop [post]
func (h *tournamentHandler) DropPlayer(c echo.Context) error {
	ctx := c.Request().Context()
	errMsg := "Dropping player got error"

	userID := c.Request().Header.Get("Authorization")
	if userID == "" {
		h.logger.Infof("%s: %s", errMsg, "user not authorized")
		return c.JSON(http.StatusUnauthorized, views.ErrorResponse{})
	}

	var req controllers.DropPlayerRequest
	if err := c.Bind(&req); err != nil {
		h.logger.Infof("%s: %s", errMsg, err.Error())
		return c.JSON(http.StatusBadRequest, views.ErrorResponse{})
	}

	if err := h.uc.Commands.DropPlayerHandler.Execute(ctx, userID, req.PlayerID, req.TournamentID); err != nil {
		h.logger.Infof("%s: %s", errMsg, err.Error())
		return c.JSON(http.StatusInternalServerError, views.ErrorResponse{})
	}

	resp := &views.DropPlayerResponse{}

	return c.JSON(http.StatusOK, resp)
}

// RecoverPlayer godoc
//
//	@Summary		Recover Player
//	@Description	recover a player
//	@Tags			tournaments,players
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	true	"Authorization info"
//	@Param			id				path		string	true	"ID of the tournament"
//	@Param			player_id		path		string	true	"ID of the player"
//	@Success		200				{object}	views.DropPlayerResponse
//	@Router			/tournaments/{id}/players/{player_id}/recover [post]
func (h *tournamentHandler) RecoverPlayer(c echo.Context) error {
	ctx := c.Request().Context()
	errMsg := "Recovering player got error"

	userID := c.Request().Header.Get("Authorization")
	if userID == "" {
		h.logger.Infof("%s: %s", errMsg, "user not authorized")
		return c.JSON(http.StatusUnauthorized, views.ErrorResponse{})
	}

	var req controllers.DropPlayerRequest
	if err := c.Bind(&req); err != nil {
		h.logger.Infof("%s: %s", errMsg, err.Error())
		return c.JSON(http.StatusBadRequest, views.ErrorResponse{})
	}

	if err := h.uc.Commands.RecoverPlayerHandler.Execute(ctx, userID, req.PlayerID, req.TournamentID); err != nil {
		h.logger.Infof("%s: %s", errMsg, err.Error())
		return c.JSON(http.StatusInternalServerError, views.ErrorResponse{})
	}

	resp := &views.RecoverPlayerResponse{}

	return c.JSON(http.StatusOK, resp)
}

func makeTournamentRoutes(e *echo.Group, uc *usecases.UseCases, logger *logrus.Logger) {
	h := tournamentHandler{uc: uc, logger: logger}

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
			e.POST("/enroll", h.SelfEnroll)

			{
				e := e.Group("/players")
				e.GET("", h.GetPlayers)
				e.POST("", h.EnrollGuestPlayer)
				e.PUT("/:userID", h.EnrollPlayer)
				e.POST("/:playerID/drop", h.DropPlayer)
				e.POST("/:playerID/recover", h.RecoverPlayer)
			}
		}
	}
}
