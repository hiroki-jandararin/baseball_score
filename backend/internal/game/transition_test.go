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

func (s *GameStateSuite) TestApplyPlay_WalkRunnersOnFirstAndSecond() {
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

func (s *GameStateSuite) TestApplyPlay_WalkRunnersOnFirstAndThird() {
	s.Run("0 outs runners on first and third -> 0 outs bases loaded and 1 run scored", func() {
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
		s.Equal(0, next.Teams["team1"].Score)
		s.Equal(0, next.Teams["team2"].Score)
		s.Equal(true, next.Runner.First)
		s.Equal(true, next.Runner.Second)
		s.Equal(true, next.Runner.Third)
	})
}

func (s *GameStateSuite) TestApplyPlay_WalkRunnersOnSecondAndThird() {
	s.Run("0 outs runners on second and third -> 0 outs bases loaded and 1 run scored", func() {
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
				First:  false,
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
		s.Equal(0, next.Teams["team1"].Score)
		s.Equal(0, next.Teams["team2"].Score)
		s.Equal(true, next.Runner.First)
		s.Equal(true, next.Runner.Second)
		s.Equal(true, next.Runner.Third)
	})
}

func (s *GameStateSuite) TestApplyPlay_WalkRunnerOnFirst() {
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

func (s *GameStateSuite) TestApplyPlay_WalkRunnerOnSecond() {
	s.Run("0 outs runner on second -> 0 outs runners on first and second", func() {
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
				First:  false,
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
		s.Equal(false, next.Runner.Third)
	})
}

func (s *GameStateSuite) TestApplyPlay_WalkRunnerOnThird() {
	s.Run("0 outs runner on third -> 0 outs runners on first and third", func() {
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
				First:  false,
				Second: false,
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
		s.Equal(0, next.Teams["team1"].Score)
		s.Equal(0, next.Teams["team2"].Score)
		s.Equal(true, next.Runner.First)
		s.Equal(false, next.Runner.Second)
		s.Equal(true, next.Runner.Third)
	})
}

func (s *GameStateSuite) TestApplyPlay_WalkRunnersEmpty() {
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

func (s *GameStateSuite) TestApplyPlay_HitByPitchPutsRunnersOnFirst() {
	s.Run("0 outs empty bases -> 0 outs runner on first", func() {
		state := NewGameState()
		play := Play{
			Type: HitByPitch,
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

func (s *GameStateSuite) TestApplyPlay_SingleRunnerOccupied() {
	s.Run("0 outs runners on first, second and third -> 0 outs bases loaded and 1 run scored", func() {
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
			Type: Single,
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

func (s *GameStateSuite) TestApplyPlay_SingleRunnerOnFirstAndSecond() {
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
			Type: Single,
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

func (s *GameStateSuite) TestApplyPlay_SingleRunnerOnFirstAndThird() {
	s.Run("0 outs runners on first and third -> 0 outs runners on first, second and third with a run scored", func() {
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
				Third:  true,
			},
			Outs: 0,
		}
		play := Play{
			Type: Single,
		}
		next := ApplyPlay(&state, play)
		s.Equal(0, next.Outs)
		s.Equal(1, next.Inning)
		s.Equal(Top, next.InningHalf)
		s.Equal(1, next.Teams["team1"].Score)
		s.Equal(0, next.Teams["team2"].Score)
		s.Equal(true, next.Runner.First)
		s.Equal(true, next.Runner.Second)
		s.Equal(false, next.Runner.Third)
	})	
}

func (s *GameStateSuite) TestApplyPlay_SingleRunnerOnSecondAndThird() {
	s.Run("0 outs runners on second and third -> 0 outs runner on first and third with a run scored", func() {
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
				First:  false,
				Second: true,
				Third:  true,
			},
			Outs: 0,
		}
		play := Play{
			Type: Single,
		}
		next := ApplyPlay(&state, play)
		s.Equal(0, next.Outs)
		s.Equal(1, next.Inning)
		s.Equal(Top, next.InningHalf)
		s.Equal(1, next.Teams["team1"].Score)
		s.Equal(0, next.Teams["team2"].Score)
		s.Equal(true, next.Runner.First)
		s.Equal(false, next.Runner.Second)
		s.Equal(true, next.Runner.Third)
	})	
}

func (s *GameStateSuite) TestApplyPlay_SingleRunnersOnFirst() {
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
			Type: Single,
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

func (s *GameStateSuite) TestApplyPlay_SingleRunnersOnSecond() {
	s.Run("0 outs runner on second -> 0 outs runners on first and third", func() {
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
				First:  false,
				Second: true,
				Third:  false,
			},
			Outs: 0,
		}
		play := Play{
			Type: Single,
		}
		next := ApplyPlay(&state, play)
		s.Equal(0, next.Outs)
		s.Equal(1, next.Inning)
		s.Equal(Top, next.InningHalf)
		s.Equal(0, next.Teams["team1"].Score)
		s.Equal(0, next.Teams["team2"].Score)
		s.Equal(true, next.Runner.First)
		s.Equal(false, next.Runner.Second)
		s.Equal(true, next.Runner.Third)
	})	
}

func (s *GameStateSuite) TestApplyPlay_SingleRunnersOnThird() {
	s.Run("0 outs runner on third -> 0 outs runners on first with a run scored", func() {
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
				First:  false,
				Second: false,
				Third:  true,
			},
			Outs: 0,
		}
		play := Play{
			Type: Single,
		}
		next := ApplyPlay(&state, play)
		s.Equal(0, next.Outs)
		s.Equal(1, next.Inning)
		s.Equal(Top, next.InningHalf)
		s.Equal(1, next.Teams["team1"].Score)
		s.Equal(0, next.Teams["team2"].Score)
		s.Equal(true, next.Runner.First)
		s.Equal(false, next.Runner.Second)
		s.Equal(false, next.Runner.Third)
	})
}

func (s *GameStateSuite) TestApplyPlay_SingleRunnersEmpty() {
	s.Run("0 outs empty bases -> 0 outs runner on first", func() {
		state := NewGameState()
		play := Play{
			Type: Single,
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

func (s *GameStateSuite) TestApplyPlay_DoubleRunnersOnFirstAndSecondAndThird() {
	s.Run("0 outs runners on first, second and third -> 0 outs runner on second and third with two runs scored", func() {
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
			Type: Double,
		}
		next := ApplyPlay(&state, play)
		s.Equal(0, next.Outs)
		s.Equal(1, next.Inning)
		s.Equal(Top, next.InningHalf)
		s.Equal(2, next.Teams["team1"].Score)
		s.Equal(0, next.Teams["team2"].Score)
		s.Equal(false, next.Runner.First)
		s.Equal(true, next.Runner.Second)
		s.Equal(true, next.Runner.Third)
	})
}

func (s *GameStateSuite) TestApplyPlay_DoubleRunnersOnFirstAndSecond() {
	s.Run("0 outs runners on first and second -> 0 outs runner on second and third with a run scored", func() {
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
			Type: Double,
		}
		next := ApplyPlay(&state, play)
		s.Equal(0, next.Outs)
		s.Equal(1, next.Inning)
		s.Equal(Top, next.InningHalf)
		s.Equal(1, next.Teams["team1"].Score)
		s.Equal(0, next.Teams["team2"].Score)
		s.Equal(false, next.Runner.First)
		s.Equal(true, next.Runner.Second)
		s.Equal(true, next.Runner.Third)
	})
}

func (s *GameStateSuite) TestApplyPlay_DoubleRunnersOnFirstAndThird() {
	s.Run("0 outs runners on first and third -> 0 outs runner on  second and third with a run scored", func() {
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
				Third:  true,
			},
			Outs: 0,
		}
		play := Play{
			Type: Double,
		}
		next := ApplyPlay(&state, play)
		s.Equal(0, next.Outs)
		s.Equal(1, next.Inning)
		s.Equal(Top, next.InningHalf)
		s.Equal(1, next.Teams["team1"].Score)
		s.Equal(0, next.Teams["team2"].Score)
		s.Equal(false, next.Runner.First)
		s.Equal(true, next.Runner.Second)
		s.Equal(true, next.Runner.Third)
	})
}

func (s *GameStateSuite) TestApplyPlay_DoubleRunnersOnSecondAndThird() {
	s.Run("0 outs runners on second and third -> 0 outs runner on second with two runs scored", func() {
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
				First:  false,
				Second: true,
				Third:  true,
			},
			Outs: 0,
		}
		play := Play{
			Type: Double,
		}
		next := ApplyPlay(&state, play)
		s.Equal(0, next.Outs)
		s.Equal(1, next.Inning)
		s.Equal(Top, next.InningHalf)
		s.Equal(2, next.Teams["team1"].Score)
		s.Equal(0, next.Teams["team2"].Score)
		s.Equal(false, next.Runner.First)
		s.Equal(true, next.Runner.Second)
		s.Equal(false, next.Runner.Third)
	})
}

func (s *GameStateSuite) TestApplyPlay_DoubleRunnersOnFirst() {
	s.Run("0 outs runner on first -> 0 outs runner on second and third", func() {
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
			Type: Double,
		}
		next := ApplyPlay(&state, play)
		s.Equal(0, next.Outs)
		s.Equal(1, next.Inning)
		s.Equal(Top, next.InningHalf)
		s.Equal(0, next.Teams["team1"].Score)
		s.Equal(0, next.Teams["team2"].Score)
		s.Equal(false, next.Runner.First)
		s.Equal(true, next.Runner.Second)
		s.Equal(true, next.Runner.Third)
	})
}

func (s *GameStateSuite) TestApplyPlay_DoubleRunnersOnSecond() {
	s.Run("0 outs runner on second -> 0 outs runner on second with a run scored", func() {
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
				First:  false,
				Second: true,
				Third:  false,
			},
			Outs: 0,
		}
		play := Play{
			Type: Double,
		}
		next := ApplyPlay(&state, play)
		s.Equal(0, next.Outs)
		s.Equal(1, next.Inning)
		s.Equal(Top, next.InningHalf)
		s.Equal(1, next.Teams["team1"].Score)
		s.Equal(0, next.Teams["team2"].Score)
		s.Equal(false, next.Runner.First)
		s.Equal(true, next.Runner.Second)
		s.Equal(false, next.Runner.Third)
	})
}

