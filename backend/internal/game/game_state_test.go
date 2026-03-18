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
		s.Equal(0, state.HomeScore)
		s.Equal(0, state.AwayScore)
		s.Equal(Top, state.InningHalf)
		s.Equal(false, state.Runner.First)
		s.Equal(false, state.Runner.Second)
		s.Equal(false, state.Runner.Third)
	})
}



