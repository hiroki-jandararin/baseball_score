package game

type PlayType string

const (
	Strikeout PlayType = "strikeout"
	Flyout    PlayType = "flyout"
	Groundout PlayType = "groundout"
	Walk      PlayType = "walk"
	HitByPitch PlayType = "hitByPitch"
	Single   PlayType = "single"
	Double   PlayType = "double"
	Triple   PlayType = "triple"
	HomeRun  PlayType = "homerun"
	Error	PlayType = "error"
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
		if next.Runner.First && next.Runner.Second && next.Runner.Third {
			next.AddRun(1)
			next.Runner = RunnerState{
				First:  true,
				Second: true,
				Third:  true,
			}
		}else if next.Runner.First && next.Runner.Second {
			next.Runner = RunnerState{
				First:  true,
				Second: true,
				Third:  true,
			}
		}else if next.Runner.First && next.Runner.Third {
			next.Runner = RunnerState{
				First:  true,
				Second: true,
				Third:  true,
			}
		} else if next.Runner.Second && next.Runner.Third {
			next.Runner = RunnerState{
				First:  true,
				Second: true,
				Third:  true,
			}
		} else if next.Runner.First {
			next.Runner = RunnerState{
				First:  true,
				Second: true,
				Third:  false,
			}
		} else if next.Runner.Second {
			next.Runner = RunnerState{
				First:  true,
				Second: true,
				Third:  false,
			}
		} else if next.Runner.Third {
			next.Runner = RunnerState{
				First:  true,
				Second: false,
				Third:  true,
			}
		} else {
			next.Runner.First = true
		}
	case Single:
		if next.Runner.First && next.Runner.Second && next.Runner.Third {
			next.AddRun(1)
			next.Runner = RunnerState{
				First:  true,
				Second: true,
				Third:  true,
			}
		}else if next.Runner.First && next.Runner.Second {
			next.Runner = RunnerState{
				First:  true,
				Second: true,
				Third:  true,
			}
		}else if next.Runner.First && next.Runner.Third {
			next.AddRun(1)
			next.Runner = RunnerState{
				First:  true,
				Second: true,
				Third:  false,
			}
		} else if next.Runner.Second && next.Runner.Third {
			next.AddRun(1)
			next.Runner = RunnerState{
				First:  true,
				Second: false,
				Third:  true,
			}
		} else if next.Runner.First {
			next.Runner = RunnerState{
				First:  true,
				Second: true,
				Third:  false,
			}
		} else if next.Runner.Second {
			next.Runner = RunnerState{
				First:  true,
				Second: false,
				Third:  true,
			}
		} else if next.Runner.Third {
			next.AddRun(1)
			next.Runner = RunnerState{
				First:  true,
				Second: false,
				Third:  false,
			}
		} else {
			next.Runner.First = true
		}
	case Double:
		if next.Runner.First && next.Runner.Second && next.Runner.Third {
			next.AddRun(2)
			next.Runner = RunnerState{
				First:  false,
				Second: true,
				Third:  true,
			}
		}else if next.Runner.First && next.Runner.Second {
			next.AddRun(1)
			next.Runner = RunnerState{
				First:  false,
				Second: true,
				Third:  true,
			}
		}else if next.Runner.First && next.Runner.Third {
			next.AddRun(1)
			next.Runner = RunnerState{
				First:  false,
				Second: true,
				Third:  true,
			}
		} else if next.Runner.Second && next.Runner.Third {
			next.AddRun(2)
			next.Runner = RunnerState{
				First:  false,
				Second: true,
				Third:  false,
			}
		} else if next.Runner.First {
			next.Runner = RunnerState{
				First:  false,
				Second: true,
				Third:  true,
			}
		} else if next.Runner.Second {
			next.AddRun(1)
			next.Runner = RunnerState{
				First:  false,
				Second: true,
				Third:  false,
			}
		} else if next.Runner.Third {
			next.AddRun(1)
			next.Runner = RunnerState{
				First:  false,
				Second: true,
				Third:  false,
			}
		} else {
			next.Runner.Second = true
		}
	case Triple:
		if next.Runner.First && next.Runner.Second && next.Runner.Third {
			next.AddRun(3)
			next.Runner = RunnerState{
				First:  false,
				Second: false,
				Third:  true,
			}
		}else if next.Runner.First && next.Runner.Second {
			next.AddRun(2)
			next.Runner = RunnerState{
				First:  false,
				Second: false,
				Third:  true,
			}
		}else if next.Runner.First && next.Runner.Third {
			next.AddRun(2)
			next.Runner = RunnerState{
				First:  false,
				Second: false,
				Third:  true,
			}
		} else if next.Runner.Second && next.Runner.Third {
			next.AddRun(2)
			next.Runner = RunnerState{
				First:  false,
				Second: false,
				Third:  true,
			}
		} else if next.Runner.First {
			next.AddRun(1)
			next.Runner = RunnerState{
				First:  false,
				Second: false,
				Third:  true,
			}
		} else if next.Runner.Second {
			next.AddRun(1)
			next.Runner = RunnerState{
				First:  false,
				Second: false,
				Third:  true,
			}
		} else if next.Runner.Third {
			next.AddRun(1)
			next.Runner = RunnerState{
				First:  false,
				Second: false,
				Third:  true,
			}
		} else {
			next.Runner.Third = true
		}
	case HomeRun:
		if next.Runner.First && next.Runner.Second && next.Runner.Third {
			next.AddRun(4)
			next.Runner = RunnerState{
				First:  false,
				Second: false,
				Third:  false,
			}
		}else if next.Runner.First && next.Runner.Second {
			next.AddRun(3)
			next.Runner = RunnerState{
				First:  false,
				Second: false,
				Third:  false,
			}
		}else if next.Runner.First && next.Runner.Third {
			next.AddRun(3)
			next.Runner = RunnerState{
				First:  false,
				Second: false,
				Third:  false,
			}
		} else if next.Runner.Second && next.Runner.Third {
			next.AddRun(3)
			next.Runner = RunnerState{
				First:  false,
				Second: false,
				Third:  false,
			}
		} else if next.Runner.First {
			next.AddRun(2)
			next.Runner = RunnerState{
				First:  false,
				Second: false,
				Third:  false,
			}
		} else if next.Runner.Second {
			next.AddRun(2)
			next.Runner = RunnerState{
				First:  false,
				Second: false,
				Third:  false,
			}
		} else if next.Runner.Third {
			next.AddRun(2)
			next.Runner = RunnerState{
				First:  false,
				Second: false,
				Third:  false,
			}
		} else {
			next.AddRun(1)
			next.Runner = RunnerState{
				First:  false,
				Second: false,
				Third:  false,
			}
		}
		case Error:
		if next.Runner.First && next.Runner.Second && next.Runner.Third {
			next.AddRun(1)
			next.Runner = RunnerState{
				First:  true,
				Second: true,
				Third:  true,
			}
		}else if next.Runner.First && next.Runner.Second {
			next.Runner = RunnerState{
				First:  true,
				Second: true,
				Third:  true,
			}
		}else if next.Runner.First && next.Runner.Third {
			next.AddRun(1)
			next.Runner = RunnerState{
				First:  true,
				Second: true,
				Third:  false,
			}
		} else if next.Runner.Second && next.Runner.Third {
			next.AddRun(1)
			next.Runner = RunnerState{
				First:  true,
				Second: false,
				Third:  true,
			}
		} else if next.Runner.First {
			next.Runner = RunnerState{
				First:  true,
				Second: true,
				Third:  false,
			}
		} else if next.Runner.Second {
			next.Runner = RunnerState{
				First:  true,
				Second: false,
				Third:  true,
			}
		} else if next.Runner.Third {
			next.AddRun(1)
			next.Runner = RunnerState{
				First:  true,
				Second: false,
				Third:  false,
			}
		} else {
			next.Runner.First = true
		}		
	}

	return &next
}