func (s *GameStateSuite) TestApplyPlay_DoubleRunnersOnThird() {
	s.Run("0 outs runner on third -> 0 outs runner on second with a run scored", func() {
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
				First:  false,
				Second: false,
				Third:  true,
			},
			Outs: 0,
		}
		play := Play{
			Type: Double,
		}
		next := ApplyPlay(&state, play)
		s.Equal(0, next.Outs)
		s.Equal(1, next.Inning)
		s.Equal(Top, next.InningHalf)
		s.Equal(1, next.Teams["team1"].Score)
		s.Equal(0, next.Teams["team2"].Score)
		s.Equal(false, next.Runner.First)
		s.Equal(true, next.Runner.Second)
		s.Equal(false, next.Runner.Third)
	})
}

func (s *GameStateSuite) TestApplyPlay_DoubleRunnersEmpty() {
	s.Run("0 outs empty bases -> 0 outs runner on second", func() {
		state := NewGameState()
		play := Play{
			Type: Double,
		}
		next := ApplyPlay(state, play)
		s.Equal(0, next.Outs)
		s.Equal(1, next.Inning)
		s.Equal(Top, next.InningHalf)
		s.Equal(0, next.Teams["team1"].Score)
		s.Equal(0, next.Teams["team2"].Score)
		s.Equal(false, next.Runner.First)
		s.Equal(true, next.Runner.Second)
		s.Equal(false, next.Runner.Third)
	})
}

