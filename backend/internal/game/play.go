package game

type PlayType string

const (
	Strikeout PlayType = "strikeout"
)

func ApplyPlay(state *GameState, play Play) *GameState {
	next := *state

	switch play.Type {
	case Strikeout:
		next.Outs++
		if next.Outs >= 3 {
			next.Outs = 0
			if next.InningHalf == Top {
				next.InningHalf = Bottom
				next.Runner = RunnerState{}
			} else {
				next.InningHalf = Top
				next.Inning++
				next.Runner = RunnerState{}
			}
		}
	}

	return &next
}