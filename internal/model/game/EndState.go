package game

import (
	"fmt"
	
)

type EndState struct {
	mainGame *Game
	GameState
}

func (es *EndState) Activate() {
	fmt.Println("EndState: game ended")
	if es.mainGame != nil {
		es.mainGame.Finished = true
	}
	fmt.Println("WInner is: ",es.mainGame.mainBoard.PlayerWhoCompleted)
}

func (es *EndState) MoveCommand(i, j, x, y int, player Player) error {
	return fmt.Errorf("game already ended")
}
