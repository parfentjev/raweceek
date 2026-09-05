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

	pgxpoolArgs := fmt.Sprintf("host=%v dbname=%v user=%v password=%v",
		cfg.DatabaseHost,
		cfg.DatabaseName,
		cfg.DatabaseUser,
		cfg.DatabasePassword)
	pool, err := pgxpool.New(context.Background(), pgxpoolArgs)
	if err != nil {
		return err
	}

	queries := db.New(pool)
	service := schedule.New(logger, queries)

	apiHandler := handler.NewAPIHandler(logger, service)
	staticHandler, err := handler.NewStaticHandler()
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.Handle("GET /", staticHandler)
	api.HandlerFromMux(&apiHandler, mux)

	server := &http.Server{
		Addr:              "127.0.0.1:8080",
		Handler:           mux,
		ReadHeaderTimeout: ReadHeaderTimeout,
	}

	return server.ListenAndServe()
}
