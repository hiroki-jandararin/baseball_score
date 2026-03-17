package game

import "github.com/stretchr/testify/suite"

type RunnerStateSuite struct {
	suite.Suite
}

func (s *RunnerStateSuite) TestRunnerState() {
	s.Run("empty bases", func() {
		runner := RunnerState{
			First:  false,
			Second: false,
			Third:  false,
		}
		s.Equal(false, runner.First)
		s.Equal(false, runner.Second)
		s.Equal(false, runner.Third)
	})
	s.Run("bases loaded", func() {
		runner := RunnerState{
			First:  true,
			Second: true,
			Third:  true,
		}
		s.Equal(true, runner.First)
		s.Equal(true, runner.Second)
		s.Equal(true, runner.Third)
	})
	s.Run("first base occupied", func() {
		runner := RunnerState{
			First:  true,
			Second: false,
			Third:  false,
		}
		s.Equal(true, runner.First)
		s.Equal(false, runner.Second)
		s.Equal(false, runner.Third)
	})
}
