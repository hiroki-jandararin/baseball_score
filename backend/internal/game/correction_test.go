package game

func (s *GameStateSuite) TestRecalculate_ReplaysCorrectedHistoryFromBeginning() {
	originalPlays := []Play{
		{Type: Walk},
		{Type: Walk},
		{Type: Strikeout},
	}

	correctedPlays := []Play{
		originalPlays[0],
		{Type: Strikeout},
		originalPlays[2],
	}

	result := Recalculate(correctedPlays)

	s.Equal(1, result.Inning)
	s.Equal(Top, result.InningHalf)
	s.Equal(2, result.Outs)
	s.True(result.Runner.First)
	s.False(result.Runner.Second)
	s.False(result.Runner.Third)
	s.Equal(0, result.Teams["team1"].Score)
	s.Equal(0, result.Teams["team2"].Score)
}
