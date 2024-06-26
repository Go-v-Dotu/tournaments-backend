package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	*mongo.Client
}

type Config struct {
	IP         string `yaml:"ip"`
	Port       int    `yaml:"port"`
	User       string `yaml:"user"`
	Password   string `yaml:"password"`
	AuthSource string `yaml:"auth_source"`
}

func NewClient(ctx context.Context, config Config) (*Client, error) {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d/?authSource=%s", config.User, config.Password, config.IP, config.Port, config.AuthSource)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return &Client{client}, nil
}
