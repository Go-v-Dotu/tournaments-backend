package mongodb

import (
	"context"
	"errors"

	"tournaments_backend/internal/tournament_management/domain"
	"tournaments_backend/internal/tournament_management/infrastructure/mongodb/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type playerRepository struct {
	client Client
	coll   *mongo.Collection
}

var (
	_ domain.PlayerRepository = (*playerRepository)(nil)
)

func NewPlayerRepository(client Client) domain.PlayerRepository {
	coll := client.Database("tournament_management").Collection("players")
	return &playerRepository{client: client, coll: coll}
}

func (r *playerRepository) GetByUserID(ctx context.Context, userID string) (*domain.Player, error) {
	f := bson.D{{Key: "user_id", Value: userID}}
	res := r.coll.FindOne(ctx, f)

	var playerModel models.Player
	if err := res.Decode(&playerModel); err != nil {
		//if errors.Is(err, mongo.ErrNoDocuments) {
		//	return nil, nil
		//}
		return nil, errors.New("")
	}

	return playerModel.ToEntity(), nil
}

func (r *playerRepository) Save(ctx context.Context, player *domain.Player) error {
	playerModel := models.NewPlayer(player)

	if _, err := r.coll.InsertOne(ctx, playerModel); err != nil {
		return errors.New("")
	}

	return nil
}

func (r *playerRepository) Delete(ctx context.Context, player *domain.Player) error {
	return nil
}

func (r *playerRepository) NextID(_ context.Context) string {
	return primitive.NewObjectID().Hex()
}
