package game

//import "fmt"s
import (
	"sync"

	"github.com/LorenzoDOrtona/Tris_Inception/internal/model/board"
	"github.com/LorenzoDOrtona/Tris_Inception/internal/model/positionable"
	"github.com/google/uuid"
)

type Game struct {
	Mu               sync.Mutex
	GameUuid         uuid.UUID
	Creator          *Player
	Opponent         *Player
	PlayingPlayer    *Player
	ObservingPlayer  *Player
	Winner           *Player
	Loser            *Player
	CurrentGameState GameState
	MainBoard        board.BigBoard
	Finished         bool
}

func New(player, opp *Player) *Game {
	g := Game{
		Creator:  player,
		Opponent: opp,
		// Set PlayingPlayer to the default player
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
	// Initializes the board and the initial state,
	// thus the whole model
	game.MainBoard = board.BigBoard{}
	game.CurrentGameState = &BeginState{MainGame: game}
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
This function makes the game proceed to the next state!
*/
func (g *Game) GoNextState(gs GameState) {
	g.CurrentGameState = gs
	g.CurrentGameState.Activate()
}
func (g *Game) CheckWin(m positionable.Mark) bool {
	// Check if there is a WINNER
	win := g.MainBoard.CheckWin(m, g.PlayingPlayer.Uuid)
	if win {
		g.Finished = true
		//player turn already changed so if player won it's not it's turn now (swapped before)
		g.Winner = g.PlayingPlayer
	}
	return win
}
