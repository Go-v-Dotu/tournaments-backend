package domain

type Settings struct {
	MaxPlayers int
}

func DefaultSettings() *Settings {
	return NewSettings(24)
}

func NewSettings(maxPlayers int) *Settings {
	return &Settings{MaxPlayers: maxPlayers}
}
