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

func (s *GameStateSuite) TestRecalculate_ReplaysHistoryAfterDeletingPlay() {
	originalPlays := []Play{
		{Type: Walk},
		{Type: Walk},
		{Type: Strikeout},
	}

	correctedPlays := []Play{
		originalPlays[0],
		originalPlays[2],
	}

	result := Recalculate(correctedPlays)

	s.Equal(1, result.Inning)
	s.Equal(Top, result.InningHalf)
	s.Equal(1, result.Outs)
	s.True(result.Runner.First)
	s.False(result.Runner.Second)
	s.False(result.Runner.Third)
	s.Equal(0, result.Teams["team1"].Score)
	s.Equal(0, result.Teams["team2"].Score)
}

func (s *GameStateSuite) TestRecalculate_DeletingHomeRunReducesScore() {
	originalPlays := []Play{
		{Type: Walk},
		{Type: HomeRun},
		{Type: Strikeout},
	}

	correctedPlays := []Play{
		originalPlays[0],
		originalPlays[2],
	}

	result := Recalculate(correctedPlays)

	s.Equal(1, result.Inning)
	s.Equal(Top, result.InningHalf)
	s.Equal(1, result.Outs)
	s.True(result.Runner.First)
	s.False(result.Runner.Second)
	s.False(result.Runner.Third)
	s.Equal(0, result.Teams["team1"].Score)
	s.Equal(0, result.Teams["team2"].Score)
}

func (s *GameStateSuite) TestRecalculate_Integration_WalkSingleHomeRun() {
	plays := []Play{
		{Type: Walk},
		{Type: Single},
		{Type: HomeRun},
	}

	result := Recalculate(plays)

	s.Equal(1, result.Inning)
	s.Equal(Top, result.InningHalf)
	s.Equal(0, result.Outs)
	s.False(result.Runner.First)
	s.False(result.Runner.Second)
	s.False(result.Runner.Third)
	s.Equal(3, result.Teams["team1"].Score)
	s.Equal(0, result.Teams["team2"].Score)
}

func (s *GameStateSuite) TestRecalculate_Integration_ThreeOutsChangeSides() {
	plays := []Play{
		{Type: Strikeout},
		{Type: Groundout},
		{Type: Flyout},
	}

	result := Recalculate(plays)

	s.Equal(1, result.Inning)
	s.Equal(Bottom, result.InningHalf)
	s.Equal(0, result.Outs)
	s.False(result.Runner.First)
	s.False(result.Runner.Second)
	s.False(result.Runner.Third)
	s.Equal(0, result.Teams["team1"].Score)
	s.Equal(0, result.Teams["team2"].Score)
}

func (s *GameStateSuite) TestRecalculate_Integration_WalkSacrificeBuntSingleScores() {
	plays := []Play{
		{Type: Walk},
		{Type: SacrificeBunt},
		{
			Type: Single,
			Override: &Advancement{
				FromSecond: destinationPtr(DestinationHome),
			},
		},
	}

	result := Recalculate(plays)

	s.Equal(1, result.Inning)
	s.Equal(Top, result.InningHalf)
	s.Equal(1, result.Outs)
	s.True(result.Runner.First)
	s.False(result.Runner.Second)
	s.False(result.Runner.Third)
	s.Equal(1, result.Teams["team1"].Score)
	s.Equal(0, result.Teams["team2"].Score)
}
