package game

type GameState struct {
	Inning      int
	Outs        int
	HomeScore   int
	AwayScore   int
	InningHalf  InningHalf
	Runner RunnerState
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
		Inning:     1,
		Outs:       0,
		HomeScore:  0,
		AwayScore:  0,
		InningHalf: Top,
		Runner: RunnerState{
			First:  false,
			Second: false,
			Third:  false,
		},
	}
}