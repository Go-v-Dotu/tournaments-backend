package api

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "tournaments_backend/api"
	"tournaments_backend/internal/tournament_management/application"
	"tournaments_backend/internal/tournament_management/infrastructure/api/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"
)

// Server
// @title			Tournament Management Service
// @version		1.0.0
// @description	Service for managing lifecycle of the tournaments
//
// @host			127.0.0.1:30001
// @BasePath		/api/v1
type Server struct {
	srv *echo.Echo
}

func NewServer(app *application.App) *Server {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.Use(middleware.Logger())

	e.GET("/docs/*", echoSwagger.WrapHandler)

	routes.Make(e.Group("/api/v1"), app)
	return &Server{srv: e}
}

func (s *Server) Start() {
	go func() {
		listener := make(chan os.Signal, 1)
		signal.Notify(listener, os.Interrupt, syscall.SIGTERM)
		// Listen on application shutdown signals.
		s.srv.Logger.Info("Received a shutdown signal:", <-listener)

		// Shutdown HTTP server.
		if err := s.srv.Shutdown(context.Background()); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.srv.Logger.Fatal(err)
		}
	}()

	// Start HTTP server.
	if err := s.srv.Start(":30001"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		s.srv.Logger.Fatal(err)
	}
}
