package review

import "encoding/json"



type AIInput struct {
	Match   AIMatch    `json:"match"`
	MVP     AIMVP      `json:"mvp"`
	Players []AIPlayer `json:"players"`
}

type AIMatch struct {
	ID            int    `json:"id"`
	OpponentName  string `json:"opponent_name"`
	TeamScore     int    `json:"team_score"`
	OpponentScore int    `json:"opponent_score"`
}

type AIMVP struct {
	PlayerID   int    `json:"player_id"`
	PlayerName string `json:"player_name"`
	Score      int    `json:"score"`
}

type AIPlayer struct {
	PlayerID   int              `json:"player_id"`
	PlayerName string           `json:"player_name"`
	Title      string           `json:"title"`
	Score      int              `json:"score"`
	IsMVP      bool             `json:"is_mvp"`
	Stats      AIPlayerStats `json:"stats"`
}

type AIPlayerStats struct {
	Hits int `json:"hits"`
	RBI  int `json:"rbi"`
	Runs int `json:"runs"`
}

func buildAIInput(match Match, mvp MVPResult, players []PlayerReview) AIInput {
	aiPlayers := make([]AIPlayer, 0, len(players))
	for _, player := range players {
		aiPlayers = append(aiPlayers, AIPlayer{
			PlayerID:   player.PlayerID,
			PlayerName: player.PlayerName,
			Title:      player.Title,
			Score:      CalculatePlayerScore(player.Stats),
			IsMVP:      player.IsMVP,
			Stats: AIPlayerStats{
				Hits: player.Stats.Hits,
				RBI:  player.Stats.RBI,
				Runs: player.Stats.Runs,
			},
		})
	}

	return AIInput{
		Match: AIMatch{
			ID:            match.ID,
			OpponentName:  match.OpponentName,
			TeamScore:     match.TeamScore,
			OpponentScore: match.OpponentScore,
		},
		MVP: AIMVP{
			PlayerID:   mvp.PlayerID,
			PlayerName: mvp.PlayerName,
			Score:      mvp.Score,
		},
		Players: aiPlayers,
	}
}

func GenerateAIInputJSON(match Match, mvp MVPResult, players []PlayerReview) string {
	input := buildAIInput(match, mvp, players)
	jsonBytes, err := json.MarshalIndent(input, "", "  ")
	if err != nil {
		return "{}"
	}
	return string(jsonBytes)
}