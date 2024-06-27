package views

import (
	"time"

	"tournaments_backend/internal/tournament_management/domain"
)

type ErrorResponse struct{}

type HostedTournamentsResponse struct {
	Tournaments []*TournamentPreview `json:"tournaments"`
}

type HostTournamentResponse struct {
	ID string `json:"id"`
}

type GetTournamentResponse struct {
	Tournament *Tournament `json:"tournament"`
}

type GetPlayersResponse struct {
	Players []*Player `json:"players"`
}

type EnrollPlayerResponse struct{}

type EnrollGuestPlayerResponse struct {
	ID string `json:"id"`
}

type Tournament struct {
	ID    string    `json:"id"`
	Title string    `json:"title"`
	Date  time.Time `json:"date"`
}

type TournamentPreview struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	Date         time.Time `json:"date"`
	TotalPlayers int       `json:"total_players"`
}

type Player struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Dropped  bool   `json:"dropped"`
}

func NewTournamentPreview(tournament *domain.Tournament) *TournamentPreview {
	//
	return nil
}
