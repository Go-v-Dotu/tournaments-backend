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
	Players  []*EnrolledPlayer
	Rounds   []*Round
	Date     time.Time
	State    TournamentState
}

func CreateTournament(id string, host *Host, title string, date time.Time) *Tournament {
	return NewTournament(
		id,
		host.ID,
		title,
		DefaultSettings(),
		make([]*EnrolledPlayer, 0),
		make([]*Round, 0),
		date,
		TournamentStateCreated,
	)
}

func NewTournament(
	id string,
	hostID string,
	title string,
	settings *Settings,
	players []*EnrolledPlayer,
	rounds []*Round,
	date time.Time,
	state TournamentState,
) *Tournament {
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

func (t *Tournament) IsHostedBy(host *Host) bool {
	return t.HostID == host.ID
}

func (t *Tournament) EnrollPlayer(host *Host, player *Player) error {
	if !t.IsHostedBy(host) {
		return errors.New("you aren't a host of this tournament")
	}
	if t.State != TournamentStateCreated {
		return errors.New("tournament should be created to enroll the player")
	}
	if slices.ContainsFunc(t.Players, func(other *EnrolledPlayer) bool {
		return other.PlayerID == player.ID
	}) {
		return errors.New("already enrolled")
	}
	if len(t.Players) == t.Settings.MaxPlayers {
		return errors.New("too many players")
	}
	t.Players = append(t.Players, &EnrolledPlayer{
		PlayerID:      player.ID,
		Dropped:       false,
		HasCommanders: false,
	})
	return nil
}

func (t *Tournament) RemovePlayer(host *Host, player *Player) error {
	if !t.IsHostedBy(host) {
		return errors.New("you aren't a host of this tournament")
	}
	if t.State != TournamentStateCreated {
		return errors.New("tournament should be created to remove the player")
	}
	for i, enrolledPlayer := range t.Players {
		if enrolledPlayer.PlayerID == player.ID {
			slices.Delete(t.Players, i, i+1)
			return nil
		}
	}
	return errors.New("not enrolled")
}

func (t *Tournament) DropPlayer(host *Host, player *Player) error {
	if !t.IsHostedBy(host) {
		return errors.New("you aren't a host of this tournament")
	}
	if t.State != TournamentStateStarted {
		return errors.New("tournament should be started to drop the player")
	}
	for _, enrolledPlayer := range t.Players {
		if enrolledPlayer.PlayerID == player.ID {
			enrolledPlayer.Drop()
			return nil
		}
	}
	return errors.New("not enrolled")
}

func (t *Tournament) RecoverPlayer(host *Host, player *Player) error {
	if !t.IsHostedBy(host) {
		return errors.New("you aren't a host of this tournament")
	}
	if t.State != TournamentStateStarted {
		return errors.New("tournament should be started to recover the player")
	}
	for _, enrolledPlayer := range t.Players {
		if enrolledPlayer.PlayerID == player.ID {
			enrolledPlayer.Recover()
			return nil
		}
	}
	return errors.New("not enrolled")
}

func (t *Tournament) Start(host *Host) error {
	if !t.IsHostedBy(host) {
		return errors.New("you aren't a host of this tournament")
	}
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

func (t *Tournament) Finish(host *Host) error {
	if !t.IsHostedBy(host) {
		return errors.New("you aren't a host of this tournament")
	}
	if t.State != TournamentStateStarted {
		return errors.New("")
	}
	t.State = TournamentStateFinished
	return nil
}
