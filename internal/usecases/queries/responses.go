package queries

import "time"

type Tournament struct {
	ID           string
	Title        string
	Host         *Host
	Date         time.Time
	TotalPlayers int
}

type Host struct {
	ID       string
	UserID   string
	Username string
}

type Player struct {
	ID       string
	UserID   string
	Username string
	Dropped  bool
}
