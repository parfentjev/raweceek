package main

import (
	"embed"
	"html/template"
	"io/fs"
	"log/slog"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/parfentjev/raweceek/internal"
	"github.com/parfentjev/raweceek/internal/codegen/db"
	"github.com/parfentjev/raweceek/internal/handlers"
	"github.com/parfentjev/raweceek/internal/session"
)

var (
	//go:embed public/*
	publicFS embed.FS
)

const (
	templatesPath = "public/templates/*.html"
	staticPath    = "public/static"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	if err := run(logger); err != nil {
		logger.Error("initalization failed", slog.Any("err", err))
		os.Exit(1)
	}
}

func run(logger *slog.Logger) error {
	config, err := internal.LoadConfig()
	if err != nil {
		return err
	}

	pool, err := internal.NewPool(config.Database)
	if err != nil {
		return err
	}

	defer pool.Close()

	queries := db.New(pool)
	repository := session.NewRepository(queries)
	service := session.NewService(repository)

	engine := gin.New()
	engine.Use(gin.Recovery())

	templates, err := template.ParseFS(publicFS, templatesPath)
	if err != nil {
		return err
	}
	engine.SetHTMLTemplate(templates)

	sub, err := fs.Sub(publicFS, staticPath)
	if err != nil {
		return err
	}
	engine.StaticFS("/static", http.FS(sub))

	router := handlers.NewRouter(service, engine, logger)
	router.RegisterRoutes()

	return engine.Run()
}
