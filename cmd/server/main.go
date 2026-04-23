package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/parfentjev/raweceek/internal/codegen/db"
	"github.com/parfentjev/raweceek/internal/session"
)

func main() {
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
			ctx.Status(500)
			return
		}

		ctx.String(200, session.Summary)
	})

	if err := router.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
