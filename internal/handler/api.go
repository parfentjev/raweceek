package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/parfentjev/raweceek/internal/generated/api"
)

type APIHandler struct {
	logger *slog.Logger
}

func NewAPIHandler(logger *slog.Logger) APIHandler {
	return APIHandler{logger}
}

func (s *APIHandler) GetNextSession(_ http.ResponseWriter, _ *http.Request) {}

func (s *APIHandler) GetStatus(_ http.ResponseWriter, _ *http.Request) {}

func (s *APIHandler) GetStatusV2(w http.ResponseWriter, r *http.Request) {
	response := api.StatusDtoV2{
		RaceWeek: true,
		UpcomingSessions: []api.SessionDtoV2{
			{
				Countdowns: []api.CountdownDto{
					{Value: "0.123", Type: api.CEEKS},
					{Value: "1 day 23 hours and 69 minutes", Type: api.TIMEUNTIL},
				},
				Location:  "Estonia",
				StartTime: time.Now(),
				Summary:   "GRAND PRIX OF LASNAMÄE",
				ThisWeek:  true,
			},
		},
	}

	body, err := json.Marshal(response)
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
