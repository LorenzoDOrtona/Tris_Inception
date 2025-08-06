package game

//import "fmt"s
import (
	"github.com/google/uuid"
	"github.com/LorenzoDOrtona/Tris_Inception/internal/model/board"

)

type Game struct {
	PlayerUuid     uuid.UUID
	OpponentUuid   uuid.UUID
	CurrentPlaying uuid.UUID
	gameUuid       uuid.UUID
	CurrentGameState      GameState
	winner         uuid.UUID
	looser         uuid.UUID
	mainBoard      board.BigBoard
	Finished 	bool=false
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
