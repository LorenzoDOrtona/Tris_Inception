package model

type MatchState_Validation struct {
}

// Returns next state from this one
func (cgs MatchState_Validation) GetNextState(hasOtherChoice bool) (ngs GameState) {
	//from validation I go to exec
	if !hasOtherChoice {
		return MatchState_Execution{}
	} else {
		return MatchState_Choice{}
	}
}
