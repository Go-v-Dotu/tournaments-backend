package models

import "time"

type Tournament struct {
	ID       string
	HostID   string
	Title    string
	Settings *Settings
	Players  []*EnrolledPlayer
	Date     time.Time
	State    TournamentState
}

type Settings struct {
	MaxPlayers int
}

type EnrolledPlayer struct {
	PlayerID      string
	Dropped       bool
	HasCommanders bool
}

type TournamentState int
