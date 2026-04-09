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

func TestBuildAIInput(t *testing.T) {
	t.Parallel()

	match := Match{
		ID:            1,
		OpponentName:  "Opponent Team",
		TeamScore:     5,
		OpponentScore: 3,
	}
	mvp := MVPResult{
		PlayerID:   1,
		PlayerName: "Player A",
		Score:      8,
	}
	players := []PlayerReview{
		{
			PlayerID:   1,
			PlayerName: "Player A",
			Stats: PlayerMatchStats{
				Hits: 2,
				RBI:  1,
				Runs: 1,
			},
			Title: "ヒット王",
			IsMVP: true,
		},
		{
			PlayerID:   2,
			PlayerName: "Player B",
			Stats: PlayerMatchStats{
				Hits: 1,
				RBI:  2,
				Runs: 0,
			},
			Title: "活躍選手",
			IsMVP: false,
		},
	}

	got := buildAIInput(match, mvp, players)

	want := AIInput{
		Match: AIMatch{
			ID:            1,
			OpponentName:  "Opponent Team",
			TeamScore:     5,
			OpponentScore: 3,
		},
		MVP: AIMVP{
			PlayerID:   1,
			PlayerName: "Player A",
			Score:      8,
		},
		Players: []AIPlayer{
			{
				PlayerID:   1,
				PlayerName: "Player A",
				Title:      "ヒット王",
				Score:      9,
				IsMVP:      true,
				Stats: AIPlayerStats{
					Hits: 2,
					RBI:  1,
					Runs: 1,
				},
			},
			{
				PlayerID:   2,
				PlayerName: "Player B",
				Title:      "活躍選手",
				Score:      8,
				IsMVP:      false,
				Stats: AIPlayerStats{
					Hits: 1,
					RBI:  2,
					Runs: 0,
				},
			},
		},
	}

	assert.Equal(t, want, got)
}

func TestGenerateAIInputJSON(t *testing.T) {
	t.Parallel()

	match := Match{
		ID:            1,
		OpponentName:  "Opponent Team",
		TeamScore:     5,
		OpponentScore: 3,
	}
	mvp := MVPResult{
		PlayerID:   1,
		PlayerName: "Player A",
		Score:      8,
	}
	players := []PlayerReview{
		{
			PlayerID:   1,
			PlayerName: "Player A",
			Stats: PlayerMatchStats{
				Hits: 2,
				RBI:  1,
				Runs: 1,
			},
			Title: "ヒット王",
			IsMVP: true,
		},
		{
			PlayerID:   2,
			PlayerName: "Player B",
			Stats: PlayerMatchStats{
				Hits: 1,
				RBI:  2,
				Runs: 0,
			},
			Title: "活躍選手",
			IsMVP: false,
		},
	}

	got := GenerateAIInputJSON(match, mvp, players)

	expected := `{
  "match": {
    "id": 1,
    "opponent_name": "Opponent Team",
    "team_score": 5,
    "opponent_score": 3
  },
  "mvp": {
    "player_id": 1,
    "player_name": "Player A",
    "score": 8
  },
  "players": [
    {
      "player_id": 1,
      "player_name": "Player A",
      "title": "ヒット王",
      "score": 9,
      "is_mvp": true,
      "stats": {
        "hits": 2,
        "rbi": 1,
        "runs": 1
      }
    },
    {
      "player_id": 2,
      "player_name": "Player B",
      "title": "活躍選手",
      "score": 8,
      "is_mvp": false,
      "stats": {
        "hits": 1,
        "rbi": 2,
        "runs": 0
      }
    }
  ]
}`

	assert.JSONEq(t, expected, got)
}
