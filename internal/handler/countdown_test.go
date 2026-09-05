package handler

import (
	"testing"
	"time"

	"github.com/parfentjev/raweceek/internal/generated/api"
)

func TestNewCeeksCountdown(t *testing.T) {
	tests := []struct {
		name          string
		remainingTime time.Duration
		wantValue     string
	}{
		{
			name:          "formats exact weeks",
			remainingTime: 14 * 24 * time.Hour,
			wantValue:     "2.000 ceeks",
		},
		{
			name:          "rounds to three decimal places",
			remainingTime: 10 * 24 * time.Hour,
			wantValue:     "1.429 ceeks",
		},
		{
			name:          "enforces minimum for zero duration",
			remainingTime: 0,
			wantValue:     "0.001 ceeks",
		},
		{
			name:          "enforces minimum for short duration",
			remainingTime: time.Second,
			wantValue:     "0.001 ceeks",
		},
		{
			name:          "enforces minimum for negative duration",
			remainingTime: -time.Second,
			wantValue:     "0.001 ceeks",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			countdown := NewCeeksCountdown(tt.remainingTime)

			if countdown.Type != api.CEEKS {
				t.Errorf("NewCeeks() type = %q, want %q", countdown.Type, api.CEEKS)
			}

			if countdown.Value != tt.wantValue {
				t.Errorf("NewCeeks() value = %q, want %q", countdown.Value, tt.wantValue)
			}
		})
	}
}

func TestNewTimeUntilCountdown(t *testing.T) {
	tests := []struct {
		name          string
		remainingTime time.Duration
		wantValue     string
	}{
		{
			name: "formats all units",
			remainingTime: 77*24*time.Hour + 4*time.Hour +
				5*time.Minute + 6*time.Second,
			wantValue: "2 months, 2 weeks, 3 days, 4 hours, 5 minutes and 6 seconds",
		},
		{
			name: "uses singular units",
			remainingTime: 38*24*time.Hour + time.Hour +
				time.Minute + time.Second,
			wantValue: "1 month, 1 week, 1 day, 1 hour, 1 minute and 1 second",
		},
		{
			name:          "omits zero units",
			remainingTime: 30*24*time.Hour + 2*time.Second,
			wantValue:     "1 month and 2 seconds",
		},
		{
			name:          "uses and when last non-zero unit is not seconds",
			remainingTime: time.Hour + time.Minute,
			wantValue:     "1 hour and 1 minute",
		},
		{
			name:          "formats zero duration as zero",
			remainingTime: 0,
			wantValue:     "0s",
		},
		{
			name:          "formats sub-second duration as zero",
			remainingTime: 999 * time.Millisecond,
			wantValue:     "0s",
		},
		{
			name:          "formats negative duration as zero",
			remainingTime: -time.Second,
			wantValue:     "0s",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			countdown := NewTimeUntilCountdown(tt.remainingTime)

			if countdown.Type != api.TIMEUNTIL {
				t.Errorf("NewTimeUntil() type = %q, want %q", countdown.Type, api.TIMEUNTIL)
			}

			if countdown.Value != tt.wantValue {
				t.Errorf("NewTimeUntil() value = %q, want %q", countdown.Value, tt.wantValue)
			}
		})
	}
}
