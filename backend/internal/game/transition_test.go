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