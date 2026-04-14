package game

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type GameStateSuite struct {
	suite.Suite
}

func TestGameStateSuite(t *testing.T) {
	suite.Run(t, new(GameStateSuite))
}

func (s *GameStateSuite) TestInitializeGameState() {
	s.Run("initial state", func() {
		state := NewGameState()
		s.Equal(1, state.Inning)
		s.Equal(0, state.Outs)
		s.NotNil(state.Teams)
		s.Equal(2, len(state.Teams))

		team1, ok1 := state.Teams[TeamID("team1")]
		team2, ok2 := state.Teams[TeamID("team2")]

		s.True(ok1)
		s.True(ok2)

		s.Equal(TeamID("team1"), team1.TeamID)
		s.Equal(TeamID("team2"), team2.TeamID)
		s.Equal(Top, state.InningHalf)
		s.Equal(false, state.Runner.First)
		s.Equal(false, state.Runner.Second)
		s.Equal(false, state.Runner.Third)
		s.Equal(0, team1.Score)
		s.Equal(0, team2.Score)
	})
}

func (s *GameStateSuite) TestCurrentBattingTeamID() {
	s.Run("top of inning returns first batting team", func() {
		state := GameState{
			Inning:              1,
			InningHalf:          Top,
			FirstBattingTeamID:  TeamID("team1"),
			SecondBattingTeamID: TeamID("team2"),
		}
		s.Equal(TeamID("team1"), state.CurrentBattingTeamID())
	})

	s.Run("bottom of inning returns second batting team", func() {
		state := GameState{
			Inning:              1,
			InningHalf:          Bottom,
			FirstBattingTeamID:  TeamID("team1"),
			SecondBattingTeamID: TeamID("team2"),
		}
		s.Equal(TeamID("team2"), state.CurrentBattingTeamID())
	})
}

func (s *GameStateSuite) TestAddRun() {
	s.Run("adds runs to first batting team in top of inning", func() {
		state := GameState{
			Inning:             1,
			InningHalf:         Top,
			FirstBattingTeamID:  TeamID("team1"),
			SecondBattingTeamID: TeamID("team2"),
			Teams: map[TeamID]*TeamState{
				"team1": {TeamID: "team1", Score: 0},
				"team2": {TeamID: "team2", Score: 0},
			},
		}
		state.AddRun(2)
		s.Equal(2, state.Teams["team1"].Score)
		s.Equal(0, state.Teams["team2"].Score)
	})
	
	s.Run("adds runs to second batting team in bottom of inning", func() {
		state := GameState{
			Inning:             1,
			InningHalf:         Bottom,
			FirstBattingTeamID:  TeamID("team1"),
			SecondBattingTeamID: TeamID("team2"),
			Teams: map[TeamID]*TeamState{
				"team1": {TeamID: "team1", Score: 0},
				"team2": {TeamID: "team2", Score: 0},
			},
		}
		state.AddRun(3)
		s.Equal(0, state.Teams["team1"].Score)
		s.Equal(3, state.Teams["team2"].Score)
	})	
}