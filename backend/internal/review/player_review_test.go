package review

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssignTitle(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		stats PlayerMatchStats
		want  string
	}{
		{
			name: "hits 3 or more",
			stats: PlayerMatchStats{
				Hits: 3,
			},
			want: "ヒット王",
		},
		{
			name: "rbi 3 or more",
			stats: PlayerMatchStats{
				RBI: 3,
			},
			want: "打点王",
		},
		{
			name: "runs 2 or more",
			stats: PlayerMatchStats{
				Runs: 2,
			},
			want: "盗塁王",
		},
		{
			name: "strikeouts 3 or more",
			stats: PlayerMatchStats{
				Strikeouts: 3,
			},
			want: "三振王",
		},
		{
			name: "good play",
			stats: PlayerMatchStats{
				GoodPlay: 1,
			},
			want: TitleGoodPlay,
		},
		{
			name: "base resident",
			stats: PlayerMatchStats{
				Walks: 1,
				AtBats: 2,
			},
			want: TitleBaseResident,
		},
		{
			name: "script writer",
			stats: PlayerMatchStats{
				RBI: 2,
			},
			want: TitleScriptWriter,
		},
		{
			name: "worker",
			stats: PlayerMatchStats{
				AtBats: 3,
				Hits:   1,
				RBI:    1,
				Runs:   1,
			},
			want: TitleWorker,
		},
		{
			name: "next expectation",
			stats: PlayerMatchStats{
				AtBats:      3,
				Hits:        0,
				Strikeouts:  1,
			},
			want: TitleNextExpectation,
		},
		{
			name: "warmup pending",
			stats: PlayerMatchStats{
				AtBats: 1,
				Hits:   0,
				Walks:  1,
			},
			want: TitleWarmupPending,
		},
		{
			name: "observation day",
			stats: PlayerMatchStats{
				AtBats:      2,
				Hits:        0,
				Walks:       0,
				Strikeouts:  0,
			},
			want: TitleObservationDay,
		},
		{
			name: "warmup long",
			stats: PlayerMatchStats{
				AtBats:      3,
				Hits:        0,
				Strikeouts:  2,
			},
			want: TitleWarmupLong,
		},
		{
			name: "dream chaser",
			stats: PlayerMatchStats{
				AtBats: 1,
				Hits:   1,
			},
			want: TitleDreamChaser,
		},
		{
			name: "default",
			stats: PlayerMatchStats{},
			want: TitleActivePlayer,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := AssignTitle(tt.stats)
			assert.Equal(t, tt.want, got, "AssignTitle() = %q, want %q", got, tt.want)
		})
	}
}