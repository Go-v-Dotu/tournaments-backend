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

type tournamentQueryService struct {
	client Client
	coll   *mongo.Collection
}

var (
	_ queries.TournamentQueryService = (*tournamentQueryService)(nil)
)

func NewTournamentQueryService(client Client) queries.TournamentQueryService {
	coll := client.Database("tournament_management").Collection("tournaments")
	return &tournamentQueryService{client: client, coll: coll}
}

func (r *tournamentQueryService) GetByID(ctx context.Context, id string) (*queries.Tournament, error) {
	oID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("")
	}

	f := bson.D{{"_id", oID}}
	res := r.coll.FindOne(ctx, f)

	var tournamentModel models.Tournament
	if err := res.Decode(&tournamentModel); err != nil {
		return nil, errors.New("")
	}

	respik := &queries.Tournament{
		ID:           tournamentModel.ID.Hex(),
		Title:        tournamentModel.Title,
		Date:         tournamentModel.Date.Time(),
		TotalPlayers: len(tournamentModel.Players),
	}

	return respik, nil
}

func (r *tournamentQueryService) GetByHostID(ctx context.Context, hostID string) ([]*queries.Tournament, error) {
	oID, err := primitive.ObjectIDFromHex(hostID)
	if err != nil {
		return nil, errors.New("")
	}

	f := bson.D{{"host_id", oID}}
	cur, err := r.coll.Find(ctx, f)
	if err != nil {
		return nil, err
	}

	tournaments := make(models.Tournaments, 0)
	if err := cur.All(ctx, &tournaments); err != nil {
		return nil, errors.New("")
	}

	respik := make([]*queries.Tournament, 0, len(tournaments))
	for _, tour := range tournaments {
		respik = append(respik, &queries.Tournament{
			ID:           tour.ID.Hex(),
			Title:        tour.Title,
			Date:         tour.Date.Time(),
			TotalPlayers: len(tour.Players),
		})
	}

	return respik, nil
}
