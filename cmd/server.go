package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/parfentjev/raweceek/internal/config"
	"github.com/parfentjev/raweceek/internal/generated/api"
	"github.com/parfentjev/raweceek/internal/generated/db"
	"github.com/parfentjev/raweceek/internal/handler"
	"github.com/parfentjev/raweceek/internal/schedule"
)

const (
	ShutdownTimeout         = 5 * time.Second
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

	pool, err := newDbPool(cfg.Database)
	if err != nil {
		return err
	}

	defer pool.Close()

	service := schedule.New(db.New(pool))
	server, err := newHttpServer(logger, cfg.Server, service)
	if err != nil {
		return fmt.Errorf("failed to create http server: %w", err)
	}

	return startHttpServer(logger, server)
}

func newDbPool(cfg config.Database) (*pgxpool.Pool, error) {
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

func newHttpServer(logger *slog.Logger, cfg config.Server, service schedule.Service) (*http.Server, error) {
	apiHandler := handler.NewAPIHandler(logger, service)
	staticHandler, err := handler.NewStaticHandler()
	if err != nil {
		return nil, err
	}

	mux := http.NewServeMux()
	mux.Handle("GET /", staticHandler)
	api.HandlerFromMux(&apiHandler, mux)

	return &http.Server{
		Addr:              cfg.BindAddressPublic,
		Handler:           mux,
		ReadTimeout:       ServerReadTimeout,
		ReadHeaderTimeout: ServerReadHeaderTimeout,
		WriteTimeout:      ServerWriteTimeout,
		IdleTimeout:       ServerIdleTimeout,
	}, nil
}

func startHttpServer(logger *slog.Logger, server *http.Server) error {
	signalContext, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	serverError := make(chan error, 1)
	go func() {
		serverError <- server.ListenAndServe()
	}()

	logger.Info("ready to handle requests", slog.String("address", server.Addr))

	var shutdownError error
	select {
	case shutdownError = <-serverError:
		// Server exited voluntarily, nothing to do here.
	case <-signalContext.Done():
		logger.Info("shutting down")
		stop()

		shutdownContext, cancelShutdown := context.WithTimeout(context.Background(), ShutdownTimeout)
		defer cancelShutdown()

		if err := server.Shutdown(shutdownContext); err != nil {
			return fmt.Errorf("failed to shutdown server: %w", err)
		}

		shutdownError = <-serverError
	}

	if shutdownError != nil {
		if errors.Is(shutdownError, http.ErrServerClosed) {
			return nil
		}

		return fmt.Errorf("fatal server error: %w", shutdownError)
	}

	return nil
}
