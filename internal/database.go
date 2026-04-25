package internal

import (
	"context"
	"fmt"
	"net"
	"net/url"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPool(config DatabaseConfig) (*pgxpool.Pool, error) {
	url := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s",
		config.User,
		url.QueryEscape(config.Password),
		net.JoinHostPort(config.Host, config.Port),
		config.Name,
		config.SSL,
	)

	return pgxpool.New(context.Background(), url)
}