func (s *GameStateSuite) TestApplyPlay_TripleRunnersOnFirstAndSecondAndThird() {
	s.Run("0 outs runners on first, second and third -> 0 outs runner on third with three runs scored", func() {
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
			Type: Triple,
		}
		next := ApplyPlay(&state, play)
		s.Equal(0, next.Outs)
		s.Equal(1, next.Inning)
		s.Equal(Top, next.InningHalf)
		s.Equal(3, next.Teams["team1"].Score)
		s.Equal(0, next.Teams["team2"].Score)
		s.Equal(false, next.Runner.First)
		s.Equal(false, next.Runner.Second)
		s.Equal(true, next.Runner.Third)
	})
}

func (s *GameStateSuite) TestApplyPlay_TripleRunnersOnFirstAndSecond() {
	s.Run("0 outs runners on first and second -> 0 outs runner on third with two runs scored", func() {
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
			Type: Triple,
		}
		next := ApplyPlay(&state, play)
		s.Equal(0, next.Outs)
		s.Equal(1, next.Inning)
		s.Equal(Top, next.InningHalf)
		s.Equal(2, next.Teams["team1"].Score)
		s.Equal(0, next.Teams["team2"].Score)
		s.Equal(false, next.Runner.First)
		s.Equal(false, next.Runner.Second)
		s.Equal(true, next.Runner.Third)
	})
}

