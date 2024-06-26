package usecases

import (
	"tournaments_backend/internal/tournament_management/domain"
	"tournaments_backend/internal/tournament_management/usecases/commands"
	"tournaments_backend/internal/tournament_management/usecases/queries"
)

type UseCases struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	EnrollPlayerHandler      *commands.EnrollPlayerHandler
	EnrollGuestPlayerHandler *commands.EnrollGuestPlayerHandler
	HostTournamentHandler    *commands.HostTournamentHandler
}

type Queries struct {
	HostedTournamentsHandler *queries.HostedTournamentsHandler
}

func NewUseCases(
	hostRepo domain.HostRepository,
	playerRepo domain.PlayerRepository,
	tournamentRepo domain.TournamentRepository,
	tournamentQueryService queries.TournamentQueryService,
) *UseCases {
	return &UseCases{
		Commands: Commands{
			EnrollPlayerHandler:      commands.NewEnrollPlayerHandler(nil, hostRepo, playerRepo, tournamentRepo),
			EnrollGuestPlayerHandler: commands.NewEnrollGuestPlayerHandler(nil, hostRepo, playerRepo, tournamentRepo),
			HostTournamentHandler:    commands.NewHostTournamentHandler(nil, hostRepo, tournamentRepo),
		},
		Queries: Queries{
			HostedTournamentsHandler: queries.NewHostedTournamentsHandler(hostRepo, tournamentQueryService),
		},
	}
}
