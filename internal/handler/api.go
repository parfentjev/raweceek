package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/parfentjev/raweceek/internal/generated/api"
	"github.com/parfentjev/raweceek/internal/generated/db"
)

type APIHandler struct {
	logger  *slog.Logger
	queries *db.Queries
}

func NewAPIHandler(logger *slog.Logger, queries *db.Queries) APIHandler {
	return APIHandler{logger, queries}
}

func (s *APIHandler) GetNextSession(_ http.ResponseWriter, _ *http.Request) {}

func (s *APIHandler) GetStatus(_ http.ResponseWriter, _ *http.Request) {}

func (s *APIHandler) GetStatusV2(w http.ResponseWriter, r *http.Request) {
	upcomingSessions, err := s.queries.FindUpcoming(r.Context())
	if err != nil {
		s.logger.ErrorContext(r.Context(), "failed query upcoming sessions", slog.Any("error", err))
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	remainingTime := upcomingSessions[0].StartTime.Time.UTC().Sub(time.Now().UTC())

	response := api.StatusDtoV2{
		RaceWeek: upcomingSessions[0].ThisWeek,
		UpcomingSessions: []api.SessionDtoV2{
			{
				Countdowns: []api.CountdownDto{
					NewCeeksCountdown(remainingTime),
					NewTimeUntilCountdown(remainingTime),
				},
				Location:  upcomingSessions[0].Location,
				StartTime: upcomingSessions[0].StartTime.Time.UTC(),
				Summary:   upcomingSessions[0].Summary,
				ThisWeek:  upcomingSessions[0].ThisWeek,
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
