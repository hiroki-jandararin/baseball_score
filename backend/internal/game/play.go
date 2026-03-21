package game

type PlayType string

const (
	Strikeout PlayType = "strikeout"
	Flyout    PlayType = "flyout"
	Groundout PlayType = "groundout"
	Walk      PlayType = "walk"
	HitByPitch PlayType = "hitByPitch"
	Single   PlayType = "single"
	// Double   PlayType = "double"
	// Triple   PlayType = "triple"
	// HomeRun  PlayType = "homeRun"
	// Error	PlayType = "error"
	// SacrificeBunt PlayType = "sacrificeBunt"
	// Steal PlayType = "steal"
	// FieldersChoice PlayType = "fieldersChoice"
	// DoublePlay PlayType = "doublePlay"
)

func ApplyPlay(state *GameState, play Play) *GameState {
	next := *state

	switch play.Type {
	case Strikeout, Flyout, Groundout:
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
	case Walk, HitByPitch:
		if !next.Runner.First {
			next.Runner.First = true
		} else if next.Runner.First && !next.Runner.Second {
			next.Runner.Second = true
		} else if next.Runner.First && next.Runner.Second && !next.Runner.Third {
			next.Runner.Third = true
		} else if next.Runner.First && next.Runner.Second && next.Runner.Third {
			next.AddRun(1)
		}
	}

	return &next
}
