package queries

import "time"

type Tournaments []*Tournament

type Tournament struct {
	ID           string
	Title        string
	Date         time.Time
	TotalPlayers int
}
