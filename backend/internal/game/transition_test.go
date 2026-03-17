package game


func (s *GameStateSuite) TestApplyPlay_StrikeoutIncreasesOuts() {
	s.Run("strikeout", func() {
		state := NewGameState()
		play := Play{
			Type: "strikeout",
		}
		next := ApplyPlay(state, play)
		s.Equal(1, next.Outs)
	})
}

func (s *GameStateSuite) TestApplyPlay_StrikeoutOnlyAffectsOuts() {
	s.Run("strikeout", func() {
		state := GameState {
			Inning:     1,
			InningHalf: "top",
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
			Type: "strikeout",
		}
		next := ApplyPlay(&state, play)
		s.Equal(1, next.Outs)
		s.Equal(1, next.Inning)
		s.Equal("top", next.InningHalf)
		s.Equal(0, next.HomeScore)
		s.Equal(0, next.AwayScore)
		s.Equal(true, next.Runner.First)
		s.Equal(false, next.Runner.Second)
		s.Equal(false, next.Runner.Third)
	})
}