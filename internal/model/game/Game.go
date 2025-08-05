package game

//import "fmt"s
import (
	"github.com/google/uuid"
	"github.com/LorenzoDOrtona/Tris_Inception/internal/model/board"

)

type Game struct {
	playerUuid     uuid.UUID
	opponentUuid   uuid.UUID
	currentPlaying uuid.UUID
	gameUuid       uuid.UUID
	gameState      GameState
	winner         uuid.UUID
	looser         uuid.UUID
	mainBoard      board.BigBoard
}

/*
Starts the game by activating the first
gameState
*/
func (game *Game) Init() {
	game.gameState=&BeginState{mainGame:game}
	game.mainBoard.SetupBigBoard()
	game.gameState.Activate()
}

/*
 */
func ChangePlayerTurn() {

}

/*
 */
