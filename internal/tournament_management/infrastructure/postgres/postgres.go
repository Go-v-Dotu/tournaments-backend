package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Client struct {
	*pgxpool.Pool
}

type Config struct {
	IP       string `yaml:"ip"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

func NewClient(ctx context.Context, config Config) (*Client, error) {
	pool, err := pgxpool.New(ctx, fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", config.User, config.Password, config.IP, config.Port, config.Database))
	if err != nil {
		return nil, fmt.Errorf("failed to create pool: %w", err)
	}

	if err = pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	return &Client{
		client: pool,
	}, nil
}
