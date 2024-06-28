package domain

import "time"

type TournamentHostedEvent struct {
	ID      string
	HostID  string
	Title   string
	Players []*EnrolledPlayer
	Date    time.Time
	State   TournamentState
}

type TournamentUpdatedEvent struct {
	ID      string
	HostID  string
	Title   string
	Players []*EnrolledPlayer
	Date    time.Time
	State   TournamentState
}
