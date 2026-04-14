package review

import "encoding/json"

type PlayerReview struct {
	PlayerID      int
	PlayerName    string
	Stats         PlayerMatchStats
	Title         string
	Comment       string
	IsMVP         bool
}

type TitleRule struct {
	Condition func(PlayerMatchStats) bool
	Title     string
}

var (
	TitleHitKing         = "ヒット王"
	TitleRBIKing         = "打点王"
	TitleRunKing         = "盗塁王"
	TitleKKing           = "三振王"
	TitleGoodPlay        = "Good Play"
	TitleHighlightMoment = "流れを変えた男"
	TitleBaseResident    = "塁に住んでる人"
	TitleDreamFirstClass = "夢だけは一級品"
	TitleNextExpectation = "次に期待枠"
	TitleWarmupPending   = "アップ完了待ち"
	TitleMainCharacter   = "本日の主役"
	TitleWorker          = "仕事人"
	TitleDreamChaser     = "夢追い人2026"
	TitleScriptWriter    = "試合の脚本書いた人"
	TitleObservationDay  = "今日は様子見回"
	TitleWarmupLong      = "ウォーミングアップ長め勢"
	TitleActivePlayer    = "活躍選手"
)

var titleRules = []TitleRule{
	{Condition: func(s PlayerMatchStats) bool { return s.Hits >= 3 }, Title: TitleHitKing},
	{Condition: func(s PlayerMatchStats) bool { return s.RBI >= 3 }, Title: TitleRBIKing},
	{Condition: func(s PlayerMatchStats) bool { return s.Runs >= 2 }, Title: TitleRunKing},
	{Condition: func(s PlayerMatchStats) bool { return s.Strikeouts >= 3 }, Title: TitleKKing},
	{Condition: func(s PlayerMatchStats) bool { return s.GoodPlay == 1 }, Title: TitleGoodPlay},
	{Condition: func(s PlayerMatchStats) bool { return s.HighlightMoment == 1 }, Title: TitleHighlightMoment},
	{Condition: func(s PlayerMatchStats) bool { return s.Hits >= 2 && s.RBI >= 2 }, Title: TitleMainCharacter},
	{Condition: func(s PlayerMatchStats) bool { return s.RBI == 2 }, Title: TitleScriptWriter},
	{Condition: func(s PlayerMatchStats) bool { return s.AtBats == 1 && s.Hits == 0 && s.Walks == 1 }, Title: TitleWarmupPending},
	{Condition: func(s PlayerMatchStats) bool { return s.Walks == 1 && s.AtBats > 0 && s.Hits == 0 }, Title: TitleBaseResident},
	{Condition: func(s PlayerMatchStats) bool { return s.AtBats >= 3 && s.Hits >= 1 && s.RBI <= 1 && s.Runs <= 1 }, Title: TitleWorker},
	{Condition: func(s PlayerMatchStats) bool { return s.AtBats >= 3 && s.Hits == 0 && s.Strikeouts < 2 }, Title: TitleNextExpectation},
	{Condition: func(s PlayerMatchStats) bool { return s.AtBats <= 2 && s.Hits == 0 && s.Walks == 0 && s.Strikeouts == 0 && s.AtBats > 0 }, Title: TitleObservationDay},
	{Condition: func(s PlayerMatchStats) bool { return s.Strikeouts >= 2 && s.Hits == 0 }, Title: TitleWarmupLong},
	{Condition: func(s PlayerMatchStats) bool { return s.AtBats > 0 }, Title: TitleDreamChaser},
}

// AssignTitle assigns a title based on player stats
func AssignTitle(stats PlayerMatchStats) string {
	for _, rule := range titleRules {
		if rule.Condition(stats) {
			return rule.Title
		}
	}
	return TitleActivePlayer
}

func calculateBasePlayerScore(stats PlayerMatchStats) int {
	score := 0
	score += stats.Hits * 2
	score += stats.RBI * 3
	score += stats.Runs * 2
	score += stats.Walks
	score -= stats.Errors * 2
	score -= stats.Strikeouts
	return score
}

func CalculatePlayerScore(stats PlayerMatchStats) int {
	return calculateBasePlayerScore(stats)
}
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