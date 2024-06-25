package domain

import (
	"errors"
	"slices"
	"time"
)

type Tournament struct {
	ID       string
	HostID   string
	Title    string
	Settings *Settings
	Players  []*Player
	Rounds   []*Round
	Date     time.Time
	State    TournamentState
}

func NewTournament(id string, hostID string, title string, settings *Settings, players []*Player, rounds []*Round, date time.Time, state TournamentState) *Tournament {
	t := &Tournament{
		ID:       id,
		HostID:   hostID,
		Title:    title,
		Settings: settings,
		Players:  players,
		Rounds:   rounds,
		Date:     date,
		State:    state,
	}
	return t
}

func (t *Tournament) IsHostedBy(hostID string) bool {
	return t.HostID == hostID
}

func (t *Tournament) EnrollPlayer(p *Player) error {
	if t.State != TournamentStateCreated {
		return errors.New("tournament should be created to enroll the player")
	}
	if slices.ContainsFunc(t.Players, func(other *Player) bool {
		return other.ID == p.ID
	}) {
		return errors.New("already enrolled")
	}
	if len(t.Players) == t.Settings.MaxPlayers {
		return errors.New("too many players")
	}
	t.Players = append(t.Players, p)
	return nil
}

func (t *Tournament) RemovePlayer(p *Player) error {
	if t.State != TournamentStateCreated {
		return errors.New("tournament should be created to remove the player")
	}
	for i, enrolledPlayer := range t.Players {
		if enrolledPlayer.ID == p.ID {
			slices.Delete(t.Players, i, i+1)
			return nil
		}
	}
	return errors.New("not enrolled")
}

func (t *Tournament) DropPlayer(p *Player) error {
	if t.State != TournamentStateStarted {
		return errors.New("tournament should be started to drop the player")
	}
	for _, enrolledPlayer := range t.Players {
		if enrolledPlayer.ID == p.ID {
			enrolledPlayer.Drop()
			return nil
		}
	}
	return errors.New("not enrolled")
}

func (t *Tournament) RecoverPlayer(p *Player) error {
	if t.State != TournamentStateStarted {
		return errors.New("tournament should be started to recover the player")
	}
	for _, enrolledPlayer := range t.Players {
		if enrolledPlayer.ID == p.ID {
			enrolledPlayer.Recover()
			return nil
		}
	}
	return errors.New("not enrolled")
}

func (t *Tournament) Start() error {
	if t.State != TournamentStateCreated {
		return errors.New("")
	}
	for _, p := range t.Players {
		if !p.HasCommanders {
			return errors.New("")
		}
	}
	t.State = TournamentStateStarted
	return nil
}

func (t *Tournament) Finish() error {
	if t.State != TournamentStateStarted {
		return errors.New("")
	}
	t.State = TournamentStateFinished
	return nil
}
