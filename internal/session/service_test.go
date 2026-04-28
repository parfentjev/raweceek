package session

import (
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
