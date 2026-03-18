package game

func (s *GameStateSuite) TestApplyPlay_StrikeoutIncreasesOuts() {
	s.Run("0 outs empty bases -> 1 out", func() {
		state := NewGameState()
		play := Play{
			Type: Strikeout,
		}
		next := ApplyPlay(state, play)
		s.Equal(1, next.Outs)
	})
}

func (s *GameStateSuite) TestApplyPlay_StrikeoutOnlyAffectsOuts() {
	s.Run("0 outs runner on first -> 1 out runner stays on first", func() {
		state := GameState {
			Inning:     1,
			InningHalf: Top,
			HomeScore:  0,
			AwayScore:  0,
			Runner: RunnerState{
				First:  true,
				Second: false,
				Third:  false,
			},
			Outs: 0,
		}
		play := Play{
			Type: Strikeout,
		}
		next := ApplyPlay(&state, play)
		s.Equal(1, next.Outs)
		s.Equal(1, next.Inning)
		s.Equal(Top, next.InningHalf)
		s.Equal(0, next.HomeScore)
		s.Equal(0, next.AwayScore)
		s.Equal(true, next.Runner.First)
		s.Equal(false, next.Runner.Second)
		s.Equal(false, next.Runner.Third)
	})
}

func (s *GameStateSuite) TestApplyPlay_ThreeOutsChangeHalfInning() {
	s.Run("2 outs top of 1st -> bottom of 1st and bases cleared", func() {
		state := GameState {
			Inning:     1,
			InningHalf: Top,
			HomeScore:  0,
			AwayScore:  0,
			Runner: RunnerState{
				First:  true,
				Second: false,
				Third:  false,
			},
			Outs: 2,
		}
		play := Play{
			Type: Strikeout,
		}
		next := ApplyPlay(&state, play)
		s.Equal(0, next.Outs)
		s.Equal(1, next.Inning)
		s.Equal(Bottom, next.InningHalf)
		s.Equal(0, next.HomeScore)
		s.Equal(0, next.AwayScore)
		s.Equal(false, next.Runner.First)
		s.Equal(false, next.Runner.Second)
		s.Equal(false, next.Runner.Third)
	})
}

func (s *GameStateSuite) TestApplyPlay_ThreeOutsChangeInning() {
	s.Run("2 outs bottom of 1st -> top of 2nd and bases cleared", func() {
		state := GameState {
			Inning:     1,
			InningHalf: Bottom,
			HomeScore:  0,
			AwayScore:  0,
			Runner: RunnerState{
				First:  true,
				Second: true,
				Third:  true,
			},
			Outs: 2,
		}
		play := Play{
			Type: Strikeout,
		}
		next := ApplyPlay(&state, play)
		s.Equal(0, next.Outs)
		s.Equal(2, next.Inning)
		s.Equal(Top, next.InningHalf)
		s.Equal(0, next.HomeScore)
		s.Equal(0, next.AwayScore)
		s.Equal(false, next.Runner.First)
		s.Equal(false, next.Runner.Second)
		s.Equal(false, next.Runner.Third)
	})
}