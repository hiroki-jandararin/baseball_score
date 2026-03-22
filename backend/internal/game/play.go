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
	SacrificeBunt PlayType = "sacrificeBunt"
	Steal PlayType = "steal"
	// TODO ランナーの状態を持てるようにして固定ルール以外のプレイも記録できるようにする
	// FieldersChoice PlayType = "fieldersChoice"
	// DoublePlay PlayType = "doublePlay"
)
func RecordPlay(state *GameState, play Play) *GameState {
	//TODO ValidatePlay(state, play) を追加する
	return applyPlay(state, play)
}

func applyPlay(state *GameState, play Play) *GameState {
	next := *state
	switch play.Type {
	case Strikeout, Flyout, Groundout:
		if addOut(&next) {
			return &next
		}
	case Walk, HitByPitch:
		handleWalk(&next)
	case Single:
		handleSingle(&next)
	case Double:
		handleDouble(&next)
	case Triple:
		handleTriple(&next)
	case HomeRun:
		handleHomeRun(&next)
	case Error:
		handleError(&next)
	case SacrificeBunt:
		if addSacrificeBunt(&next) {
			return &next
		}
	case Steal:
		handleSteal(&next)
	}
	return &next
}

func addOut(s *GameState) bool {
	s.Outs++
	if s.Outs >= 3 {
		s.ChangeInning()
		return true
	}
	return false
}

func handleWalk(s *GameState) {
	if s.Runner.First && s.Runner.Second && s.Runner.Third {
		s.AddRun(1)
		s.Runner = RunnerState{First: true, Second: true, Third: true}
	} else if s.Runner.First && s.Runner.Second {
		s.Runner = RunnerState{First: true, Second: true, Third: true}
	} else if s.Runner.First && s.Runner.Third {
		s.Runner = RunnerState{First: true, Second: true, Third: true}
	} else if s.Runner.Second && s.Runner.Third {
		s.Runner = RunnerState{First: true, Second: true, Third: true}
	} else if s.Runner.First {
		s.Runner = RunnerState{First: true, Second: true, Third: false}
	} else if s.Runner.Second {
		s.Runner = RunnerState{First: true, Second: true, Third: false}
	} else if s.Runner.Third {
		s.Runner = RunnerState{First: true, Second: false, Third: true}
	} else {
		s.Runner.First = true
	}
}

func handleSingle(s *GameState) {
	if s.Runner.First && s.Runner.Second && s.Runner.Third {
		s.AddRun(1)
		s.Runner = RunnerState{First: true, Second: true, Third: true}
	} else if s.Runner.First && s.Runner.Second {
		s.Runner = RunnerState{First: true, Second: true, Third: true}
	} else if s.Runner.First && s.Runner.Third {
		s.AddRun(1)
		s.Runner = RunnerState{First: true, Second: true, Third: false}
	} else if s.Runner.Second && s.Runner.Third {
		s.AddRun(1)
		s.Runner = RunnerState{First: true, Second: false, Third: true}
	} else if s.Runner.First {
		s.Runner = RunnerState{First: true, Second: true, Third: false}
	} else if s.Runner.Second {
		s.Runner = RunnerState{First: true, Second: false, Third: true}
	} else if s.Runner.Third {
		s.AddRun(1)
		s.Runner = RunnerState{First: true, Second: false, Third: false}
	} else {
		s.Runner.First = true
	}
}

func handleDouble(s *GameState) {
	if s.Runner.First && s.Runner.Second && s.Runner.Third {
		s.AddRun(2)
		s.Runner = RunnerState{First: false, Second: true, Third: true}
	} else if s.Runner.First && s.Runner.Second {
		s.AddRun(1)
		s.Runner = RunnerState{First: false, Second: true, Third: true}
	} else if s.Runner.First && s.Runner.Third {
		s.AddRun(1)
		s.Runner = RunnerState{First: false, Second: true, Third: true}
	} else if s.Runner.Second && s.Runner.Third {
		s.AddRun(2)
		s.Runner = RunnerState{First: false, Second: true, Third: false}
	} else if s.Runner.First {
		s.Runner = RunnerState{First: false, Second: true, Third: true}
	} else if s.Runner.Second {
		s.AddRun(1)
		s.Runner = RunnerState{First: false, Second: true, Third: false}
	} else if s.Runner.Third {
		s.AddRun(1)
		s.Runner = RunnerState{First: false, Second: true, Third: false}
	} else {
		s.Runner.Second = true
	}
}

func handleTriple(s *GameState) {
	if s.Runner.First && s.Runner.Second && s.Runner.Third {
		s.AddRun(3)
		s.Runner = RunnerState{First: false, Second: false, Third: true}
	} else if s.Runner.First && s.Runner.Second {
		s.AddRun(2)
		s.Runner = RunnerState{First: false, Second: false, Third: true}
	} else if s.Runner.First && s.Runner.Third {
		s.AddRun(2)
		s.Runner = RunnerState{First: false, Second: false, Third: true}
	} else if s.Runner.Second && s.Runner.Third {
		s.AddRun(2)
		s.Runner = RunnerState{First: false, Second: false, Third: true}
	} else if s.Runner.First {
		s.AddRun(1)
		s.Runner = RunnerState{First: false, Second: false, Third: true}
	} else if s.Runner.Second {
		s.AddRun(1)
		s.Runner = RunnerState{First: false, Second: false, Third: true}
	} else if s.Runner.Third {
		s.AddRun(1)
		s.Runner = RunnerState{First: false, Second: false, Third: true}
	} else {
		s.Runner.Third = true
	}
}

