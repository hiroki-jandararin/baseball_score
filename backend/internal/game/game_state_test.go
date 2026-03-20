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
		s.Equal(0, state.FirstBattingScore)
		s.Equal(0, state.SecondBattingScore)
		s.Equal(Top, state.InningHalf)
		s.Equal(false, state.Runner.First)
		s.Equal(false, state.Runner.Second)
		s.Equal(false, state.Runner.Third)
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
			FirstBattingScore:  0,
			SecondBattingScore: 0,
		}
		state.AddRun(2)
		s.Equal(2, state.FirstBattingScore)
		s.Equal(0, state.SecondBattingScore)
	})
	
	s.Run("adds runs to second batting team in bottom of inning", func() {
		state := GameState{
			Inning:             1,
			InningHalf:         Bottom,
			FirstBattingScore:  0,
			SecondBattingScore: 0,
		}
		state.AddRun(3)
		s.Equal(0, state.FirstBattingScore)
		s.Equal(3, state.SecondBattingScore)
	})	
}