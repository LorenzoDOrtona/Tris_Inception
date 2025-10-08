package game

//import "fmt"s
import (
	"github.com/google/uuid"
	"github.com/LorenzoDOrtona/Tris_Inception/internal/model/board"

)

type Game struct {
	playerUuid     uuid.UUID
	opponentUuid   uuid.UUID
	CurrentPlaying uuid.UUID
	gameUuid       uuid.UUID
	CurrentGameState      GameState
	winner         uuid.UUID
	looser         uuid.UUID
	mainBoard      board.BigBoard
	Finished 	bool
}
func New(player,opp uuid.UUID) *Game{
	g := Game{
		playerUuid:   player,
		opponentUuid: opp,
		// CurrentPlaying lo settiamo al player di default
		CurrentPlaying: player,
		gameUuid:       uuid.New(),
		Finished:       false,
	}
	return &g
}
/*
Starts the game by activating the first
gameState
*/
func (game *Game) Init() {
	// inizializza la board e lo stato iniziale
	game.mainBoard = board.BigBoard{}
	game.CurrentGameState=&BeginState{mainGame:game}
	game.mainBoard.SetupBigBoard()
	game.CurrentGameState.Activate()
}

/*
 */
func ChangePlayerTurn() {

}

/*
 */
func (g*Game) GoNextState(gs GameState){
	g.CurrentGameState=gs
	g.CurrentGameState.Activate()
}
