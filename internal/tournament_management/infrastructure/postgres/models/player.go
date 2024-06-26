package models

import "tournaments_backend/internal/tournament_management/domain"

type Player struct {
	ID       string
	UserID   string
	Username string
}

func NewPlayer(player *domain.Player) *Player {
	return &Player{
		ID:       player.ID,
		UserID:   player.UserID,
		Username: player.Username,
	}
}

func (m *Player) ToEntity() *domain.Player {
	return domain.NewPlayer(m.ID, m.UserID, m.Username)
}
