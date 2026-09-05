package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/parfentjev/raweceek/internal/generated/api"
	"github.com/parfentjev/raweceek/internal/handler"
)

const ReadHeaderTimeout = 5 * time.Second

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	if err := run(logger); err != nil {
		logger.Error("program exited with an error", slog.Any("error", err))
	}
}

func run(logger *slog.Logger) error {
	apiHandler := handler.NewAPIHandler(logger)
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
