package testutil

import "github.com/LorenzoDOrtona/Tris_Inception/internal/model/game"

var MoveFirstPLayerWins = [][]int{
	{1, 1, 1, 1},
	{1, 1, 0, 0},
	{0, 0, 0, 1},
	{0, 1, 2, 2},
	{2, 2, 0, 0},
	{0, 0, 1, 2},
	{1, 2, 0, 0},
	{0, 0, 2, 1},
	{2, 1, 0, 0},
	{0, 0, 0, 2},
	{0, 2, 1, 1},
	{1, 1, 2, 0},
	{2, 0, 1, 1},
	{1, 1, 0, 2},
	{0, 2, 2, 2},
	{2, 2, 1, 2},
	{1, 2, 2, 0},
	{2, 0, 2, 1},
	{2, 1, 1, 0},
	{1, 0, 2, 2},
	{2, 2, 2, 0},
	{2, 0, 0, 1},
	{0, 1, 1, 0},
	{1, 0, 0, 0},
	{0, 0, 2, 0},
	{2, 0, 0, 2},
	{0, 2, 0, 0},
	{0, 0, 2, 2},
	{2, 2, 0, 1},
	{0, 1, 1, 1},
	{1, 1, 1, 0},
	{1, 0, 0, 2},
	{1, 2, 1, 0},
	{1, 0, 0, 1},
	{0, 1, 2, 0},
	{2, 0, 0, 0},
}

func CreateGameWaitingOpponent() *game.Game {
	/*
		Creates a game with 2 players of fixed Username and UUID
	*/
	newCreator := game.NewPlayer("CreatorUUID", "CreatorUsername")
	newGame := game.New(&newCreator, nil)
	return newGame
}
func CreateBasicGame() *game.Game {
	/*
		Creates a game with 2 players of fixed Username and UUID
	*/
	newCreator := game.NewPlayer("CreatorUUID", "CreatorUsername")
	newGame := game.New(&newCreator, nil)
	//creating oppenent player
	newGame.CurrentGameState.AddOpponent("OpponentUUID", "OpponentUsername", false)
	return newGame
}
