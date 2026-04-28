package session

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateCeeksType(t *testing.T) {
	expected := "CEEKS"
	actual := calculateCeeks(1)

	assert.Equal(t, expected, actual.Type)
}

func TestCalculateCeeksValue(t *testing.T) {
	tests := []struct {
		name          string
		remainingTime int64
		expected      string
	}{
		{"input: -1s", -1, "0.001 ceeks"},
		{"input: 0s", 0, "0.001 ceeks"},
		{"input: 1 minute", 60, "0.001 ceeks"},
		{"input: 10 minutes", 600, "0.001 ceeks"},
		{"input: 30 minutes", 1800, "0.003 ceeks"},
		{"input: 1 hour", 3600, "0.006 ceeks"},
		{"input: 12 hours", 43200, "0.071 ceeks"},
		{"input: 24 hours", 86400, "0.143 ceeks"},
		{"output: just below 1 ceek", 604200, "0.999 ceeks"},
		{"output: exactly 1 ceek", 604800, "1.000 ceeks"},
		{"output: just above 1 ceek", 605400, "1.001 ceeks"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := calculateCeeks(test.remainingTime)
			assert.Equal(t, test.expected, actual.Value)
		})
	}
}

func TestCalculateTimeUntilType(t *testing.T) {
	expected := "TIME_UNTIL"
	actual := calculateTimeUntil(1)

	assert.Equal(t, expected, actual.Type)
}

func TestCalculateTimeUntilValue(t *testing.T) {
	tests := []struct {
		name          string
		remainingTime int64
		expected      string
	}{
		{"input: -1s", -1, "-1 seconds"},
		{"input: 0s", 0, "0 seconds"},
		{"input: 1s", 1, "1 second"},
		{"input: 59s", 59, "59 seconds"},
		{"input: 60s", 60, "1 minute and 0 seconds"},
		{"input: 61s", 61, "1 minute and 1 second"},
		{"input: 119s", 119, "1 minute and 59 seconds"},
		{"input: 120s", 120, "2 minutes and 0 seconds"},
		{"input: 1h", 3600, "1 hour and 0 seconds"},
		{"input: 1h 1m", 3660, "1 hour 1 minute and 0 seconds"},
		{"input: 23h 59m 59s", 86399, "23 hours 59 minutes and 59 seconds"},
		{"input: 24h", 86400, "1 day and 0 seconds"},
		{"input: 6d", 518400, "6 days and 0 seconds"},
		{"input: 6d 23h 59m 59s", 604799, "6 days 23 hours 59 minutes and 59 seconds"},
		{"input: 7d", 604800, "1 week and 0 seconds"},
		{"input: 7d 1s", 604801, "1 week and 1 second"},
		{"input: 13d 23h 59m 59s", 1209599, "1 week 6 days 23 hours 59 minutes and 59 seconds"},
		{"input: 29d", 2505600, "4 weeks 1 day and 0 seconds"},
		{"input: 30d", 2592000, "1 month and 0 seconds"},
		{"input: 31d", 2678400, "1 month 1 day and 0 seconds"},
		{"input: 1h 0m 1s", 3601, "1 hour and 1 second"},
		{"input: 1d 0h 1m", 86460, "1 day 1 minute and 0 seconds"},
		{"input: 59d", 5097600, "1 month 4 weeks 1 day and 0 seconds"},
		{"input: 60d", 5184000, "2 months and 0 seconds"},
		{"input: 61d", 5270400, "2 months 1 day and 0 seconds"},
		{"input: 89d", 7689600, "2 months 4 weeks 1 day and 0 seconds"},
		{"input: 90d", 7776000, "3 months and 0 seconds"},
		{"input: 91d", 7862400, "3 months 1 day and 0 seconds"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := calculateTimeUntil(test.remainingTime)
			assert.Equal(t, test.expected, actual.Value)
		})
	}
}

func TestAppendTimeUnit(t *testing.T) {
	tests := []struct {
		name     string
		value    int64
		unitName string
		end      string
		expected string
	}{
		{"output: singular", 1, "second", "", "1 second"},
		{"output: plural (zero)", 0, "second", "", "0 seconds"},
		{"output: plural (two)", 2, "second", "", "2 seconds"},
		{"output: negative", -1, "second", "", "-1 seconds"},
		{"output: end appended", 1, "second", " ", "1 second "},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var sb strings.Builder
			appendTimeUnit(&sb, test.value, test.unitName, test.end)
			assert.Equal(t, test.expected, sb.String())
		})
	}
}
