package game

import (
	//"github.com/google/uuid"
	//"github.com/LorenzoDOrtona/Tris_Inception/internal/model/player"
)

// import "fmt"
type GameState interface {
	Activate()
	MoveCommand(i, j, x, y int, player Player) error
}