func (s *GameStateSuite) TestApplyPlay_TripleRunnersOnFirstAndThird() {
	s.Run("0 outs runners on first and third -> 0 outs runner on third with two runs scored", func() {
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
				Third:  true,
			},
			Outs: 0,
		}
		play := Play{
			Type: Triple,
		}
		next := ApplyPlay(&state, play)
		s.Equal(0, next.Outs)
		s.Equal(1, next.Inning)
		s.Equal(Top, next.InningHalf)
		s.Equal(2, next.Teams["team1"].Score)
		s.Equal(0, next.Teams["team2"].Score)
		s.Equal(false, next.Runner.First)
		s.Equal(false, next.Runner.Second)
		s.Equal(true, next.Runner.Third)
	})
}

func (s *GameStateSuite) TestApplyPlay_TripleRunnersOnSecondAndThird() {
	s.Run("0 outs runners on second and third -> 0 outs runner on third with two runs scored", func() {
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
				First:  false,
				Second: true,
				Third:  true,
			},
			Outs: 0,
		}
		play := Play{
			Type: Triple,
		}
		next := ApplyPlay(&state, play)
		s.Equal(0, next.Outs)
		s.Equal(1, next.Inning)
		s.Equal(Top, next.InningHalf)
		s.Equal(2, next.Teams["team1"].Score)
		s.Equal(0, next.Teams["team2"].Score)
		s.Equal(false, next.Runner.First)
		s.Equal(false, next.Runner.Second)
		s.Equal(true, next.Runner.Third)
	})
}

func (s *GameStateSuite) TestApplyPlay_TripleRunnersOnFirst() {
	s.Run("0 outs runner on first -> 0 outs runner third with a run scored", func() {
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
			Type: Triple,
		}
		next := ApplyPlay(&state, play)
		s.Equal(0, next.Outs)
		s.Equal(1, next.Inning)
		s.Equal(Top, next.InningHalf)
		s.Equal(1, next.Teams["team1"].Score)
		s.Equal(0, next.Teams["team2"].Score)
		s.Equal(false, next.Runner.First)
		s.Equal(false, next.Runner.Second)
		s.Equal(true, next.Runner.Third)
	})
}

func (s *GameStateSuite) TestApplyPlay_TripleRunnersOnSecond() {
	s.Run("0 outs runner on second -> 0 outs runner on third with a run scored", func() {
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
				First:  false,
				Second: true,
				Third:  false,
			},
			Outs: 0,
		}
		play := Play{
			Type: Triple,
		}
		next := ApplyPlay(&state, play)
		s.Equal(0, next.Outs)
		s.Equal(1, next.Inning)
		s.Equal(Top, next.InningHalf)
		s.Equal(1, next.Teams["team1"].Score)
		s.Equal(0, next.Teams["team2"].Score)
		s.Equal(false, next.Runner.First)
		s.Equal(false, next.Runner.Second)
		s.Equal(true, next.Runner.Third)
	})
}

