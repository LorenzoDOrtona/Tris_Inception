package game

type MatchState_Execution struct {

}

// Returns next state from this one
// If I ended myMove (marked position) I go to StatusCheck,
// If something must be completed (card effect activated), I come back to choice
func (cgs MatchState_Execution) GetNextState(hasOtherChoice bool) (ngs GameState) {
	if hasOtherChoice{ 
		return MatchState_Choice{}
	}else{
		return MatchState_StatusCheck{}
	}
}
