package game

// import "fmt"
type GameState interface {
	//var game Game
	GetNextState(hasOtherChoice bool) GameState
}

// Returns next state from this one
func GetNextState(hasOtherChoice bool) (ngs GameState) {
	return
}
func (cs GameState) GoNextState(hasOtherChoice bool) {
	game.state = GoNextState(hasOtherChoice)
	game.state.activate()
}
