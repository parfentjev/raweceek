package schedule

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/parfentjev/raweceek/internal/generated/api"
	"github.com/parfentjev/raweceek/internal/generated/db"
)

var ErrSessionNotFound = errors.New("upcoming session(s) not found")

type Service struct {
	queries *db.Queries
}

func New(queries *db.Queries) Service {
	return Service{queries}
}

func (s *Service) GetNextSession(ctx context.Context) (api.SessionDto, error) {
	session, err := s.queries.FindNext(ctx)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return api.SessionDto{}, ErrSessionNotFound
		}

		return api.SessionDto{}, err
	}

	countdown := newCountdownCalc(time.Until(session.StartTime.Time))

	return api.SessionDto{
		Countdowns: []api.CountdownDto{
			countdown.ceeks(),
			countdown.timeUntil(),
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

	countdown := newCountdownCalc(time.Until(session.StartTime.Time))

	return api.StatusDto{
		RaceWeek: session.ThisWeek.Bool,
		NextSession: api.SessionDto{
			Countdowns: []api.CountdownDto{
				countdown.ceeks(),
				countdown.timeUntil(),
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
		countdown := newCountdownCalc(time.Until(row.StartTime.Time))

		sessions = append(sessions, api.SessionDtoV2{
			Countdowns: []api.CountdownDto{
				countdown.ceeks(),
				countdown.timeUntil(),
			},
			Location:  row.Location,
			StartTime: row.StartTime.Time.UTC(),
			Summary:   row.Summary,
			ThisWeek:  row.ThisWeek,
		})
	}

	return sessions
}
