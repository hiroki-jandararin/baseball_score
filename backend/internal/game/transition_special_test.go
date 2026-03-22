package game

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
		next := RecordPlay(&state, play)
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
		next := RecordPlay(&state, play)
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
		next := RecordPlay(&state, play)
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
		next := RecordPlay(&state, play)
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
		next := RecordPlay(&state, play)
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
		next := RecordPlay(&state, play)
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
		next := RecordPlay(&state, play)
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
		next := RecordPlay(state, play)
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
		next := RecordPlay(&state, play)
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
		next := RecordPlay(&state, play)
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
		next := RecordPlay(&state, play)
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
		next := RecordPlay(&state, play)
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
		next := RecordPlay(&state, play)
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
		next := RecordPlay(&state, play)
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
		next := RecordPlay(&state, play)
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
		next := RecordPlay(state, play)
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

func (s *GameStateSuite) TestApplyPlay_StealRunnerOccupied() {
	s.Run("0 outs runners on first, second and third -> 0 outs runners on second and third with 1 run scored", func() {
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
			Type: Steal,
		}
		next := RecordPlay(&state, play)
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

func (s *GameStateSuite) TestApplyPlay_StealRunnerOnFirstAndSecond() {
	s.Run("0 outs runners on first and second -> 0 outs runners on second and third", func() {
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
			Type: Steal,
		}
		next := RecordPlay(&state, play)
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

func (s *GameStateSuite) TestApplyPlay_StealRunnerOnFirstAndThird() {
	s.Run("0 outs runners on first and third -> 0 outs runners on second with a run scored", func() {
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
			Type: Steal,
		}
		next := RecordPlay(&state, play)
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

func (s *GameStateSuite) TestApplyPlay_StealRunnerOnSecondAndThird() {
	s.Run("0 outs runners on second and third -> 0 outs runner on third with a run scored", func() {
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
			Type: Steal,
		}
		next := RecordPlay(&state, play)
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

func (s *GameStateSuite) TestApplyPlay_StealRunnersOnFirst() {
	s.Run("0 outs runner on first -> 0 outs runners on second", func() {
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
			Type: Steal,
		}
		next := RecordPlay(&state, play)
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

func (s *GameStateSuite) TestApplyPlay_StealRunnersOnSecond() {
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
			Type: Steal,
		}
		next := RecordPlay(&state, play)
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

func (s *GameStateSuite) TestApplyPlay_StealRunnersOnThird() {
	s.Run("0 outs runner on third -> 0 outs no runners with a run scored", func() {
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
			Type: Steal,
		}
		next := RecordPlay(&state, play)
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