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
			FirstBattingTeamID:  TeamID("team1"),
			SecondBattingTeamID: TeamID("team2"),
			Teams: map[TeamID]*TeamState{
				"team1": {TeamID: "team1", Score: 0},
				"team2": {TeamID: "team2", Score: 0},
			},
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
		s.Equal(0, state.Teams["team1"].Score)
		s.Equal(0, state.Teams["team2"].Score)
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
			FirstBattingTeamID:  TeamID("team1"),
			SecondBattingTeamID: TeamID("team2"),
			Teams: map[TeamID]*TeamState{
				"team1": {TeamID: "team1", Score: 0},
				"team2": {TeamID: "team2", Score: 0},
			},
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
		s.Equal(0, next.Teams["team1"].Score)
		s.Equal(0, next.Teams["team2"].Score)
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
			FirstBattingTeamID:  TeamID("team1"),
			SecondBattingTeamID: TeamID("team2"),
			Teams: map[TeamID]*TeamState{
				"team1": {TeamID: "team1", Score: 0},
				"team2": {TeamID: "team2", Score: 0},
			},
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
		s.Equal(0, next.Teams["team1"].Score)
		s.Equal(0, next.Teams["team2"].Score)
		s.Equal(false, next.Runner.First)
		s.Equal(false, next.Runner.Second)
		s.Equal(false, next.Runner.Third)
	})
}

func (s *GameStateSuite) TestApplyPlay_flyOutsChangeInningAndScore() {
	s.Run("2 outs bottom of 1st with runner on third -> top of 2nd ", func() {
		state := GameState {
			Inning:     1,
			InningHalf: Bottom,
			FirstBattingTeamID:  TeamID("team1"),
			SecondBattingTeamID: TeamID("team2"),
			Teams: map[TeamID]*TeamState{
				"team1": {TeamID: "team1", Score: 0},
				"team2": {TeamID: "team2", Score: 0},
			},
			Runner: RunnerState{
				First:  false,
				Second: false,
				Third:  true,
			},
			Outs: 2,
		}
		play := Play{
			Type: Flyout,
		}
		next := ApplyPlay(&state, play)
		s.Equal(0, next.Outs)
		s.Equal(2, next.Inning)
		s.Equal(Top, next.InningHalf)
		s.Equal(0, next.Teams["team1"].Score)
		s.Equal(0, next.Teams["team2"].Score)
		s.Equal(false, next.Runner.First)
		s.Equal(false, next.Runner.Second)
		s.Equal(false, next.Runner.Third)
	})
}

func (s *GameStateSuite) TestApplyPlay_GroundOutsChangeInning() {
	s.Run("2 outs top of 2nd with runner on third -> bottom of 2nd ", func() {
		state := GameState {
			Inning:     2,
			InningHalf: Top,
			FirstBattingTeamID:  TeamID("team1"),
			SecondBattingTeamID: TeamID("team2"),
			Teams: map[TeamID]*TeamState{
				"team1": {TeamID: "team1", Score: 0},
				"team2": {TeamID: "team2", Score: 0},
			},
			Runner: RunnerState{
				First:  false,
				Second: false,
				Third:  true,
			},
			Outs: 2,
		}
		play := Play{
			Type: Groundout,
		}
		next := ApplyPlay(&state, play)
		s.Equal(0, next.Outs)
		s.Equal(2, next.Inning)
		s.Equal(Bottom, next.InningHalf)
		s.Equal(0, next.Teams["team1"].Score)
		s.Equal(0, next.Teams["team2"].Score)
		s.Equal(false, next.Runner.First)
		s.Equal(false, next.Runner.Second)
		s.Equal(false, next.Runner.Third)
	})
}

