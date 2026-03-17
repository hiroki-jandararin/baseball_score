package game

func ApplyPlay(state *GameState, play Play) *GameState {
	next := *state

	switch play.Type {
	case "strikeout":
		next.Outs++
	}

	return &next
}