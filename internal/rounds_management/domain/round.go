package domain

type Round struct {
	ID      string
	Number  int
	Matches []*Match
	State   string
}

func NewRound() *Round {
	return &Round{
		ID:      "",
		Number:  0,
		Matches: nil,
	}
}

func (r *Round) Seed() error {

}

func (r *Round) Start() error {

}
