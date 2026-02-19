package game

import (
	"fmt"
	"time"

	"github.com/LorenzoDOrtona/Tris_Inception/internal/errors"

	"github.com/LorenzoDOrtona/Tris_Inception/internal/model/board"
	"github.com/LorenzoDOrtona/Tris_Inception/internal/model/positionable"
)

type BeginState struct {
	MainGame *Game
}

func (gs *BeginState) Activate() {
	fmt.Print("Begin STATE!")
	gs.MainGame.MainBoard = board.BigBoard{}
	gs.MainGame.MainBoard.SetupBigBoard()
	gs.MainGame.MainBoard.Print()
}
func (gs *BeginState) AddOpponent(uuid string, name string, bot bool) {
	m := gs.MainGame.Creator.MarkS
	oppPlayer := Player{
		Uuid:     uuid,
		Username: name,
		MarkS:    positionable.OppositeMark(m),
	}
	//giving opponent player to the game object
	gs.MainGame.Opponent = &oppPlayer
	//
	gs.MainGame.PlayingPlayer = gs.MainGame.Creator
	gs.MainGame.ObservingPlayer = gs.MainGame.Opponent
	//update timestamp
	gs.MainGame.MainBoard.LastTimestamp = int(time.Now().UnixNano())
	gs.MainGame.GoNextState(&MatchState{MainGame: gs.MainGame})

}
func (gs *BeginState) MoveCommand(i, j, x, y int, player Player) error {
	// non accetti mosse in BeginState
	gs.MainGame.MainBoard.Print()
	return errors.ErrNotYourTurn
}
