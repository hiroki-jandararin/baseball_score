package review

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestScoreMVPPlayer(t *testing.T) {
	t.Parallel()

	match := Match{
		IsWin: 1,
	}

	tests := []struct {
		name  string
		stats PlayerMatchStats
		want  int
	}{
		{
			name: "adds hits, rbi, runs, walks",
			stats: PlayerMatchStats{
				Hits: 2,
				RBI:  1,
				Runs: 1,
				Walks: 1,
			},
			want: 10,
		},
		{
			name: "subtracts errors and strikeouts",
			stats: PlayerMatchStats{
				Errors:     2,
				Strikeouts: 3,
			},
			want: -7,
		},
		{
			name: "counts every rule together",
			stats: PlayerMatchStats{
				Hits:       1,
				RBI:        2,
				Runs:       3,
				Walks:      1,
				Errors:     1,
				Strikeouts: 1,
			},
			want: 12,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := ScoreMVP(match, tt.stats)
			assert.Equal(t, tt.want, got, "ScoreMVP() = %d, want %d", got, tt.want)
		})
	}
}

func TestSelectMVP(t *testing.T) {
	t.Parallel()

	match := Match{
		ID:           1,
		MatchDate:    time.Date(2026, time.April, 4, 0, 0, 0, 0, time.UTC),
		OpponentName: "Rivals",
		IsWin:        1,
		PlayerStats: []PlayerMatchStats{
			{PlayerID: 1, PlayerName: "A", Hits: 1},
			{PlayerID: 2, PlayerName: "B", Hits: 3},
			{PlayerID: 3, PlayerName: "C", Hits: 2, RBI: 1},
		},
	}

	got := SelectMVP(match)

	assert.Equal(t, 3, got.PlayerID, "SelectMVP() PlayerID = %d, want %d", got.PlayerID, 3)
	assert.Equal(t, "C", got.PlayerName, "SelectMVP() PlayerName = %q, want %q", got.PlayerName, "C")
	assert.Equal(t, 7, got.Score, "SelectMVP() Score = %d, want %d", got.Score, 7)
}

func TestSelectMVPWithTie(t *testing.T) {
	t.Parallel()
	
	match := Match{
		ID:           1,
		MatchDate:    time.Date(2026, time.April, 4, 0, 0, 0, 0, time.UTC),
		OpponentName: "Rivals",
		IsWin:        1,
		PlayerStats: []PlayerMatchStats{
			{PlayerID: 1, PlayerName: "A", Hits: 1},
			{PlayerID: 2, PlayerName: "B", Hits: 1},
		},
	}

	got := SelectMVP(match)
	assert.Equal(t, 1, got.PlayerID, "SelectMVP() PlayerID = %d, want %d", got.PlayerID, 1)
}