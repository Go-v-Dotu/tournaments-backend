package models

import (
	"tournaments_backend/internal/tournament_management/domain"
)

type Host struct {
	ID       string
	UserID   string
	Username string
}

func NewHost(host *domain.Host) *Host {
	return &Host{
		ID:       host.ID,
		UserID:   host.UserID,
		Username: host.Username,
	}
}

func (m *Host) ToEntity() *domain.Host {
	return domain.NewHost(m.ID, m.UserID, m.Username)
}
