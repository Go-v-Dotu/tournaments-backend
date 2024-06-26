package domain

type Tournament struct {
	ID      string
	HostID  string
	Title   string
	Players []*Player
	//Rounds  []*Round
	//Date    time.Time
	//State   TournamentState
}
