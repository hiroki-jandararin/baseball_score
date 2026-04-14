package review

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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