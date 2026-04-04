package review

import "time"

type Result string

const (
	ResultWin  Result = "win"
	ResultLose Result = "lose"
	ResultDraw Result = "draw"
)

type Match struct {
	ID            string
	Date          time.Time
	Opponent      string
	ScoreTeam     int
	ScoreOpponent int
	Result        Result
	PlayerStats   []PlayerMatchStats
}

type PlayerMatchStats struct {
	PlayerID   string
	PlayerName string
	Hits       int
	Runs       int
	RBI        int
	Strikeouts int
	Walks      int
	Errors     int
	Tags       PlayerTags
}

type PlayerTags struct {
	GoodPlay        bool
	BadPlay         bool
	HighlightMoment bool
}
