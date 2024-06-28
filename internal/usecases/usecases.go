package usecases

import (
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"tournaments_backend/internal/domain"
	"tournaments_backend/internal/usecases/commands"
	"tournaments_backend/internal/usecases/queries"
)

type UseCases struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateHostHandler        *commands.CreateHostHandler
	CreatePlayerHandler      *commands.CreatePlayerHandler
	CreateUserHandler        *commands.CreateUserHandler
	DropPlayerHandler        *commands.DropPlayerHandler
	EnrollPlayerHandler      *commands.EnrollPlayerHandler
	EnrollGuestPlayerHandler *commands.EnrollGuestPlayerHandler
	HostTournamentHandler    *commands.HostTournamentHandler
	RecoverPlayerHandler     *commands.RecoverPlayerHandler
	SelfEnrollHandler        *commands.SelfEnrollHandler
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
			CreateHostHandler:        commands.NewCreateHostHandler(eventBus, hostRepo),
			CreatePlayerHandler:      commands.NewCreatePlayerHandler(eventBus, playerRepo),
			CreateUserHandler:        commands.NewCreateUserHandler(eventBus),
			DropPlayerHandler:        commands.NewDropPlayerHandler(eventBus, hostRepo, playerRepo, tournamentRepo),
			EnrollPlayerHandler:      commands.NewEnrollPlayerHandler(eventBus, hostRepo, playerRepo, tournamentRepo),
			EnrollGuestPlayerHandler: commands.NewEnrollGuestPlayerHandler(eventBus, hostRepo, playerRepo, tournamentRepo),
			HostTournamentHandler:    commands.NewHostTournamentHandler(eventBus, hostRepo, tournamentRepo),
			RecoverPlayerHandler:     commands.NewRecoverPlayerHandler(eventBus, hostRepo, playerRepo, tournamentRepo),
			SelfEnrollHandler:        commands.NewSelfEnrollHandler(eventBus, playerRepo, tournamentRepo),
		},
		Queries: Queries{
			EnrolledPlayersHandler:   queries.NewEnrolledPlayersHandler(hostRepo, tournamentRepo, playerQueryService),
			HostedTournamentsHandler: queries.NewHostedTournamentsHandler(hostRepo, tournamentQueryService),
			TournamentByIDHandler:    queries.NewTournamentByIDHandler(hostRepo, tournamentRepo, tournamentQueryService),
		},
	}
}
