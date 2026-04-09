package review

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
