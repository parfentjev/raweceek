package handler

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/parfentjev/raweceek/internal/generated/api"
)

const secondsPerWeek = 7 * 24 * 60 * 60

type timeUntilUnit struct {
	name  string
	value int64
}

func NewCeeksCountdown(remainingTime time.Duration) api.CountdownDto {
	ceeks := remainingTime.Seconds() / secondsPerWeek
	if ceeks < 0.001 {
		ceeks = 0.001
	}

	return api.CountdownDto{
		Type:  api.CEEKS,
		Value: fmt.Sprintf("%.3f ceeks", ceeks),
	}
}

func NewTimeUntilCountdown(remainingTime time.Duration) api.CountdownDto {
	units := extractTimeUntilUnits(remainingTime)
	parts := make([]string, 0, len(units))

	for _, unit := range units {
		if unit.value <= 0 {
			continue
		}

		name := unit.name
		if unit.value > 1 {
			name += "s"
		}

		parts = append(parts, strconv.FormatInt(unit.value, 10)+" "+name)
	}

	value := "0s"
	switch len(parts) {
	case 1:
		value = parts[0]
	case 2:
		value = parts[0] + " and " + parts[1]
	default:
		if len(parts) > 2 {
			value = strings.Join(parts[:len(parts)-1], ", ") + " and " + parts[len(parts)-1]
		}
	}

	return api.CountdownDto{
		Type:  api.TIMEUNTIL,
		Value: value,
	}
}

func extractTimeUntilUnits(remainingTime time.Duration) [6]timeUntilUnit {
	totalSeconds := int64(remainingTime / time.Second)
	seconds := totalSeconds % 60
	minutes := totalSeconds / 60 % 60
	hours := totalSeconds / (60 * 60) % 24

	// This part isn't very precise because not all months are exactly 30 days long,
	// But this level of precision is acceptable.
	totalDays := totalSeconds / (60 * 60 * 24)
	months := totalDays / 30
	remainingDays := totalDays % 30
	weeks := remainingDays / 7
	days := remainingDays % 7

	return [6]timeUntilUnit{
		{name: "month", value: months},
		{name: "week", value: weeks},
		{name: "day", value: days},
		{name: "hour", value: hours},
		{name: "minute", value: minutes},
		{name: "second", value: seconds},
	}
}
