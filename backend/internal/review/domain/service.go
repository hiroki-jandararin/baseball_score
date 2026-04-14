package review

import "time"

// MatchOverview は結果画面に表示するための試合ヘッダー
type MatchOverview struct {
	ID            int
	MatchDate     time.Time
	OpponentName  string
	TeamScore     int
	OpponentScore int
	IsWin         int
}

// MatchReview は1試合分の最終生成結果
type MatchReview struct {
	Match   MatchOverview
	MVP     MVPResult
	Summary MatchSummary
	Players []PlayerReview
}

// 責務は純粋なドメインルールの実行に限定し、次を順番にまとめる。
// 1. 結果画面向けの試合ヘッダーを作る
// 2. 選手成績から MVP を選ぶ
// 3. 試合全体の総評を作る
// 4. 各選手の称号とコメントを作る
// 5. 最終的な MatchReview にまとめる
type Service struct{}

func (s Service) GenerateMatchReview(match Match) MatchReview {
	return MatchReview{}
}
