package schedule

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/parfentjev/raweceek/internal/generated/api"
)

const secondsPerWeek = 7 * 24 * 60 * 60
const minCeeksValue = 0.001

type timeUntilUnit struct {
	name  string
	value int64
}

type countdownCalc struct {
	remainingTime time.Duration
}

func newCountdownCalc(remainingTime time.Duration) countdownCalc {
	return countdownCalc{remainingTime}
}

func (rt *countdownCalc) ceeks() api.CountdownDto {
	ceeks := rt.remainingTime.Seconds() / secondsPerWeek
	if ceeks < minCeeksValue {
		ceeks = minCeeksValue
	}

	return api.CountdownDto{
		Type:  api.CEEKS,
		Value: fmt.Sprintf("%.3f ceeks", ceeks),
	}
}

func (rt *countdownCalc) timeUntil() api.CountdownDto {
	timeUnits := extractTimeUntilUnits(rt.remainingTime)

	var result strings.Builder
	for i, unit := range timeUnits {
		finalUnit := i == len(timeUnits)-1
		addSeparator(&result, finalUnit)
		appendUnit(&result, unit.value, unit.name)
	}

	if result.Len() == 0 {
		result.WriteString("0s")
	}

	return api.CountdownDto{
		Type:  api.TIMEUNTIL,
		Value: result.String(),
	}
}

//nolint:mnd // Numbers are clear from the context, defining consts would be an overkill
func extractTimeUntilUnits(remainingTime time.Duration) []timeUntilUnit {
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

	return slices.DeleteFunc([]timeUntilUnit{
		{name: "month", value: months},
		{name: "week", value: weeks},
		{name: "day", value: days},
		{name: "hour", value: hours},
		{name: "minute", value: minutes},
		{name: "second", value: seconds},
	}, func(unit timeUntilUnit) bool {
		return unit.value <= 0
	})
}

func addSeparator(result *strings.Builder, isLast bool) {
	if result.Len() == 0 {
		return
	}

	if isLast {
		result.WriteString(" and ")
	} else {
		result.WriteString(", ")
	}
}

func appendUnit(result *strings.Builder, value int64, unit string) {
	result.WriteString(strconv.FormatInt(value, 10))
	result.WriteByte(' ')
	result.WriteString(unit)

	if value > 1 {
		result.WriteByte('s')
	}
}
