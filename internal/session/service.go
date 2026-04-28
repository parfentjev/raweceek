package session

import (
	"context"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

type Service struct {
	repository *Repository
}

type Session struct {
	Summary    string      `json:"summary"`
	Location   string      `json:"location"`
	StartTime  time.Time   `json:"startTime"`
	ThisWeek   bool        `json:"thisWeek"`
	Countdowns []Countdown `json:"countdowns"`
}

type Countdown struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

const secondsInWeek = 7 * 24 * 60 * 60
const minCeeksValue = 0.001

func NewService(repository *Repository) *Service {
	return &Service{repository}
}

func (s *Service) GetNextSession(ctx context.Context) (*Session, error) {
	session, err := s.repository.GetNextSession(ctx)
	if err != nil {
		return nil, err
	}

	_, targetWeek := session.StartTime.Time.ISOWeek()
	_, currentWeek := time.Now().ISOWeek()

	thisWeek := targetWeek == currentWeek
	remainingTime := int64(time.Until(session.StartTime.Time).Seconds())

	return &Session{
		Summary:   session.Summary,
		Location:  session.Location,
		StartTime: session.StartTime.Time,
		ThisWeek:  thisWeek,
		Countdowns: []Countdown{
			calculateCeeks(remainingTime),
			calculateTimeUntil(remainingTime),
		},
	}, nil
}

func (s *Service) IsRaceWeek(ctx context.Context) (bool, error) {
	count, err := s.repository.CountSessionsThisWeek(ctx)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

//nolint:mnd //would be too verbose; context is clear
func calculateTimeUntil(remainingTime int64) Countdown {
	seconds := remainingTime % 60
	minutes := (remainingTime / 60) % 60
	hours := (remainingTime / (60 * 60)) % 24
	days := (remainingTime / (60 * 60 * 24)) % 7
	totalDays := remainingTime / (60 * 60 * 24)
	months := totalDays / 30
	weeks := (totalDays % 30) / 7

	type unit struct {
		value int64
		name  string
	}

	var sb strings.Builder
	for _, u := range []unit{{months, "month"}, {weeks, "week"}, {days, "day"}, {hours, "hour"}, {minutes, "minute"}} {
		if u.value > 0 {
			appendTimeUnit(&sb, u.value, u.name, " ")
		}
	}

	if sb.Len() > 0 {
		sb.WriteString("and ")
	}

	appendTimeUnit(&sb, seconds, "second", "")

	return Countdown{Type: "TIME_UNTIL", Value: sb.String()}
}

func appendTimeUnit(sb *strings.Builder, value int64, name, end string) {
	sb.WriteString(strconv.FormatInt(value, 10))
	sb.WriteString(" ")
	sb.WriteString(name)

	if value != 1 {
		sb.WriteString("s")
	}

	sb.WriteString(end)
}

func calculateCeeks(remainingTime int64) Countdown {
	ceeks := math.Max(float64(remainingTime)/secondsInWeek, minCeeksValue)
	value := fmt.Sprintf("%.3f ceeks", ceeks)

	return Countdown{Type: "CEEKS", Value: value}
}
