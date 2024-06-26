package models

import (
	"tournaments_backend/internal/tournament_management/domain"
	"tournaments_backend/internal/tournament_management/usecases/queries"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tournament struct {
	ID       primitive.ObjectID `bson:"_id"`
	HostID   primitive.ObjectID `bson:"host_id"`
	Title    string             `bson:"title"`
	Settings *Settings          `bson:"settings"`
	Players  EnrolledPlayers    `bson:"players"`
	Date     primitive.DateTime `bson:"date"`
	State    TournamentState    `bson:"state"`
}

type Settings struct {
	MaxPlayers int `bson:"max_players"`
}

type EnrolledPlayer struct {
	PlayerID      primitive.ObjectID `bson:"player_id"`
	Dropped       bool               `bson:"dropped"`
	HasCommanders bool               `bson:"has_commanders"`
}

type EnrolledPlayers []*EnrolledPlayer

type TournamentState uint8

func (m *Tournament) ToEntity() *domain.Tournament {
	return domain.NewTournament(
		m.ID.Hex(),
		m.HostID.Hex(),
		m.Title,
		m.Settings.ToEntity(),
		m.Players.ToEntity(),
		nil,
		m.Date.Time(),
		domain.TournamentState(m.State),
	)
}

func NewTournament(tournament *domain.Tournament) *Tournament {
	id, err := primitive.ObjectIDFromHex(tournament.ID)
	if err != nil {
		panic(err)
	}
	hostID, err := primitive.ObjectIDFromHex(tournament.HostID)
	if err != nil {
		panic(err)
	}
	return &Tournament{
		ID:       id,
		HostID:   hostID,
		Title:    tournament.Title,
		Settings: NewSettings(tournament.Settings),
		Players:  NewEnrolledPlayers(tournament.Players),
		Date:     primitive.NewDateTimeFromTime(tournament.Date),
		State:    TournamentState(tournament.State),
	}
}

func (m *Settings) ToEntity() *domain.Settings {
	return domain.NewSettings(m.MaxPlayers)
}

func NewSettings(settings *domain.Settings) *Settings {
	return &Settings{MaxPlayers: settings.MaxPlayers}
}

func (m *EnrolledPlayer) ToEntity() *domain.EnrolledPlayer {
	return domain.NewEnrolledPlayer(m.PlayerID.Hex(), m.Dropped, m.HasCommanders)
}

func NewEnrolledPlayer(player *domain.EnrolledPlayer) *EnrolledPlayer {
	playerID, err := primitive.ObjectIDFromHex(player.PlayerID)
	if err != nil {
		panic(err)
	}
	return &EnrolledPlayer{
		PlayerID:      playerID,
		Dropped:       player.Dropped,
		HasCommanders: player.HasCommanders,
	}
}

func (ms EnrolledPlayers) ToEntity() []*domain.EnrolledPlayer {
	players := make([]*domain.EnrolledPlayer, 0, len(ms))
	for _, m := range ms {
		players = append(players, m.ToEntity())
	}
	return players
}

func NewEnrolledPlayers(players []*domain.EnrolledPlayer) EnrolledPlayers {
	ms := make(EnrolledPlayers, 0, len(players))
	for _, player := range players {
		ms = append(ms, NewEnrolledPlayer(player))
	}
	return ms
}

type Tournaments []*Tournament

func (m *Tournament) ToResponse() *queries.Tournament {
	return &queries.Tournament{
		ID:           m.ID.Hex(),
		Title:        m.Title,
		Date:         m.Date.Time(),
		TotalPlayers: len(m.Players),
	}
}

func (ms Tournaments) ToResponse() queries.Tournaments {
	tournaments := make(queries.Tournaments, 0, len(ms))
	for _, m := range ms {
		tournaments = append(tournaments, m.ToResponse())
	}
	return tournaments
}
