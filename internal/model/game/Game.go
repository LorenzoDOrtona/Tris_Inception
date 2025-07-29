package model

//import "fmt"s
import (
	"github.com/google/uuid"
)

type Game struct {
	playerUuid     uuid.UUID
	opponentUuid   uuid.UUID
	currentPlaying uuid.UUID
	gameUuid       uuid.UUID
	gameState      GameState
	winner         uuid.UUID
	looser         uuid.UUID
	mainBoard      BigBoard
}

/*
Starts the game by activating the first
gameState
*/
func init() {
	
}

/*
 */
func changePlayerTurn() {

}

/*
 */
