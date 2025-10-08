package game

import "github.com/google/uuid"

// import "fmt"
type GameState interface {
	Activate()
	MoveCommand(i, j, x, y int, player uuid.UUID) error
}

