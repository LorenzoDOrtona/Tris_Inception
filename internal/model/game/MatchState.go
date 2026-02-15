package game

import (
	"fmt"
	"time"

	"github.com/LorenzoDOrtona/Tris_Inception/internal/errors"
	"github.com/LorenzoDOrtona/Tris_Inception/internal/model/positionable"
	//"github.com/LorenzoDOrtona/Tris_Inception/internal/model/player"
)

type MatchState struct {
	MainGame *Game
	GameState
	//listener dei turni per i giocatori
	//listener degli ev
}

func (ms *MatchState) Activate() {
	fmt.Println("MatchState: activated")
	// inizializzazioni se necessarie
}

// Returns next state from this one
// Returns next state from this one
func (ms *MatchState) GetNextState(hasOtherChoice bool) GameState {
	return &EndState{MainGame: ms.MainGame}
}
func (gs *MatchState) MoveCommand(i, j, x, y int, player Player) error {
	/*
		Command handler
	*/
	//1) validation
	err := gs.validateMove(i, j, x, y, player)
	if err == nil {
		//2) execution
		gs.executeMove(i, j, x, y, player)
		gs.MainGame.MainBoard.LastTimestamp = int(time.Now().UnixNano())
		//3) check status
		gs.checkStatus(player.MarkS, i, j)
		if !gs.MainGame.Finished {
			gs.MainGame.ChangePlayerTurn()
		}
	} else {
		return err
	}

	return nil
}
func (gs *MatchState) validateMove(i, j, x, y int, player Player) error {
	//this function returns true iff the proposed move is valid
	if !gs.MainGame.MainBoard.AvailableMoves[[4]int{i, j, x, y}] {
		return errors.ErrInvalidMove
	}
	if gs.MainGame.PlayingPlayer.Uuid != player.Uuid {
		return errors.ErrNotYourTurn
	} else {
		return nil
	}
	//we just need to
}
func (gs *MatchState) executeMove(i, j, x, y int, player Player) {
	//1) check if a card is selected
	// or if it is a white place
	//2) if there is a card, activate the effect and
	// 		go to another state in the card
	//3) place the mark in the specified cell

	cell := gs.MainGame.MainBoard.GetCell(i, j, x, y)
	isCard := cell.ImCard()

	m := player.MarkS
	if isCard {
		cell.Selected(player.Uuid, m)
	}
	gs.MainGame.MainBoard.InsertMark(m, i, j, x, y)

	gs.MainGame.MainBoard.ChangeBoardAvailability(i, j, x, y)

}
func (gs *MatchState) checkStatus(MarkS positionable.Mark, i, j int) {
	//1) check if there is a new little win
	gs.MainGame.MainBoard.CheckSmallWin(MarkS, i, j)
	//2) check if someone won definetelly
	weHaveAWinner := gs.MainGame.CheckWin(MarkS)
	//if someone won, we end game
	if weHaveAWinner {
		//go to end state
		gs.MainGame.MainBoard.LastTimestamp = int(time.Now().UnixNano())
		nextGamestate := gs.GetNextState(false)
		gs.MainGame.GoNextState(nextGamestate)
	}
}
