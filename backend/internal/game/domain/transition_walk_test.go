package game


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

func (s *GameStateSuite) TestApplyPlay_WalkRunnersEmpty() {
	s.Run("0 outs empty bases -> 0 outs runner on first", func() {
		state := NewGameState()
		play := Play{
			Type: Walk,
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

func (s *GameStateSuite) TestApplyPlay_HitByPitchPutsRunnersOnFirst() {
	s.Run("0 outs empty bases -> 0 outs runner on first", func() {
		state := NewGameState()
		play := Play{
			Type: HitByPitch,
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