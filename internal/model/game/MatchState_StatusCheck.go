package game
type MatchState_StatusCheck struct{
	
}
// Returns next state from this one
func (cgs MatchState_StatusCheck) GetNextState(hasOtherChoice bool) (ngs GameState) {
	if !hasOtherChoice {
		return EndState{}
	} else {
		return MatchState_Choice{}
	}
}