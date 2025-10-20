package game

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/LorenzoDOrtona/Tris_Inception/internal/model/positionable"
	//"github.com/LorenzoDOrtona/Tris_Inception/internal/model/player"
)

type MatchState struct {
	mainGame *Game
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
	return &EndState{mainGame: ms.mainGame}
}
func (gs *MatchState) MoveCommand(i, j, x, y int, player Player) error {
	/*
		Command handler
	*/
	//1) validation
	valid := gs.validateMove(i, j, x, y, player.Uuid)
	if valid {
		//2) execution
		gs.executeMove(i, j, x, y, player)
		//3) check status
		gs.checkStatus(player.MarkS,i,j)
		gs.mainGame.mainBoard.Print()
	} else {
		//e) ERROR, not correct move
		gs.mainGame.mainBoard.Print()
		return fmt.Errorf("not valid")
	}
	
	return nil
}
func (gs *MatchState) validateMove(i, j, x, y int, player uuid.UUID) bool {
	//this function returns true iff the proposed move is valid
	return gs.mainGame.mainBoard.AvailableMoves[[4]int{i,j,x,y}]
	//we just need to 
}
func (gs *MatchState)IsOccupied(i, j, x, y int) bool {
	//if there is something different
	///From empty it's false

	isEmpty:=gs.mainGame.mainBoard.GetCell(i,j,x,y).ImEmpty()
	if (isEmpty) {return false}
	return true
}
func (gs*MatchState) executeMove(i, j, x, y int, player Player) string{
	
	//1) check if a card is selected
	// or if it is a white place
	//2) if there is a card, activate the effect and
	// 		go to another state in the card
	//3) place the mark in the specified cell
	
	cell:=gs.mainGame.mainBoard.GetCell(i,j,x,y)
	isCard:=cell.ImCard()
	
	m:=player.MarkS
	if (isCard){cell.Selected(player.Uuid,m)}
	gs.mainGame.mainBoard.InsertMark(m,i,j,x,y)
	
	gs.mainGame.mainBoard.ChangeBoardAvailability(i,j,x,y)
	//change current player
	gs.mainGame.ChangePlayerTurn()


	return ""
}
func (gs *MatchState) checkStatus(MarkS positionable.Mark,i,j int) {
	//1) check if there is a new little win
	gs.mainGame.mainBoard.CheckSmallWin(MarkS,i,j)
	//2) check if someone won definetelly
	weHaveAWinner := gs.mainGame.CheckWin(MarkS)
	//if someone won, we end game
	if weHaveAWinner {
		gs.mainGame.GoNextState(&EndState{mainGame: gs.mainGame})
	}
}


func finishMove(i, j, x, y int) {
	//1) update the available plane to mark
	//2) update player to play
}

