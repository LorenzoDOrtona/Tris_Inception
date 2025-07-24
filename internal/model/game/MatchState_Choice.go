package model

type MatchState_Choice struct {
}

// Returns next state from this one
func (cgs MatchState_Choice) GetNextState(hasOtherChoice bool) (ngs GameState) {
	return MatchState_Validation{}
}
