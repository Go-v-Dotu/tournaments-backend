package views

import "time"

type HostedTournamentsResponse struct {
	Tournaments []*TournamentPreview `json:"tournaments"`
}

type TournamentPreview struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	Date         time.Time `json:"date"`
	TotalPlayers int       `json:"total_players"`
}
