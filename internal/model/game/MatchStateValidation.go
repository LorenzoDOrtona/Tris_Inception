package game

type MatchStateValidation struct {
}

// Returns next state from this one
func (cgs MatchStateValidation) GetNextState(hasOtherChoice bool) (ngs GameState) {
	//from validation I go to exec
	if !hasOtherChoice {
		return MatchStateExecution{}
	} else {
		return MatchStateChoice{}
	}
}
