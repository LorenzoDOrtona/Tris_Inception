package game

type BeginState struct {
	mainGame *Game
	GameState
}

func (gs *BeginState) activate() {
	gs.mainGame.Init()
}
