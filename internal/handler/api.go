package handler

import (
	"encoding/json"
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

func (s *APIHandler) GetNextSession(_ http.ResponseWriter, _ *http.Request) {}

func (s *APIHandler) GetStatus(_ http.ResponseWriter, _ *http.Request) {}

func (s *APIHandler) GetStatusV2(w http.ResponseWriter, r *http.Request) {
	status, err := s.service.GetStatusV2(r.Context())
	if err != nil {
		s.logger.ErrorContext(r.Context(), "failed to get status", slog.Any("error", err))
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	body, err := json.Marshal(status)
	if err != nil {
		s.logger.ErrorContext(r.Context(), "failed to encode response body", slog.Any("error", err))
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if _, err = w.Write(body); err != nil {
		s.logger.ErrorContext(r.Context(), "failed to write response body", slog.Any("error", err))
	}
}
