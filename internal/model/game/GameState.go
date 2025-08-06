package game

import "github.com/google/uuid"

// import "fmt"
type GameState interface {
	Activate()
}

func (gs *GameState) MoveCommand(i, j, x, y int, player uuid.UUID) error {
}
