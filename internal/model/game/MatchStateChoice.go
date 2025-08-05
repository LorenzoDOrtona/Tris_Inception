package game

type MatchStateChoice struct {
}

// Returns next state from this one
func (cgs MatchStateChoice) GetNextState(hasOtherChoice bool) (ngs GameState) {
	return MatchStateValidation{}
}
