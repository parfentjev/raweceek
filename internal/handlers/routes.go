package handlers

import (
	"log/slog"

	"github.com/gin-gonic/gin"

	"github.com/parfentjev/raweceek/internal/session"
)

type Router struct {
	service *session.Service
	gin     *gin.Engine
	logger  *slog.Logger
}

func NewRouter(service *session.Service, router *gin.Engine, logger *slog.Logger) *Router {
	return &Router{service, router, logger}
}

func (router *Router) RegisterRoutes() {
	router.gin.Use(errorHandler(router.logger))

	handler := newHandler(router.service)
	router.gin.GET("/", handler.index)
	router.gin.GET("/api/next-session", handler.nextSession)
}
