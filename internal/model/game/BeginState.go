package game

import "github.com/LorenzoDOrtona/Tris_Inception/internal/model/board"

type BeginState struct {
	mainGame *Game
}

func (gs *BeginState) Activate() {
	gs.mainGame.mainBoard=board.BigBoard{}
	gs.mainGame.mainBoard.SetupBigBoard()
	gs.mainGame.mainBoard.Print()
	gs.mainGame.CurrentGameState=MatchState{}
	gs.mainGame.CurrentGameState.Activate()
}
