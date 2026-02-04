package game

//"github.com/google/uuid"

// import "fmt"
type GameState interface {
	Activate()
	MoveCommand(i, j, x, y int, player Player) error
	AddOpponent(uuid string, name string, bot bool)
}
