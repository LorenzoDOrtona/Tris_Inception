package game

import (
	"fmt"

	"github.com/LorenzoDOrtona/Tris_Inception/internal/model/board"
)

type BeginState struct {
	mainGame *Game
}

func (gs *BeginState) Activate() {
	gs.mainGame.MainBoard = board.BigBoard{}
	gs.mainGame.MainBoard.SetupBigBoard()
	gs.mainGame.MainBoard.Print()
	// passa allo stato Match usando un puntatore
	gs.mainGame.CurrentGameState = &MatchState{mainGame: gs.mainGame}
	gs.mainGame.CurrentGameState.Activate()
}

func (gs *BeginState) MoveCommand(i, j, x, y int, player Player) error {
	// non accetti mosse in BeginState
	fmt.Println("CIAO")
	gs.mainGame.MainBoard.Print()
	return nil
}
