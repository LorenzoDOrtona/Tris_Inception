package game

import (
	"github.com/google/uuid"
)
type MatchState struct{
	mainGame *Game
}
// Returns next state from this one
func (cgs MatchState) GetNextState(hasOtherChoice bool) (ngs GameState) {
	return &EndState{}
}
func (gs *MatchState)moveCommand(i,j,x,y int, player uuid.UUID) {
	/*
	Command handler
	*/
	//1) validation 
	valid:=gs.validateMove(i,j,x,y,player)
	if valid{
		//2) execution
		executeMove(i,j,x,y,player)
		//3) check status
		gs.checkStatus()
		//4
		finishMove(i,j,x,y)
	}else{
		//e) ERROR, not correct move
	}
	
}
func (gs *MatchState)validateMove(i,j,x,y int, player uuid.UUID) bool{
	//returns True or False depensing on move avalability
	occupied:=IsOccupied(i,j,x,y)
	yourTurn:= player== gs.mainGame.currentPlaying
	
	if occupied || !yourTurn{
		return false
	}
	return true
}
func IsOccupied(i,j,x,y int) bool{
	//if there is something different 
	///From card or empty it's false
	return true
}
func executeMove(i,j,x,y int, player uuid.UUID) {
	//1) check if a card is selected 
	// or if it is a white place
	//2) if there is a card, activate the effect and 
	// 		go to another state in the card
	//3) place the mark in the specified cell
	
}
func (gs *MatchState)checkStatus(){
	//1) check if there is a new little win
	checkNewSmallWins()
	//2) check if someone won definetelly
	weHaveAWinner:=checkWin()
	//if someone won, we end game
	if weHaveAWinner{
		gs.GoNextState()
	}
}
func checkWin() bool{
	// Check if there is a WINNER
	return false
}
func checkNewSmallWins(){
	//
}
func finishMove(i,j,x,y int){
	//1) update the available plane to mark
	//2) update player to play
}
func (gs *MatchState) GoNextState() {
	gs.mainGame.gameState = &EndState{}
	gs.mainGame.gameState.Activate()
}