package handlers

import (
	"github.com/gin-gonic/gin"

	"github.com/parfentjev/raweceek/internal/session"
)

func RegisterRoutes(service *session.Service, router *gin.Engine) {
	h := newHandler(service)

	router.LoadHTMLGlob("public/templates/*.html")
	router.Static("/static", "public/static/")
	router.GET("/", h.index)
	router.GET("/api/next-session", h.nextSession)
}
