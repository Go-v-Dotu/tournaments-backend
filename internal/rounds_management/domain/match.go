package domain

type Match struct {
	ID      string
	Table   int
	Players []*Player
	//*Result
}

func NewMatch(id string, table int, players []*Player) *Match {
	return &Match{
		ID:      id,
		Table:   table,
		Players: players,
	}
}

//func (m *Match) SubmitResult(result *Result) {
//	m.Result = result
//}
