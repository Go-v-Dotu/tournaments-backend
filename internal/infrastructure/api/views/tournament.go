package views

import "time"

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

type SelfEnrollResponse struct{}

type GetPlayersResponse struct {
	Players []*Player `json:"players"`
}

type EnrollPlayerResponse struct{}

type EnrollGuestPlayerResponse struct {
	ID string `json:"id"`
}

type DropPlayerResponse struct{}

type RecoverPlayerResponse struct{}

type Tournament struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	Date         time.Time `json:"date"`
	TotalPlayers int       `json:"total_players"`
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
