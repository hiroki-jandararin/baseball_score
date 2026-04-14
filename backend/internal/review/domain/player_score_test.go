package review

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculatePlayerScore(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		stats PlayerMatchStats
		want  int
	}{
		{
			name: "high performer",
			stats: PlayerMatchStats{
				Hits:   3,
				RBI:    2,
				Runs:   2,
				Walks:  1,
				Errors: 0,
			},
			want: 17, // 3*2 + 2*3 + 2*2 + 1 - 0*2 = 6 + 6 + 4 + 1 = 17
		},
		{
			name: "average player",
			stats: PlayerMatchStats{
				Hits:   1,
				RBI:    1,
				Runs:   1,
				Walks:  0,
				Errors: 1,
			},
			want: 5, // 1*2 + 1*3 + 1*2 + 0 - 1*2 = 2 + 3 + 2 - 2 = 5
		},
		{
			name: "low performer",
			stats: PlayerMatchStats{
				Hits:       0,
				RBI:        0,
				Runs:       0,
				Walks:      0,
				Errors:     2,
				Strikeouts: 3,
			},
			want: -7, // 0 + 0 + 0 + 0 - 2*2 - 3 = -4 - 3 = -7
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := CalculatePlayerScore(tt.stats)
			assert.Equal(t, tt.want, got, "CalculatePlayerScore() = %d, want %d", got, tt.want)
		})
	}
}