func (s *GameStateSuite) TestApplyPlay_TripleRunnersOnThird() {
	s.Run("0 outs runner on third -> 0 outs runner on third with a run scored", func() {
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
				First:  false,
				Second: false,
				Third:  true,
			},
			Outs: 0,
		}
		play := Play{
			Type: Triple,
		}
		next := ApplyPlay(&state, play)
		s.Equal(0, next.Outs)
		s.Equal(1, next.Inning)
		s.Equal(Top, next.InningHalf)
		s.Equal(1, next.Teams["team1"].Score)
		s.Equal(0, next.Teams["team2"].Score)
		s.Equal(false, next.Runner.First)
		s.Equal(false, next.Runner.Second)
		s.Equal(true, next.Runner.Third)
	})
}

func (s *GameStateSuite) TestApplyPlay_TripleRunnersEmpty() {
	s.Run("0 outs empty bases -> 0 outs runner on third", func() {
		state := NewGameState()
		play := Play{
			Type: Triple,
		}
		next := ApplyPlay(state, play)
		s.Equal(0, next.Outs)
		s.Equal(1, next.Inning)
		s.Equal(Top, next.InningHalf)
		s.Equal(0, next.Teams["team1"].Score)
		s.Equal(0, next.Teams["team2"].Score)
		s.Equal(false, next.Runner.First)
		s.Equal(false, next.Runner.Second)
		s.Equal(true, next.Runner.Third)
	})
}

func (s *GameStateSuite) TestApplyPlay_HomerunRunnersOnFirstAndSecondAndThird() {
	s.Run("0 outs runners on first, second and third -> 0 outs no runners with four runs scored", func() {
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
			Type: HomeRun,
		}
		next := ApplyPlay(&state, play)
		s.Equal(0, next.Outs)
		s.Equal(1, next.Inning)
		s.Equal(Top, next.InningHalf)
		s.Equal(4, next.Teams["team1"].Score)
		s.Equal(0, next.Teams["team2"].Score)
		s.Equal(false, next.Runner.First)
		s.Equal(false, next.Runner.Second)
		s.Equal(false, next.Runner.Third)
	})
}

func (s *GameStateSuite) TestApplyPlay_HomerunRunnersOnFirstAndSecond() {
	s.Run("0 outs runners on first and second -> 0 outs no runners with three runs scored", func() {
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
			Type: HomeRun,
		}
		next := ApplyPlay(&state, play)
		s.Equal(0, next.Outs)
		s.Equal(1, next.Inning)
		s.Equal(Top, next.InningHalf)
		s.Equal(3, next.Teams["team1"].Score)
		s.Equal(0, next.Teams["team2"].Score)
		s.Equal(false, next.Runner.First)
		s.Equal(false, next.Runner.Second)
		s.Equal(false, next.Runner.Third)
	})
}

func (s *GameStateSuite) TestApplyPlay_HomerunRunnersOnFirstAndThird() {
	s.Run("0 outs runners on first and third -> 0 outs no runners with three runs scored", func() {
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
				Third:  true,
			},
			Outs: 0,
		}
		play := Play{
			Type: HomeRun,
		}
		next := ApplyPlay(&state, play)
		s.Equal(0, next.Outs)
		s.Equal(1, next.Inning)
		s.Equal(Top, next.InningHalf)
		s.Equal(3, next.Teams["team1"].Score)
		s.Equal(0, next.Teams["team2"].Score)
		s.Equal(false, next.Runner.First)
		s.Equal(false, next.Runner.Second)
		s.Equal(false, next.Runner.Third)
	})
}

func (s *GameStateSuite) TestApplyPlay_HomerunRunnersOnSecondAndThird() {
	s.Run("0 outs runners on second and third -> 0 outs no runners with three runs scored", func() {
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
				First:  false,
				Second: true,
				Third:  true,
			},
			Outs: 0,
		}
		play := Play{
			Type: HomeRun,
		}
		next := ApplyPlay(&state, play)
		s.Equal(0, next.Outs)
		s.Equal(1, next.Inning)
		s.Equal(Top, next.InningHalf)
		s.Equal(3, next.Teams["team1"].Score)
		s.Equal(0, next.Teams["team2"].Score)
		s.Equal(false, next.Runner.First)
		s.Equal(false, next.Runner.Second)
		s.Equal(false, next.Runner.Third)
	})
}

