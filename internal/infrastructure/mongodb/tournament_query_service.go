package mongodb

import (
	"context"
	"fmt"

	"tournaments_backend/internal/infrastructure/mongodb/models"
	"tournaments_backend/internal/usecases/queries"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type tournamentQueryService struct {
	client         Client
	tournamentColl *mongo.Collection
	hostColl       *mongo.Collection
}

var (
	_ queries.TournamentQueryService = (*tournamentQueryService)(nil)
)

func NewTournamentQueryService(client Client) queries.TournamentQueryService {
	tournamentColl := client.Database("tournament_management").Collection("tournaments")
	hostColl := client.Database("tournament_management").Collection("hosts")
	return &tournamentQueryService{client: client, tournamentColl: tournamentColl, hostColl: hostColl}
}

func (r *tournamentQueryService) GetByID(ctx context.Context, id string) (*queries.Tournament, error) {
	oID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("error getting tournament by id: %w", err)
	}

	f := bson.D{{Key: "_id", Value: oID}}
	res := r.tournamentColl.FindOne(ctx, f)

	var tournamentModel models.Tournament
	if err := res.Decode(&tournamentModel); err != nil {
		return nil, fmt.Errorf("tournament not found: %w", err)
	}

	f = bson.D{{Key: "_id", Value: tournamentModel.HostID}}
	res = r.hostColl.FindOne(ctx, f)

	var hostModel models.Host
	if err := res.Decode(&hostModel); err != nil {
		return nil, fmt.Errorf("host not found: %w", err)
	}

	tournament := &queries.Tournament{
		ID:    tournamentModel.ID.Hex(),
		Title: tournamentModel.Title,
		Host: &queries.Host{
			ID:       hostModel.ID.Hex(),
			UserID:   hostModel.UserID,
			Username: hostModel.Username,
		},
		Date:         tournamentModel.Date.Time(),
		TotalPlayers: len(tournamentModel.Players),
	}

	return tournament, nil
}

func (r *tournamentQueryService) GetByHostID(ctx context.Context, hostID string) ([]*queries.Tournament, error) {
	oID, err := primitive.ObjectIDFromHex(hostID)
	if err != nil {
		return nil, fmt.Errorf("error getting tournament by host: %w", err)
	}

	f := bson.D{{Key: "host_id", Value: oID}}
	cur, err := r.tournamentColl.Find(ctx, f)
	if err != nil {
		return nil, fmt.Errorf("tournament not found: %w", err)
	}

	tournamentModels := make(models.Tournaments, 0)
	if err := cur.All(ctx, &tournamentModels); err != nil {
		return nil, fmt.Errorf("error getting tournament by host: %w", err)
	}

	tournaments := make([]*queries.Tournament, 0, len(tournamentModels))
	for _, tournamentModel := range tournamentModels {
		f := bson.D{{Key: "_id", Value: tournamentModel.HostID}}
		res := r.hostColl.FindOne(ctx, f)

		var hostModel models.Host
		if err := res.Decode(&hostModel); err != nil {
			return nil, fmt.Errorf("host not found: %w", err)
		}
		tournaments = append(tournaments, &queries.Tournament{
			ID:    tournamentModel.ID.Hex(),
			Title: tournamentModel.Title,
			Host: &queries.Host{
				ID:       hostModel.ID.Hex(),
				UserID:   hostModel.UserID,
				Username: hostModel.Username,
			},
			Date:         tournamentModel.Date.Time(),
			TotalPlayers: len(tournamentModel.Players),
		})
	}

	return tournaments, nil
}
