package game

func (s *GameStateSuite) TestApplyPlay_SingleOverrideAllowsRunnerAndBatterToAdvanceDifferently() {
	state := GameState{
		Inning:              1,
		InningHalf:          Top,
		FirstBattingTeamID:  TeamID("team1"),
		SecondBattingTeamID: TeamID("team2"),
		Teams: map[TeamID]*TeamState{
			"team1": {TeamID: "team1", Score: 0},
			"team2": {TeamID: "team2", Score: 0},
		},
		Runner: RunnerState{
			First:  false,
			Second: true,
			Third:  false,
		},
		Outs: 0,
	}

	play := Play{
		Type: Single,
		Override: &Advancement{
			FromSecond: destinationPtr(DestinationHome),
			Batter:     destinationPtr(DestinationFirst),
		},
	}

	next := RecordPlay(&state, play)

	s.Equal(1, next.Teams["team1"].Score)
	s.True(next.Runner.First)
	s.False(next.Runner.Second)
	s.False(next.Runner.Third)
}

func (s *GameStateSuite) TestApplyPlay_SingleOverrideCanKeepSecondRunnerOnThird() {
	state := GameState{
		Inning:              1,
		InningHalf:          Top,
		FirstBattingTeamID:  TeamID("team1"),
		SecondBattingTeamID: TeamID("team2"),
		Teams: map[TeamID]*TeamState{
			"team1": {TeamID: "team1", Score: 0},
			"team2": {TeamID: "team2", Score: 0},
		},
		Runner: RunnerState{
			First:  false,
			Second: true,
			Third:  false,
		},
		Outs: 0,
	}

	play := Play{
		Type: Single,
		Override: &Advancement{
			FromSecond: destinationPtr(DestinationThird),
		},
	}

	next := RecordPlay(&state, play)

	s.Equal(0, next.Teams["team1"].Score)
	s.True(next.Runner.First)
	s.False(next.Runner.Second)
	s.True(next.Runner.Third)
}