func (s *GameStateSuite) TestApplyPlay_HomerunRunnersOnFirst() {
	s.Run("0 outs runner on first -> 0 outs no runners with two runs scored", func() {
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
			Type: HomeRun,
		}
		next := ApplyPlay(&state, play)
		s.Equal(0, next.Outs)
		s.Equal(1, next.Inning)
		s.Equal(Top, next.InningHalf)
		s.Equal(2, next.Teams["team1"].Score)
		s.Equal(0, next.Teams["team2"].Score)
		s.Equal(false, next.Runner.First)
		s.Equal(false, next.Runner.Second)
		s.Equal(false, next.Runner.Third)
	})
}

func (s *GameStateSuite) TestApplyPlay_HomerunRunnersOnSecond() {
	s.Run("0 outs runner on second -> 0 outs no runners with two runs scored", func() {
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
				First:  false,
				Second: true,
				Third:  false,
			},
			Outs: 0,
		}
		play := Play{
			Type: HomeRun,
		}
		next := ApplyPlay(&state, play)
		s.Equal(0, next.Outs)
		s.Equal(1, next.Inning)
		s.Equal(Top, next.InningHalf)
		s.Equal(2, next.Teams["team1"].Score)
		s.Equal(0, next.Teams["team2"].Score)
		s.Equal(false, next.Runner.First)
		s.Equal(false, next.Runner.Second)
		s.Equal(false, next.Runner.Third)
	})
}

func (s *GameStateSuite) TestApplyPlay_HomerunRunnersOnThird() {
	s.Run("0 outs runner on third -> 0 outs no runners with two runs scored", func() {
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
				First:  false,
				Second: false,
				Third:  true,
			},
			Outs: 0,
		}
		play := Play{
			Type: HomeRun,
		}
		next := ApplyPlay(&state, play)
		s.Equal(0, next.Outs)
		s.Equal(1, next.Inning)
		s.Equal(Top, next.InningHalf)
		s.Equal(2, next.Teams["team1"].Score)
		s.Equal(0, next.Teams["team2"].Score)
		s.Equal(false, next.Runner.First)
		s.Equal(false, next.Runner.Second)
		s.Equal(false, next.Runner.Third)
	})
}

func (s *GameStateSuite) TestApplyPlay_HomerunRunnersEmpty() {
	s.Run("0 outs empty bases -> 0 outs no runners with a run scored", func() {
		state := NewGameState()
		play := Play{
			Type: HomeRun,
		}
		next := ApplyPlay(state, play)
		s.Equal(0, next.Outs)
		s.Equal(1, next.Inning)
		s.Equal(Top, next.InningHalf)
		s.Equal(1, next.Teams["team1"].Score)
		s.Equal(0, next.Teams["team2"].Score)
		s.Equal(false, next.Runner.First)
		s.Equal(false, next.Runner.Second)
		s.Equal(false, next.Runner.Third)
	})
}

func (s *GameStateSuite) TestApplyPlay_ErrorRunnerOccupied() {
	s.Run("0 outs runners on first, second and third -> 0 outs bases loaded and 1 run scored", func() {
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
			Type: Error,
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

func (s *GameStateSuite) TestApplyPlay_ErrorRunnerOnFirstAndSecond() {
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
			Type: Error,
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

func (s *GameStateSuite) TestApplyPlay_ErrorRunnerOnFirstAndThird() {
	s.Run("0 outs runners on first and third -> 0 outs runners on first, second and third with a run scored", func() {
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
				Third:  true,
			},
			Outs: 0,
		}
		play := Play{
			Type: Error,
		}
		next := ApplyPlay(&state, play)
		s.Equal(0, next.Outs)
		s.Equal(1, next.Inning)
		s.Equal(Top, next.InningHalf)
		s.Equal(1, next.Teams["team1"].Score)
		s.Equal(0, next.Teams["team2"].Score)
		s.Equal(true, next.Runner.First)
		s.Equal(true, next.Runner.Second)
		s.Equal(false, next.Runner.Third)
	})	
}

func (s *GameStateSuite) TestApplyPlay_ErrorRunnerOnSecondAndThird() {
	s.Run("0 outs runners on second and third -> 0 outs runner on first and third with a run scored", func() {
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
				First:  false,
				Second: true,
				Third:  true,
			},
			Outs: 0,
		}
		play := Play{
			Type: Error,
		}
		next := ApplyPlay(&state, play)
		s.Equal(0, next.Outs)
		s.Equal(1, next.Inning)
		s.Equal(Top, next.InningHalf)
		s.Equal(1, next.Teams["team1"].Score)
		s.Equal(0, next.Teams["team2"].Score)
		s.Equal(true, next.Runner.First)
		s.Equal(false, next.Runner.Second)
		s.Equal(true, next.Runner.Third)
	})	
}

