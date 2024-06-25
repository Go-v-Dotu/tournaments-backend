package domain

type Player struct {
	ID            string
	Dropped       bool
	HasCommanders bool
}

func (p *Player) Drop() {
	p.Dropped = true
}

func (p *Player) Recover() {
	p.Dropped = false
}
