package game

func Recalculate(plays []Play) GameState {
	state := NewGameState()
	for _, play := range plays {
		state = RecordPlay(state, play)
	}
	return *state
}
