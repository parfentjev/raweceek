package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/parfentjev/raweceek/internal/schedule"
)

type APIHandler struct {
	logger  *slog.Logger
	service schedule.Service
}

func NewAPIHandler(logger *slog.Logger, service schedule.Service) APIHandler {
	return APIHandler{logger, service}
}

func (h *APIHandler) GetNextSession(w http.ResponseWriter, r *http.Request) {
	session, err := h.service.GetNextSession(r.Context())
	if err != nil {
		if errors.Is(err, schedule.ErrSessionNotFound) {
			h.writeNotFound(w, r, fmt.Errorf("there are no upcoming sessions: %w", err))
			return
		}

		h.writeInternalServerError(w, r, fmt.Errorf("failed to get next session: %w", err))
		return
	}

	h.writeJSON(w, r, session)
}

func (h *APIHandler) GetStatus(w http.ResponseWriter, r *http.Request) {
	status, err := h.service.GetStatus(r.Context())
	if err != nil {
		if errors.Is(err, schedule.ErrSessionNotFound) {
			h.writeNotFound(w, r, fmt.Errorf("there are no upcoming sessions: %w", err))
			return
		}

		h.writeInternalServerError(w, r, fmt.Errorf("failed to get status: %w", err))
		return
	}

	h.writeJSON(w, r, status)
}

func (h *APIHandler) GetStatusV2(w http.ResponseWriter, r *http.Request) {
	status, err := h.service.GetStatusV2(r.Context())
	if err != nil {
		if errors.Is(err, schedule.ErrSessionNotFound) {
			h.writeNotFound(w, r, fmt.Errorf("there are no upcoming sessions: %w", err))
			return
		}

		h.writeInternalServerError(w, r, fmt.Errorf("failed to get status v2: %w", err))
		return
	}

	h.writeJSON(w, r, status)
}

func (h *APIHandler) writeInternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	h.logger.ErrorContext(r.Context(), "failed to process API request", slog.Any("error", err))
	http.Error(w, "internal server error", http.StatusInternalServerError)
}

func (h *APIHandler) writeNotFound(w http.ResponseWriter, r *http.Request, err error) {
	h.logger.InfoContext(r.Context(), "failed to process API request", slog.Any("error", err))
	w.WriteHeader(http.StatusNotFound)
}

func (h *APIHandler) writeJSON(w http.ResponseWriter, r *http.Request, response any) {
	body, err := json.Marshal(response)
	if err != nil {
		h.writeInternalServerError(w, r, fmt.Errorf("failed to encode response body: %w", err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if _, err = w.Write(body); err != nil {
		h.logger.ErrorContext(r.Context(), "failed to write response body", slog.Any("error", err))
	}
}
