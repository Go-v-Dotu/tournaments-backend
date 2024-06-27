package queries

import "time"

type Tournament struct {
	ID           string
	Title        string
	Date         time.Time
	TotalPlayers int
}

type Player struct {
	ID       string
	UserID   string
	Username string
	Dropped  bool
}
