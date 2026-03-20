package game

type TeamID string

type GameState struct {
	Inning              int
	Outs                int
	FirstBattingTeamID  TeamID
	SecondBattingTeamID TeamID
	FirstBattingScore   int
	SecondBattingScore  int
	InningHalf          InningHalf
	Runner              RunnerState
}

type RunnerState struct {
	First  bool
	Second bool
	Third  bool
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
		FirstBattingScore:  0,
		SecondBattingScore: 0,
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
