package game

//import "fmt"s
import (
	"github.com/LorenzoDOrtona/Tris_Inception/internal/model/board"
	"github.com/LorenzoDOrtona/Tris_Inception/internal/model/positionable"
	"github.com/google/uuid"
)

type Game struct {
	GameUuid         uuid.UUID
	Creator          *Player
	Opponent         *Player
	PlayingPlayer    *Player
	ObservingPlayer  *Player
	Winner           *Player
	Looser           *Player
	CurrentGameState GameState
	MainBoard        board.BigBoard
	Finished         bool
}

func New(player, opp *Player) *Game {
	g := Game{
		Creator:  player,
		Opponent: opp,
		// PlayingPlayer lo settiamo al player di default
		PlayingPlayer:   player,
		ObservingPlayer: opp,
		GameUuid:        uuid.New(),
		Finished:        false,
	}
	g.Init()
	return &g
}

/*
Starts the game by activating the first
gameState
*/
func (game *Game) Init() {
	// inizializza la board e lo stato iniziale
	//quindi tutto il model
	game.MainBoard = board.BigBoard{}
	game.CurrentGameState = &BeginState{mainGame: game}
	game.MainBoard.SetupBigBoard()
	game.CurrentGameState.Activate()
}

/*
 */
func (game *Game) ChangePlayerTurn() {
	temp := game.PlayingPlayer
	game.PlayingPlayer = game.ObservingPlayer
	game.ObservingPlayer = temp
}

/*
This functions makes the game procede to next state!
*/
func (g *Game) GoNextState(gs GameState) {
	g.CurrentGameState = gs
	g.CurrentGameState.Activate()
}
func (g *Game) CheckWin(m positionable.Positionable) bool {
	// Check if there is a WINNER
	win := g.MainBoard.CheckWin(m, g.PlayingPlayer.Uuid)
	if win {
		g.Finished = true
	}
	return win
}
