package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/parfentjev/raweceek/internal/session"
)

type Handler struct {
	service *session.Service
}

func newHandler(service *session.Service) *Handler {
	return &Handler{service}
}

func (h *Handler) index(ctx *gin.Context) {
	nextSession, err := h.service.GetNextSession(ctx.Request.Context())
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	sessionJSON, err := json.Marshal(nextSession)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	isRaceWeek, err := h.service.IsRaceWeek(ctx.Request.Context())
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	data := gin.H{"session": nextSession, "sessionJSON": string(sessionJSON), "isRaceWeek": isRaceWeek}

	ctx.HTML(http.StatusOK, "index.html", data)
}

func (h *Handler) nextSession(ctx *gin.Context) {
	session, err := h.service.GetNextSession(ctx.Request.Context())
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, session)
}
