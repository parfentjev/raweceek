package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/parfentjev/raweceek/internal/codegen/db"
	"github.com/parfentjev/raweceek/internal/session"
)

func main() {
	router := gin.Default()

	postgresUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DATABASE_USER"),
		url.QueryEscape(os.Getenv("DATABASE_PASSWORD")),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_NAME"),
		os.Getenv("DATABASE_SSL"),
	)

	dbPool, err := pgxpool.New(context.Background(), postgresUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer dbPool.Close()

	queries := db.New(dbPool)
	repository := session.NewRepository(queries)
	service := session.NewService(repository)

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
