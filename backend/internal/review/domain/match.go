package review

import "time"

type Match struct {
	ID            int
	TeamID        int
	OpponentName  string
	MatchDate     time.Time
	Location      string
	IsWin         int
	TeamScore     int
	OpponentScore int
	Note          string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	PlayerStats   []PlayerMatchStats
}

type PlayerMatchStats struct {
	ID            int
	MatchID       int
	PlayerID      int
	PlayerName    string
	BattingOrder  int
	Position      string
	Hits          int
	AtBats        int
	RBI           int
	Runs          int
	Walks         int
	Strikeouts    int
	Errors        int
	GoodPlay      int
	HighlightMoment int
	Memo          string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
