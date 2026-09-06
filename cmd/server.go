package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/parfentjev/raweceek/internal/config"
	"github.com/parfentjev/raweceek/internal/generated/api"
	"github.com/parfentjev/raweceek/internal/generated/db"
	"github.com/parfentjev/raweceek/internal/handler"
	"github.com/parfentjev/raweceek/internal/schedule"
)

const ReadHeaderTimeout = 5 * time.Second

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	if err := run(logger); err != nil {
		logger.Error("program exited with an error", slog.Any("error", err))
	}
}

func run(logger *slog.Logger) error {
	cfg, err := config.New()
	if err != nil {
		return err
	}

	pool, err := dbPool(cfg)
	if err != nil {
		return err
	}

	defer pool.Close()

	service := schedule.New(logger, db.New(pool))
	server, err := httpServer(logger, cfg, service)
	if err != nil {
		return fmt.Errorf("failed to create http server: %w", err)
	}

	return server.ListenAndServe()
}

func dbPool(cfg config.Config) (*pgxpool.Pool, error) {
	pgxpoolArgs := fmt.Sprintf("host=%v dbname=%v user=%v password=%v",
		cfg.DatabaseHost,
		cfg.DatabaseName,
		cfg.DatabaseUser,
		cfg.DatabasePassword)

	return pgxpool.New(context.Background(), pgxpoolArgs)
}

func httpServer(logger *slog.Logger, cfg config.Config, service schedule.Service) (*http.Server, error) {
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
		ReadHeaderTimeout: ReadHeaderTimeout,
	}, nil
}
