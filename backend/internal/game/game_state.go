package game

type TeamID string

type GameState struct {
	Inning              int
	Outs                int
	FirstBattingTeamID  TeamID
	SecondBattingTeamID TeamID
	Teams			   map[TeamID]*TeamState
	InningHalf          InningHalf
	Runner              RunnerState
}

type RunnerState struct {
	First  bool
	Second bool
	Third  bool
}

type PlayerID string

type TeamState struct {
	TeamID TeamID
	Score  int
}

type InningHalf string

const (
	Top    InningHalf = "top"
	Bottom InningHalf = "bottom"
)

func NewGameState() *GameState {
	return &GameState{
		Inning:             1,
		Outs:               0,
		FirstBattingTeamID:  "team1",
		SecondBattingTeamID: "team2",
		Teams: map[TeamID]*TeamState{
			"team1": {TeamID: "team1", Score: 0},
			"team2": {TeamID: "team2", Score: 0},
		},
		InningHalf:         Top,
		Runner: RunnerState{
			First:  false,
			Second: false,
			Third:  false,
		},
	}
}

func (s *GameState) CurrentBattingTeamID() TeamID {
	if s.InningHalf == Top {
		return s.FirstBattingTeamID
	}

	return s.SecondBattingTeamID
}

func (s *GameState) AddRun(runs int) {
	if s.InningHalf == Top {
		s.Teams[s.FirstBattingTeamID].Score += runs
		return
	}

	s.Teams[s.SecondBattingTeamID].Score += runs
}