package game

//import "fmt"s
import (
	"github.com/google/uuid"
	"github.com/LorenzoDOrtona/Tris_Inception/internal/model/board"
	"github.com/LorenzoDOrtona/Tris_Inception/internal/model/positionable"
	

)

type Game struct {
	playerUuid     uuid.UUID
	opponentUuid   uuid.UUID
	CurrentPlaying Player
	NotCurrentlyPlaying Player
	gameUuid       uuid.UUID
	CurrentGameState      GameState
	winner         uuid.UUID
	looser         uuid.UUID
	mainBoard      board.BigBoard
	Finished 	bool
}
func New(player,opp Player) *Game{
	g := Game{
		playerUuid:   player.Uuid,
		opponentUuid: opp.Uuid,
		// CurrentPlaying lo settiamo al player di default
		CurrentPlaying: player,
		NotCurrentlyPlaying: opp,
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
	//quindi tutto il model
	game.mainBoard = board.BigBoard{}
	game.CurrentGameState=&BeginState{mainGame:game}
	game.mainBoard.SetupBigBoard()
	game.CurrentGameState.Activate()
}

/*
 */
func (game * Game)ChangePlayerTurn() {
	temp:=game.CurrentPlaying
	game.CurrentPlaying=game.NotCurrentlyPlaying
	game.NotCurrentlyPlaying=temp
}

/*
 */
func (g*Game) GoNextState(gs GameState){
	g.CurrentGameState=gs
	g.CurrentGameState.Activate()
}
func (g *Game)CheckWin(m positionable.Positionable) bool {
	// Check if there is a WINNER
	win:=g.mainBoard.CheckWin(m,g.CurrentPlaying.Uuid)
	return win
}