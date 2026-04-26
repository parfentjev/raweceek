package handlers

import (
	"encoding/json"
	"fmt"
	"log/slog"
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

func (handler *Handler) index(ctx *gin.Context) {
	nextSession, err := handler.service.GetNextSession(ctx.Request.Context())
	if err != nil {
		_ = ctx.Error(fmt.Errorf("failed to query next session: %w", err))
		return
	}

	sessionJSON, err := json.Marshal(nextSession)
	if err != nil {
		_ = ctx.Error(fmt.Errorf("failed to convert next session to JSON: %w", err))
		return
	}

	isRaceWeek, err := handler.service.IsRaceWeek(ctx.Request.Context())
	if err != nil {
		_ = ctx.Error(fmt.Errorf("failed to determine race week status: %w", err))
		return
	}

	data := gin.H{"session": nextSession, "sessionJSON": string(sessionJSON), "isRaceWeek": isRaceWeek}

	ctx.HTML(http.StatusOK, "index.html", data)
}

func (handler *Handler) nextSession(ctx *gin.Context) {
	session, err := handler.service.GetNextSession(ctx.Request.Context())
	if err != nil {
		_ = ctx.Error(fmt.Errorf("failed to query next session: %w", err))
		return
	}

	ctx.JSON(http.StatusOK, session)
}

// https://gin-gonic.com/en/docs/middleware/error-handling-middleware/
func errorHandler(logger *slog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) > 0 {
			err := ctx.Errors.Last().Err
			logger.Error("failed to process request", slog.Any("err", err))

			ctx.Status(http.StatusInternalServerError)
		}
	}
}