func (s *GameStateSuite) TestApplyPlay_ErrorRunnersOnFirst() {
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
			Type: Error,
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

func (s *GameStateSuite) TestApplyPlay_ErrorRunnersOnSecond() {
	s.Run("0 outs runner on second -> 0 outs runners on first and third", func() {
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
				First:  false,
				Second: true,
				Third:  false,
			},
			Outs: 0,
		}
		play := Play{
			Type: Error,
		}
		next := ApplyPlay(&state, play)
		s.Equal(0, next.Outs)
		s.Equal(1, next.Inning)
		s.Equal(Top, next.InningHalf)
		s.Equal(0, next.Teams["team1"].Score)
		s.Equal(0, next.Teams["team2"].Score)
		s.Equal(true, next.Runner.First)
		s.Equal(false, next.Runner.Second)
		s.Equal(true, next.Runner.Third)
	})	
}

func (s *GameStateSuite) TestApplyPlay_ErrorRunnersOnThird() {
	s.Run("0 outs runner on third -> 0 outs runners on first with a run scored", func() {
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
				First:  false,
				Second: false,
				Third:  true,
			},
			Outs: 0,
		}
		play := Play{
			Type: Error,
		}
		next := ApplyPlay(&state, play)
		s.Equal(0, next.Outs)
		s.Equal(1, next.Inning)
		s.Equal(Top, next.InningHalf)
		s.Equal(1, next.Teams["team1"].Score)
		s.Equal(0, next.Teams["team2"].Score)
		s.Equal(true, next.Runner.First)
		s.Equal(false, next.Runner.Second)
		s.Equal(false, next.Runner.Third)
	})
}

func (s *GameStateSuite) TestApplyPlay_ErrorRunnersEmpty() {
	s.Run("0 outs empty bases -> 0 outs runner on first", func() {
		state := NewGameState()
		play := Play{
			Type: Error,
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

func (s *GameStateSuite) TestApplyPlay_SacrificeBuntRunnerOccupied() {
	s.Run("0 outs runners on first, second and third -> 1 out runners on second and third with 1 run scored", func() {
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
			Type: SacrificeBunt,
		}
		next := ApplyPlay(&state, play)
		s.Equal(1, next.Outs)
		s.Equal(1, next.Inning)
		s.Equal(Top, next.InningHalf)
		s.Equal(1, next.Teams["team1"].Score)
		s.Equal(0, next.Teams["team2"].Score)
		s.Equal(false, next.Runner.First)
		s.Equal(true, next.Runner.Second)
		s.Equal(true, next.Runner.Third)
	})
}

func (s *GameStateSuite) TestApplyPlay_SacrificeBuntRunnerOnFirstAndSecond() {
	s.Run("0 outs runners on first and second -> 1 out runners on second and third", func() {
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
			Type: SacrificeBunt,
		}
		next := ApplyPlay(&state, play)
		s.Equal(1, next.Outs)
		s.Equal(1, next.Inning)
		s.Equal(Top, next.InningHalf)
		s.Equal(0, next.Teams["team1"].Score)
		s.Equal(0, next.Teams["team2"].Score)
		s.Equal(false, next.Runner.First)
		s.Equal(true, next.Runner.Second)
		s.Equal(true, next.Runner.Third)
	})	
}

func (s *GameStateSuite) TestApplyPlay_SacrificeBuntRunnerOnFirstAndThird() {
	s.Run("0 outs runners on first and third -> 1 out runners on second with a run scored", func() {
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
				Third:  true,
			},
			Outs: 0,
		}
		play := Play{
			Type: SacrificeBunt,
		}
		next := ApplyPlay(&state, play)
		s.Equal(1, next.Outs)
		s.Equal(1, next.Inning)
		s.Equal(Top, next.InningHalf)
		s.Equal(1, next.Teams["team1"].Score)
		s.Equal(0, next.Teams["team2"].Score)
		s.Equal(false, next.Runner.First)
		s.Equal(true, next.Runner.Second)
		s.Equal(false, next.Runner.Third)
	})	
}

