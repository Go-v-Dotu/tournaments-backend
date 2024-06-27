package mongodb

import (
	"context"
	"errors"

	"tournaments_backend/internal/tournament_management/infrastructure/mongodb/models"
	"tournaments_backend/internal/tournament_management/usecases/queries"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type playerQueryService struct {
	client         Client
	tournamentColl *mongo.Collection
	playerColl     *mongo.Collection
}

var (
	_ queries.PlayerQueryService = (*playerQueryService)(nil)
)

func NewPlayerQueryService(client Client) queries.PlayerQueryService {
	tournamentColl := client.Database("tournament_management").Collection("tournaments")
	playerColl := client.Database("tournament_management").Collection("players")
	return &playerQueryService{
		client:         client,
		tournamentColl: tournamentColl,
		playerColl:     playerColl,
	}
}

func (r *playerQueryService) GetByTournamentID(ctx context.Context, tournamentID string) ([]*queries.Player, error) {
	oID, err := primitive.ObjectIDFromHex(tournamentID)
	if err != nil {
		return nil, errors.New("")
	}

	f := bson.D{{"_id", oID}}
	res := r.tournamentColl.FindOne(ctx, f)

	var tournamentModel models.Tournament
	if err := res.Decode(&tournamentModel); err != nil {
		return nil, errors.New("")
	}

	pleps := make([]*queries.Player, 0, len(tournamentModel.Players))
	for _, ppp := range tournamentModel.Players {
		f := bson.D{{"_id", ppp.PlayerID}}
		res := r.playerColl.FindOne(ctx, f)

		var playerModel models.Player
		if err := res.Decode(&playerModel); err != nil {
			return nil, errors.New("")
		}

		pleps = append(pleps, &queries.Player{
			ID:       ppp.PlayerID.Hex(),
			UserID:   playerModel.UserID,
			Username: playerModel.Username,
			Dropped:  ppp.Dropped,
		})
	}

	return pleps, nil
}
