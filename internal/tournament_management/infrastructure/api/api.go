package api

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"tournaments_backend/internal/tournament_management/application"

	"tournaments_backend/internal/tournament_management/infrastructure/api/routes"

	"github.com/labstack/echo/v4"
)

type Server struct {
	srv *echo.Echo
}

func NewServer(app *application.App) *Server {
	e := echo.New()
	routes.Make(e.Group("/api/v1"), app)
	return &Server{srv: e}
}

func (s *Server) Start() {
	go func() {
		listener := make(chan os.Signal, 1)
		signal.Notify(listener, os.Interrupt, syscall.SIGTERM)
		// Listen on application shutdown signals.
		log.Println("Received a shutdown signal:", <-listener)

		// Shutdown HTTP server.
		if err := s.srv.Shutdown(context.Background()); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	// Start HTTP server.
	if err := s.srv.Start(":30001"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
