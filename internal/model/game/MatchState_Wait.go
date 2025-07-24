package model

type MatchState_Wait struct {
}

func (cgs MatchState_Wait) GetNextState(hasOtherChoice bool) (ngs GameState) {
	//from validation I go to exec
	return MatchState_Execution{}
}
