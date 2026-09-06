package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/parfentjev/raweceek/internal/config"
	"github.com/parfentjev/raweceek/internal/generated/api"
	"github.com/parfentjev/raweceek/internal/generated/db"
	"github.com/parfentjev/raweceek/internal/handler"
	"github.com/parfentjev/raweceek/internal/schedule"
)

const (
	DatabasePingTimeout     = 2 * time.Second
	ServerReadHeaderTimeout = 2 * time.Second
	ServerReadTimeout       = 5 * time.Second
	ServerWriteTimeout      = 5 * time.Second
	ServerIdleTimeout       = 60 * time.Second
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	if err := run(logger); err != nil {
		logger.Error("program exited with an error", slog.Any("error", err))
		os.Exit(1)
	}
}

func run(logger *slog.Logger) error {
	cfg, err := config.New()
	if err != nil {
		return err
	}

	pool, err := dbPool(cfg.Database)
	if err != nil {
		return err
	}

	defer pool.Close()

	service := schedule.New(db.New(pool))
	server, err := httpServer(logger, cfg.Server, service)
	if err != nil {
		return fmt.Errorf("failed to create http server: %w", err)
	}

	return server.ListenAndServe()
}

func dbPool(cfg config.Database) (*pgxpool.Pool, error) {
	pgxpoolArgs := (&url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(cfg.User, cfg.Password),
		Host:   cfg.Host,
		Path:   cfg.Name,
	}).String()

	pool, err := pgxpool.New(context.Background(), pgxpoolArgs)
	if err != nil {
		return nil, fmt.Errorf("failed to create database pool: %w", err)
	}

	pingContext, cancelPingContext := context.WithTimeout(context.Background(), DatabasePingTimeout)
	defer cancelPingContext()

	if err = pool.Ping(pingContext); err != nil {
		pool.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return pool, nil
}

func httpServer(logger *slog.Logger, cfg config.Server, service schedule.Service) (*http.Server, error) {
	apiHandler := handler.NewAPIHandler(logger, service)
	staticHandler, err := handler.NewStaticHandler()
	if err != nil {
		return nil, err
	}

	mux := http.NewServeMux()
	mux.Handle("GET /", staticHandler)
	api.HandlerFromMux(&apiHandler, mux)

	return &http.Server{
		Addr:              cfg.BindAddress,
		Handler:           mux,
		ReadTimeout:       ServerReadTimeout,
		ReadHeaderTimeout: ServerReadHeaderTimeout,
		WriteTimeout:      ServerWriteTimeout,
		IdleTimeout:       ServerIdleTimeout,
	}, nil
}
