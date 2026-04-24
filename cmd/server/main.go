package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/parfentjev/raweceek/internal/codegen/db"
	"github.com/parfentjev/raweceek/internal/handlers"
	"github.com/parfentjev/raweceek/internal/session"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	config, err := loadConfig()
	if err != nil {
		return err
	}

	pool, err := newPool(config.Database)
	if err != nil {
		return err
	}

	defer pool.Close()

	queries := db.New(pool)
	repository := session.NewRepository(queries)
	service := session.NewService(repository)

	router := gin.Default()
	handlers.RegisterRoutes(service, router)

	return router.Run()
}