func handleHomeRun(s *GameState) {
	if s.Runner.First && s.Runner.Second && s.Runner.Third {
		s.AddRun(4)
		s.Runner = RunnerState{First: false, Second: false, Third: false}
	} else if s.Runner.First && s.Runner.Second {
		s.AddRun(3)
		s.Runner = RunnerState{First: false, Second: false, Third: false}
	} else if s.Runner.First && s.Runner.Third {
		s.AddRun(3)
		s.Runner = RunnerState{First: false, Second: false, Third: false}
	} else if s.Runner.Second && s.Runner.Third {
		s.AddRun(3)
		s.Runner = RunnerState{First: false, Second: false, Third: false}
	} else if s.Runner.First {
		s.AddRun(2)
		s.Runner = RunnerState{First: false, Second: false, Third: false}
	} else if s.Runner.Second {
		s.AddRun(2)
		s.Runner = RunnerState{First: false, Second: false, Third: false}
	} else if s.Runner.Third {
		s.AddRun(2)
		s.Runner = RunnerState{First: false, Second: false, Third: false}
	} else {
		s.AddRun(1)
		s.Runner = RunnerState{First: false, Second: false, Third: false}
	}
}

func handleError(s *GameState) {
	if s.Runner.First && s.Runner.Second && s.Runner.Third {
		s.AddRun(1)
		s.Runner = RunnerState{First: true, Second: true, Third: true}
	} else if s.Runner.First && s.Runner.Second {
		s.Runner = RunnerState{First: true, Second: true, Third: true}
	} else if s.Runner.First && s.Runner.Third {
		s.AddRun(1)
		s.Runner = RunnerState{First: true, Second: true, Third: false}
	} else if s.Runner.Second && s.Runner.Third {
		s.AddRun(1)
		s.Runner = RunnerState{First: true, Second: false, Third: true}
	} else if s.Runner.First {
		s.Runner = RunnerState{First: true, Second: true, Third: false}
	} else if s.Runner.Second {
		s.Runner = RunnerState{First: true, Second: false, Third: true}
	} else if s.Runner.Third {
		s.AddRun(1)
		s.Runner = RunnerState{First: true, Second: false, Third: false}
	} else {
		s.Runner.First = true
	}
}

func addSacrificeBunt(s *GameState) bool {
	// returns true if inning changed and we should return early
	if s.Runner.First && s.Runner.Second && s.Runner.Third {
		if addOut(s) {
			return true
		}
		s.AddRun(1)
		s.Runner = RunnerState{First: false, Second: true, Third: true}
		return false
	} else if s.Runner.First && s.Runner.Second {
		if addOut(s) {
			return true
		}
		s.Runner = RunnerState{First: false, Second: true, Third: true}
		return false
	} else if s.Runner.First && s.Runner.Third {
		if addOut(s) {
			return true
		}
		s.AddRun(1)
		s.Runner = RunnerState{First: false, Second: true, Third: false}
		return false
	} else if s.Runner.Second && s.Runner.Third {
		if addOut(s) {
			return true
		}
		s.AddRun(1)
		s.Runner = RunnerState{First: false, Second: false, Third: true}
		return false
	} else if s.Runner.First {
		if addOut(s) {
			return true
		}
		s.Runner = RunnerState{First: false, Second: true, Third: false}
		return false
	} else if s.Runner.Second {
		if addOut(s) {
			return true
		}
		s.Runner = RunnerState{First: false, Second: false, Third: true}
		return false
	} else if s.Runner.Third {
		if addOut(s) {
			return true
		}
		s.AddRun(1)
		s.Runner = RunnerState{First: false, Second: false, Third: false}
		return false
	} else {
		if addOut(s) {
			return true
		}
		return false
	}
}

func handleSteal(s *GameState) {
	if s.Runner.First && s.Runner.Second && s.Runner.Third {
		s.AddRun(1)
		s.Runner = RunnerState{First: false, Second: true, Third: true}
	} else if s.Runner.First && s.Runner.Second {
		s.Runner = RunnerState{First: false, Second: true, Third: true}
	} else if s.Runner.First && s.Runner.Third {
		s.AddRun(1)
		s.Runner = RunnerState{First: false, Second: true, Third: false}
	} else if s.Runner.Second && s.Runner.Third {
		s.AddRun(1)
		s.Runner = RunnerState{First: false, Second: false, Third: true}
	} else if s.Runner.First {
		s.Runner = RunnerState{First: false, Second: true, Third: false}
	} else if s.Runner.Second {
		s.Runner = RunnerState{First: false, Second: false, Third: true}
	} else if s.Runner.Third {
		s.AddRun(1)
		s.Runner = RunnerState{First: false, Second: false, Third: false}
	}
}

func (s *GameState) ChangeInning() {
	s.Outs = 0
	if s.InningHalf == Top {
		s.InningHalf = Bottom
		s.Runner = RunnerState{}
	} else {
		s.InningHalf = Top
		s.Inning++
		s.Runner = RunnerState{}
	}
}