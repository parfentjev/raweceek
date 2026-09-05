package schedule

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/parfentjev/raweceek/internal/generated/api"
	"github.com/parfentjev/raweceek/internal/generated/db"
)

type Service struct {
	logger  *slog.Logger
	queries *db.Queries
}

func New(logger *slog.Logger, queries *db.Queries) Service {
	return Service{logger, queries}
}

func (s *Service) GetStatusV2(ctx context.Context) (api.StatusDtoV2, error) {
	upcomingSessions, err := s.queries.FindUpcoming(ctx)
	if err != nil {
		return api.StatusDtoV2{}, fmt.Errorf("failed to query database: %w", err)
	}

	return api.StatusDtoV2{
		RaceWeek:         upcomingSessions[0].ThisWeek,
		UpcomingSessions: mapRowsToSessions(upcomingSessions),
	}, nil
}

func mapRowsToSessions(rows []db.FindUpcomingRow) []api.SessionDtoV2 {
	sessions := make([]api.SessionDtoV2, 0, len(rows))
	remainingTime := NewRemainingTime(time.Until(rows[0].StartTime.Time))

	for _, row := range rows {
		sessions = append(sessions, api.SessionDtoV2{
			Countdowns: []api.CountdownDto{
				remainingTime.Ceeks(),
				remainingTime.TimeUntil(),
			},
			Location:  row.Location,
			StartTime: row.StartTime.Time.UTC(),
			Summary:   row.Summary,
			ThisWeek:  row.ThisWeek,
		})
	}

	return sessions
}
