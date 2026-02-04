package game

import (
	"time"

	"github.com/LorenzoDOrtona/Tris_Inception/internal/model/board"
	"github.com/LorenzoDOrtona/Tris_Inception/internal/model/positionable"
)

type BeginState struct {
	mainGame *Game
}

func (gs *BeginState) Activate() {
	gs.mainGame.MainBoard = board.BigBoard{}
	gs.mainGame.MainBoard.SetupBigBoard()
	gs.mainGame.MainBoard.Print()
}
func (gs *BeginState) AddOpponent(uuid string, name string, bot bool) {
	m := gs.mainGame.Creator.MarkS
	oppPlayer := Player{
		Uuid:     uuid,
		Username: name,
		MarkS:    positionable.OppositeMark(m),
	}
	//giving opponent player to the game object
	gs.mainGame.Opponent = &oppPlayer
	//
	gs.mainGame.PlayingPlayer = gs.mainGame.Creator
	gs.mainGame.ObservingPlayer = gs.mainGame.Opponent
	//update timestamp
	gs.mainGame.MainBoard.LastTimestamp = int(time.Now().UnixNano())
	gs.mainGame.GoNextState(&MatchState{mainGame: gs.mainGame})

}
func (gs *BeginState) MoveCommand(i, j, x, y int, player Player) error {
	// non accetti mosse in BeginState
	gs.mainGame.MainBoard.Print()
	return nil
}
