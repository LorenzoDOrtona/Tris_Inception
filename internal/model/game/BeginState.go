package game

import (
	"fmt"
	
	"github.com/LorenzoDOrtona/Tris_Inception/internal/model/board"
	
)
type BeginState struct {
	mainGame *Game
}

func (gs *BeginState) Activate() {
	gs.mainGame.mainBoard = board.BigBoard{}
	gs.mainGame.mainBoard.SetupBigBoard()
	gs.mainGame.mainBoard.Print()
	// passa allo stato Match usando un puntatore
	gs.mainGame.CurrentGameState = &MatchState{mainGame: gs.mainGame}
	gs.mainGame.CurrentGameState.Activate()
}

func (gs *BeginState) MoveCommand(i, j, x, y int, player Player) error {
	// non accetti mosse in BeginState
	fmt.Println("CIAO")
	gs.mainGame.mainBoard.Print()
	return nil
}