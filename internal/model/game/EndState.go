package game

import (
	"fmt"

	"github.com/google/uuid"
)

type EndState struct {
	mainGame *Game
}

func (es *EndState) Activate() {
	fmt.Println("EndState: game ended")
	if es.mainGame != nil {
		es.mainGame.Finished = true
	}
}

func (es *EndState) MoveCommand(i, j, x, y int, player uuid.UUID) error {
	return fmt.Errorf("game already ended")
}
