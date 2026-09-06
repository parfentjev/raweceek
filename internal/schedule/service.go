package schedule

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/parfentjev/raweceek/internal/generated/api"
	"github.com/parfentjev/raweceek/internal/generated/db"
)

var ErrSessionNotFound = errors.New("next session(s) not found")

type Service struct {
	logger  *slog.Logger
	queries *db.Queries
}

func New(logger *slog.Logger, queries *db.Queries) Service {
	return Service{logger, queries}
}

func (s *Service) GetNextSession(ctx context.Context) (api.SessionDto, error) {
	session, err := s.queries.FindNext(ctx)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return api.SessionDto{}, ErrSessionNotFound
		}

		return api.SessionDto{}, err
	}

	remainingTime := startToRemainingTime(session.StartTime.Time)

	return api.SessionDto{
		Countdowns: []api.CountdownDto{
			remainingTime.Ceeks(),
			remainingTime.TimeUntil(),
		},
		Location:  session.Location,
		StartTime: session.StartTime.Time.UTC(),
		Summary:   session.Summary,
		ThisWeek:  session.ThisWeek.Bool}, nil
}

func (s *Service) GetStatus(ctx context.Context) (api.StatusDto, error) {
	session, err := s.queries.FindNext(ctx)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return api.StatusDto{}, ErrSessionNotFound
		}

		return api.StatusDto{}, err
	}

	remainingTime := startToRemainingTime(session.StartTime.Time)

	return api.StatusDto{
		RaceWeek: session.ThisWeek.Bool,
		NextSession: api.SessionDto{
			Countdowns: []api.CountdownDto{
				remainingTime.Ceeks(),
				remainingTime.TimeUntil(),
			},
			Location:  session.Location,
			StartTime: session.StartTime.Time.UTC(),
			Summary:   session.Summary,
			ThisWeek:  session.ThisWeek.Bool},
	}, nil
}

func (s *Service) GetStatusV2(ctx context.Context) (api.StatusDtoV2, error) {
	upcomingSessions, err := s.queries.FindUpcoming(ctx)
	if err != nil {
		return api.StatusDtoV2{}, fmt.Errorf("failed to query database: %w", err)
	}

	if len(upcomingSessions) == 0 {
		return api.StatusDtoV2{}, ErrSessionNotFound
	}

	return api.StatusDtoV2{
		RaceWeek:         upcomingSessions[0].ThisWeek,
		UpcomingSessions: mapRowsToSessions(upcomingSessions),
	}, nil
}

func mapRowsToSessions(rows []db.FindUpcomingRow) []api.SessionDtoV2 {
	sessions := make([]api.SessionDtoV2, 0, len(rows))
	for _, row := range rows {
		remainingTime := startToRemainingTime(row.StartTime.Time)

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

func startToRemainingTime(startTime time.Time) RemainingTime {
	return NewRemainingTime(time.Until(startTime))
}
