package model

// import "fmt"
type GameState interface {
	GetNextState(hasOtherChoice bool) GameState
}

// Returns next state from this one
func GetNextState(hasOtherChoice bool) (ngs GameState) {
	return
}
