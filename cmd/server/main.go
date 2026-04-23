package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/parfentjev/raweceek/internal/codegen/db"
	"github.com/parfentjev/raweceek/internal/session"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	config, err := loadConfig()
	if err != nil {
		log.Fatal(err)
	}

	pool, err := newPool(config.Database)
	if err != nil {
		log.Fatal(err)
	}

	defer pool.Close()

	queries := db.New(pool)
	repository := session.NewRepository(queries)
	service := session.NewService(repository)

	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		session, err := service.GetNextSession(ctx.Request.Context())
		if err != nil {
			log.Println(err)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		ctx.String(http.StatusOK, session.Summary)
	})

	return router.Run()
}
