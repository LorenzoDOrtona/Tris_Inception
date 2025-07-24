package model
type MatchState_Validation struct{
	
}
// Returns next state from this one
func (cgs MatchState_Validation) GetNextState(hasOtherChoice bool) (ngs GameState) {
	//from validation I go to exec
	return MatchState_Execution{}
}