func (s *GameStateSuite) TestApplyPlay_WalkPutsRunnersOnFirst() {
	s.Run("0 outs empty bases -> 0 outs runner on first", func() {
		state := NewGameState()
		play := Play{
			Type: Walk,
		}
		next := ApplyPlay(state, play)
		s.Equal(0, next.Outs)
		s.Equal(1, next.Inning)
		s.Equal(Top, next.InningHalf)
		s.Equal(0, next.Teams["team1"].Score)
		s.Equal(0, next.Teams["team2"].Score)
		s.Equal(true, next.Runner.First)
		s.Equal(false, next.Runner.Second)
		s.Equal(false, next.Runner.Third)
	})
}

func (s *GameStateSuite) TestApplyPlay_WalkForcesRunnerToSecond() {
	s.Run("0 outs runner on first -> 0 outs runners on first and second", func() {
		state := GameState {
			Inning:     1,
			InningHalf: Top,
			FirstBattingTeamID:  TeamID("team1"),
			SecondBattingTeamID: TeamID("team2"),
			Teams: map[TeamID]*TeamState{
				"team1": {TeamID: "team1", Score: 0},
				"team2": {TeamID: "team2", Score: 0},
			},
			Runner: RunnerState{
				First:  true,
				Second: false,
				Third:  false,
			},
			Outs: 0,
		}
		play := Play{
			Type: Walk,
		}
		next := ApplyPlay(&state, play)
		s.Equal(0, next.Outs)
		s.Equal(1, next.Inning)
		s.Equal(Top, next.InningHalf)
		s.Equal(0, next.Teams["team1"].Score)
		s.Equal(0, next.Teams["team2"].Score)
		s.Equal(true, next.Runner.First)
		s.Equal(true, next.Runner.Second)
		s.Equal(false, next.Runner.Third)
	})	
}

func (s *GameStateSuite) TestApplyPlay_WalkLoadsBases() {
	s.Run("0 outs runners on first and second -> 0 outs bases loaded", func() {
		state := GameState {
			Inning:     1,
			InningHalf: Top,
			FirstBattingTeamID:  TeamID("team1"),
			SecondBattingTeamID: TeamID("team2"),
			Teams: map[TeamID]*TeamState{
				"team1": {TeamID: "team1", Score: 0},
				"team2": {TeamID: "team2", Score: 0},
			},
			Runner: RunnerState{
				First:  true,
				Second: true,
				Third:  false,
			},
			Outs: 0,
		}
		play := Play{
			Type: Walk,
		}
		next := ApplyPlay(&state, play)
		s.Equal(0, next.Outs)
		s.Equal(1, next.Inning)
		s.Equal(Top, next.InningHalf)
		s.Equal(0, next.Teams["team1"].Score)
		s.Equal(0, next.Teams["team2"].Score)
		s.Equal(true, next.Runner.First)
		s.Equal(true, next.Runner.Second)
		s.Equal(true, next.Runner.Third)
	})	
}

func (s *GameStateSuite) TestApplyPlay_WalkWithBasesLoadedScoresOne() {
	s.Run("0 outs bases loaded -> 0 outs bases loaded and 1 run scored", func() {
		state := GameState {
			Inning:     1,
			InningHalf: Top,
			FirstBattingTeamID:  TeamID("team1"),
			SecondBattingTeamID: TeamID("team2"),
			Teams: map[TeamID]*TeamState{
				"team1": {TeamID: "team1", Score: 0},
				"team2": {TeamID: "team2", Score: 0},
			},
			Runner: RunnerState{
				First:  true,
				Second: true,
				Third:  true,
			},
			Outs: 0,
		}
		play := Play{
			Type: Walk,
		}
		next := ApplyPlay(&state, play)
		s.Equal(0, next.Outs)
		s.Equal(1, next.Inning)
		s.Equal(Top, next.InningHalf)
		s.Equal(1, next.Teams["team1"].Score)
		s.Equal(0, next.Teams["team2"].Score)
		s.Equal(true, next.Runner.First)
		s.Equal(true, next.Runner.Second)
		s.Equal(true, next.Runner.Third)
	})
}