func (s *GameStateSuite) TestApplyPlay_SacrificeBuntRunnerOnSecondAndThird() {
	s.Run("0 outs runners on second and third -> 1 outs runner on third with a run scored", func() {
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
				First:  false,
				Second: true,
				Third:  true,
			},
			Outs: 0,
		}
		play := Play{
			Type: SacrificeBunt,
		}
		next := ApplyPlay(&state, play)
		s.Equal(1, next.Outs)
		s.Equal(1, next.Inning)
		s.Equal(Top, next.InningHalf)
		s.Equal(1, next.Teams["team1"].Score)
		s.Equal(0, next.Teams["team2"].Score)
		s.Equal(false, next.Runner.First)
		s.Equal(false, next.Runner.Second)
		s.Equal(true, next.Runner.Third)
	})	
}

func (s *GameStateSuite) TestApplyPlay_SacrificeBuntRunnersOnFirst() {
	s.Run("0 outs runner on first -> 1 out runners on second", func() {
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
			Type: SacrificeBunt,
		}
		next := ApplyPlay(&state, play)
		s.Equal(1, next.Outs)
		s.Equal(1, next.Inning)
		s.Equal(Top, next.InningHalf)
		s.Equal(0, next.Teams["team1"].Score)
		s.Equal(0, next.Teams["team2"].Score)
		s.Equal(false, next.Runner.First)
		s.Equal(true, next.Runner.Second)
		s.Equal(false, next.Runner.Third)
	})	
}

func (s *GameStateSuite) TestApplyPlay_SacrificeBuntRunnersOnSecond() {
	s.Run("0 outs runner on second -> 0 outs runners on third", func() {
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
				First:  false,
				Second: true,
				Third:  false,
			},
			Outs: 0,
		}
		play := Play{
			Type: SacrificeBunt,
		}
		next := ApplyPlay(&state, play)
		s.Equal(1, next.Outs)
		s.Equal(1, next.Inning)
		s.Equal(Top, next.InningHalf)
		s.Equal(0, next.Teams["team1"].Score)
		s.Equal(0, next.Teams["team2"].Score)
		s.Equal(false, next.Runner.First)
		s.Equal(false, next.Runner.Second)
		s.Equal(true, next.Runner.Third)
	})	
}

func (s *GameStateSuite) TestApplyPlay_SacrificeBuntRunnersOnThird() {
	s.Run("0 outs runner on third -> 1 out no runners with a run scored", func() {
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
				First:  false,
				Second: false,
				Third:  true,
			},
			Outs: 0,
		}
		play := Play{
			Type: SacrificeBunt,
		}
		next := ApplyPlay(&state, play)
		s.Equal(1, next.Outs)
		s.Equal(1, next.Inning)
		s.Equal(Top, next.InningHalf)
		s.Equal(1, next.Teams["team1"].Score)
		s.Equal(0, next.Teams["team2"].Score)
		s.Equal(false, next.Runner.First)
		s.Equal(false, next.Runner.Second)
		s.Equal(false, next.Runner.Third)
	})
}

func (s *GameStateSuite) TestApplyPlay_SacrificeBuntRunnersEmpty() {
	s.Run("0 outs empty bases -> 1 out no runners", func() {
		state := NewGameState()
		play := Play{
			Type: SacrificeBunt,
		}
		next := ApplyPlay(state, play)
		s.Equal(1, next.Outs)
		s.Equal(1, next.Inning)
		s.Equal(Top, next.InningHalf)
		s.Equal(0, next.Teams["team1"].Score)
		s.Equal(0, next.Teams["team2"].Score)
		s.Equal(false, next.Runner.First)
		s.Equal(false, next.Runner.Second)
		s.Equal(false, next.Runner.Third)
	})
}