package game

type MatchStateExecution struct {

}

// Returns next state from this one
// If I ended myMove (marked position) I go to StatusCheck,
// If something must be completed (card effect activated), I come back to choice
func (cgs MatchStateExecution) GetNextState(hasOtherChoice bool) (ngs GameState) {
	if hasOtherChoice{ 
		return MatchStateChoice{}
	}else{
		return MatchStateStatusCheck{}
	}
}
