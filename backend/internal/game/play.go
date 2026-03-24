package game

type PlayType string

const (
	Strikeout     PlayType = "strikeout"
	Flyout        PlayType = "flyout"
	Groundout     PlayType = "groundout"
	Walk          PlayType = "walk"
	HitByPitch    PlayType = "hitByPitch"
	Single        PlayType = "single"
	Double        PlayType = "double"
	Triple        PlayType = "triple"
	HomeRun       PlayType = "homerun"
	Error         PlayType = "error"
	SacrificeBunt PlayType = "sacrificeBunt"
	Steal         PlayType = "steal"
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
	default:
		preset, ok := presetForPlay(state.Runner, play.Type)
		if !ok {
			return &next
		}
		for i := 0; i < preset.outs; i++ {
			if addOut(&next) {
				return &next
			}
		}
		advancement := mergeAdvancement(preset.advancement, play.Override)
		applyAdvancement(&next, advancement)
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

type preset struct {
	advancement Advancement
	outs        int
}

func presetForPlay(runner RunnerState, playType PlayType) (preset, bool) {
	switch playType {
	case Walk, HitByPitch:
		return preset{advancement: walkAdvancement(runner)}, true
	case Single:
		return preset{advancement: singleAdvancement(runner)}, true
	case Double:
		return preset{advancement: doubleAdvancement(runner)}, true
	case Triple:
		return preset{advancement: tripleAdvancement(runner)}, true
	case HomeRun:
		return preset{advancement: homeRunAdvancement(runner)}, true
	case Error:
		return preset{advancement: errorAdvancement(runner)}, true
	case SacrificeBunt:
		return preset{advancement: sacrificeBuntAdvancement(runner), outs: 1}, true
	case Steal:
		return preset{advancement: stealAdvancement(runner)}, true
	default:
		return preset{}, false
	}
}

func walkAdvancement(r RunnerState) Advancement {
	a := Advancement{Batter: destinationPtr(DestinationFirst)}
	if r.First {
		a.FromFirst = destinationPtr(DestinationSecond)
	}
	if r.Second {
		if r.First {
			a.FromSecond = destinationPtr(DestinationThird)
		} else {
			a.FromSecond = destinationPtr(DestinationSecond)
		}
	}
	if r.Third {
		if r.First && r.Second {
			a.FromThird = destinationPtr(DestinationHome)
		} else {
			a.FromThird = destinationPtr(DestinationThird)
		}
	}
	return a
}

func singleAdvancement(r RunnerState) Advancement {
	a := Advancement{Batter: destinationPtr(DestinationFirst)}
	if r.First {
		a.FromFirst = destinationPtr(DestinationSecond)
	}
	if r.Second {
		a.FromSecond = destinationPtr(DestinationThird)
	}
	if r.Third {
		if r.Second {
			a.FromThird = destinationPtr(DestinationThird)
		} else {
			a.FromThird = destinationPtr(DestinationHome)
		}
	}
	if r.First && r.Second {
		a.FromSecond = destinationPtr(DestinationThird)
	}
	if r.First && r.Second && r.Third {
		a.FromThird = destinationPtr(DestinationHome)
	}
	if r.First && r.Third && !r.Second {
		a.FromThird = destinationPtr(DestinationHome)
	}
	if r.Second && r.Third && !r.First {
		a.FromSecond = destinationPtr(DestinationHome)
		a.FromThird = destinationPtr(DestinationThird)
	}
	return a
}

func doubleAdvancement(r RunnerState) Advancement {
	a := Advancement{Batter: destinationPtr(DestinationSecond)}
	if r.First {
		if r.Second || r.Third {
			a.FromFirst = destinationPtr(DestinationThird)
		} else {
			a.FromFirst = destinationPtr(DestinationThird)
		}
	}
	if r.Second {
		a.FromSecond = destinationPtr(DestinationHome)
	}
	if r.Third {
		a.FromThird = destinationPtr(DestinationHome)
	}
	if r.First && r.Second && r.Third {
		a.FromFirst = destinationPtr(DestinationThird)
	}
	if r.First && r.Second && !r.Third {
		a.FromFirst = destinationPtr(DestinationThird)
		a.FromSecond = destinationPtr(DestinationHome)
	}
	if r.First && r.Third && !r.Second {
		a.FromFirst = destinationPtr(DestinationThird)
	}
	if r.Second && r.Third && !r.First {
		a.FromSecond = destinationPtr(DestinationHome)
		a.FromThird = destinationPtr(DestinationHome)
	}
	return a
}

func tripleAdvancement(r RunnerState) Advancement {
	a := Advancement{Batter: destinationPtr(DestinationThird)}
	if r.First {
		a.FromFirst = destinationPtr(DestinationHome)
	}
	if r.Second {
		a.FromSecond = destinationPtr(DestinationHome)
	}
	if r.Third {
		a.FromThird = destinationPtr(DestinationHome)
	}
	return a
}

func homeRunAdvancement(r RunnerState) Advancement {
	a := Advancement{Batter: destinationPtr(DestinationHome)}
	if r.First {
		a.FromFirst = destinationPtr(DestinationHome)
	}
	if r.Second {
		a.FromSecond = destinationPtr(DestinationHome)
	}
	if r.Third {
		a.FromThird = destinationPtr(DestinationHome)
	}
	return a
}

func errorAdvancement(r RunnerState) Advancement {
	a := Advancement{Batter: destinationPtr(DestinationFirst)}
	if r.First {
		a.FromFirst = destinationPtr(DestinationSecond)
	}
	if r.Second {
		if r.First {
			a.FromSecond = destinationPtr(DestinationThird)
		} else {
			a.FromSecond = destinationPtr(DestinationThird)
		}
	}
	if r.Third {
		if r.Second {
			a.FromThird = destinationPtr(DestinationHome)
		} else {
			a.FromThird = destinationPtr(DestinationHome)
		}
	}
	if r.Second && !r.First {
		a.FromSecond = destinationPtr(DestinationThird)
	}
	if r.Second && !r.First && !r.Third {
		a.FromSecond = destinationPtr(DestinationThird)
	}
	if r.Second && !r.First && r.Third {
		a.FromSecond = destinationPtr(DestinationThird)
	}
	return a
}

func sacrificeBuntAdvancement(r RunnerState) Advancement {
	a := Advancement{Batter: destinationPtr(DestinationOut)}
	if r.First {
		a.FromFirst = destinationPtr(DestinationSecond)
	}
	if r.Second {
		if r.First {
			a.FromSecond = destinationPtr(DestinationThird)
		} else {
			a.FromSecond = destinationPtr(DestinationThird)
		}
	}
	if r.Third {
		if r.First || r.Second {
			a.FromThird = destinationPtr(DestinationHome)
		} else {
			a.FromThird = destinationPtr(DestinationHome)
		}
	}
	if !r.First && !r.Second && !r.Third {
		a.Batter = destinationPtr(DestinationOut)
	}
	return a
}

func stealAdvancement(r RunnerState) Advancement {
	a := Advancement{}
	if r.First {
		a.FromFirst = destinationPtr(DestinationSecond)
	}
	if r.Second {
		a.FromSecond = destinationPtr(DestinationThird)
	}
	if r.Third {
		a.FromThird = destinationPtr(DestinationHome)
	}
	if r.Second && !r.First {
		a.FromSecond = destinationPtr(DestinationThird)
	}
	return a
}

func mergeAdvancement(base Advancement, override *Advancement) Advancement {
	if override == nil {
		return base
	}
	if override.Batter != nil {
		base.Batter = override.Batter
	}
	if override.FromFirst != nil {
		base.FromFirst = override.FromFirst
	}
	if override.FromSecond != nil {
		base.FromSecond = override.FromSecond
	}
	if override.FromThird != nil {
		base.FromThird = override.FromThird
	}
	return base
}

func applyAdvancement(s *GameState, a Advancement) {
	nextRunner := RunnerState{}
	runs := 0

	if a.Batter != nil {
		placeDestination(*a.Batter, &nextRunner, &runs)
	}
	if s.Runner.First && a.FromFirst != nil {
		placeDestination(*a.FromFirst, &nextRunner, &runs)
	}
	if s.Runner.Second && a.FromSecond != nil {
		placeDestination(*a.FromSecond, &nextRunner, &runs)
	}
	if s.Runner.Third && a.FromThird != nil {
		placeDestination(*a.FromThird, &nextRunner, &runs)
	}

	s.Runner = nextRunner
	if runs > 0 {
		s.AddRun(runs)
	}
}

func placeDestination(destination BaseDestination, runner *RunnerState, runs *int) {
	switch destination {
	case DestinationFirst:
		runner.First = true
	case DestinationSecond:
		runner.Second = true
	case DestinationThird:
		runner.Third = true
	case DestinationHome:
		*runs++
	case DestinationOut:
		return
	}
}

func destinationPtr(destination BaseDestination) *BaseDestination {
	return &destination
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
