package usecases

import (
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"tournaments_backend/internal/tournament_management/domain"
	"tournaments_backend/internal/tournament_management/usecases/commands"
	"tournaments_backend/internal/tournament_management/usecases/queries"
)

type UseCases struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateHost               *commands.CreateHostHandler
	CreatePlayer             *commands.CreatePlayerHandler
	CreateUser               *commands.CreateUserHandler
	EnrollPlayerHandler      *commands.EnrollPlayerHandler
	EnrollGuestPlayerHandler *commands.EnrollGuestPlayerHandler
	HostTournamentHandler    *commands.HostTournamentHandler
}

type Queries struct {
	EnrolledPlayersHandler   *queries.EnrolledPlayersHandler
	HostedTournamentsHandler *queries.HostedTournamentsHandler
	TournamentByIDHandler    *queries.TournamentByIDHandler
}

func NewUseCases(
	eventBus *cqrs.EventBus,
	hostRepo domain.HostRepository,
	playerRepo domain.PlayerRepository,
	tournamentRepo domain.TournamentRepository,
	tournamentQueryService queries.TournamentQueryService,
	playerQueryService queries.PlayerQueryService,
) *UseCases {
	return &UseCases{
		Commands: Commands{
			CreateHost:               commands.NewCreateHostHandler(hostRepo),
			CreatePlayer:             commands.NewCreatePlayerHandler(playerRepo),
			CreateUser:               commands.NewCreateUserHandler(eventBus),
			EnrollPlayerHandler:      commands.NewEnrollPlayerHandler(eventBus, hostRepo, playerRepo, tournamentRepo),
			EnrollGuestPlayerHandler: commands.NewEnrollGuestPlayerHandler(eventBus, hostRepo, playerRepo, tournamentRepo),
			HostTournamentHandler:    commands.NewHostTournamentHandler(eventBus, hostRepo, tournamentRepo),
		},
		Queries: Queries{
			EnrolledPlayersHandler:   queries.NewEnrolledPlayersHandler(hostRepo, tournamentRepo, playerQueryService),
			HostedTournamentsHandler: queries.NewHostedTournamentsHandler(hostRepo, tournamentQueryService),
			TournamentByIDHandler:    queries.NewTournamentByIDHandler(hostRepo, tournamentRepo, tournamentQueryService),
		},
	}
}
