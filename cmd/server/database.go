package main

import (
	"context"
	"fmt"
	"net/url"

	"github.com/jackc/pgx/v5/pgxpool"
)

func newPool(config DatabaseConfig) (*pgxpool.Pool, error) {
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		config.User,
		url.QueryEscape(config.Password),
		config.Host,
		config.Port,
		config.Name,
		config.SSL,
	)

	return pgxpool.New(context.Background(), url)
}
