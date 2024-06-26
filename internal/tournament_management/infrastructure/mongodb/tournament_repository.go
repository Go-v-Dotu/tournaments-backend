package mongodb

import (
	"context"
	"errors"

	"tournaments_backend/internal/tournament_management/domain"
	"tournaments_backend/internal/tournament_management/infrastructure/mongodb/models"
	"tournaments_backend/internal/tournament_management/usecases/queries"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type tournamentRepository struct {
	client Client
	coll   *mongo.Collection
}

var (
	_ domain.TournamentRepository = (*tournamentRepository)(nil)
)

func NewTournamentRepository(client Client) domain.TournamentRepository {
	coll := client.Database("tournament_management").Collection("tournaments")
	return &tournamentRepository{client: client, coll: coll}
}

func NewTournamentQueryService(client Client) queries.TournamentQueryService {
	coll := client.Database("tournament_management").Collection("tournaments")
	return &tournamentRepository{client: client, coll: coll}
}

func (r *tournamentRepository) Get(ctx context.Context, id string) (*domain.Tournament, error) {
	oID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("")
	}

	f := bson.D{{"_id", oID}}
	res := r.coll.FindOne(ctx, f)

	var tournamentModel models.Tournament
	if err := res.Decode(&tournamentModel); err != nil {
		//if errors.Is(err, mongo.ErrNoDocuments) {
		//	return nil, nil
		//}
		return nil, errors.New("")
	}

	return tournamentModel.ToEntity(), nil
}

func (r *tournamentRepository) GetByHostID(ctx context.Context, hostID string) (queries.Tournaments, error) {
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

	return tournaments.ToResponse(), nil
}

func (r *tournamentRepository) Save(ctx context.Context, tournament *domain.Tournament) error {
	tournamentModel := models.NewTournament(tournament)

	if _, err := r.coll.InsertOne(ctx, tournamentModel); err != nil {
		return errors.New("")
	}

	return nil
}

func (r *tournamentRepository) Delete(ctx context.Context, tournament *domain.Tournament) error {
	return nil
}

func (r *tournamentRepository) NextID(_ context.Context) string {
	return primitive.NewObjectID().Hex()
}
