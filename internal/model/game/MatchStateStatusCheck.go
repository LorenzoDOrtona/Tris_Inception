package game
type MatchStateStatusCheck struct{
	
}
// Returns next state from this one
func (cgs MatchStateStatusCheck) GetNextState(hasOtherChoice bool) (ngs GameState) {
	if !hasOtherChoice {
		return EndState{}
	} else {
		return MatchStateChoice{}
	}
}