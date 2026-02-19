package game

import (
	"fmt"
)

type EndState struct {
	MainGame *Game
	GameState
}

func (es *EndState) Activate() {
	fmt.Println("EndState: game ended")
	if es.MainGame != nil {
		es.MainGame.Finished = true
	}
	fmt.Println("WInner is: ", es.MainGame.MainBoard.PlayerWhoCompleted)
}

func (es *EndState) MoveCommand(i, j, x, y int, player Player) error {
	return fmt.Errorf("game already ended")
}
