package handler

import (
	"encoding/json"
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
		h.writeError(w, r, fmt.Errorf("failed to get next session: %w", err), http.StatusInternalServerError)
		return
	}

	h.writeJSON(w, r, session)
}

func (h *APIHandler) GetStatus(w http.ResponseWriter, r *http.Request) {
	status, err := h.service.GetStatus(r.Context())
	if err != nil {
		h.writeError(w, r, fmt.Errorf("failed to get status: %w", err), http.StatusInternalServerError)
		return
	}

	h.writeJSON(w, r, status)
}

func (h *APIHandler) GetStatusV2(w http.ResponseWriter, r *http.Request) {
	status, err := h.service.GetStatusV2(r.Context())
	if err != nil {
		h.writeError(w, r, fmt.Errorf("failed to get status v2: %w", err), http.StatusInternalServerError)
		return
	}

	h.writeJSON(w, r, status)
}

//nolint:unparam // Always receives 500, but this is remporary—I'll add 404 later
func (h *APIHandler) writeError(w http.ResponseWriter, r *http.Request, err error, statusCode int) {
	h.logger.ErrorContext(r.Context(), "failed to process API request", slog.Any("error", err))
	http.Error(w, "internal server error", statusCode)
}

func (h *APIHandler) writeJSON(w http.ResponseWriter, r *http.Request, response any) {
	body, err := json.Marshal(response)
	if err != nil {
		h.writeError(w, r, fmt.Errorf("failed to encode response body: %w", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if _, err = w.Write(body); err != nil {
		h.writeError(w, r, fmt.Errorf("failed to write response body: %w", err), http.StatusInternalServerError)
	}
}
