package game_test

import (
	"fmt"
	"testing"

	"github.com/LorenzoDOrtona/Tris_Inception/internal/model/testutil"
)

func TestValidateMove(t *testing.T) {
	TestGame := testutil.CreateBasicGame()
	err := TestGame.CurrentGameState.MoveCommand(1, 1, 1, 1, *TestGame.ObservingPlayer)
	if err == nil {
		t.Errorf("Wrong player allowed to play when it should not")
	}
	err = TestGame.CurrentGameState.MoveCommand(1, 1, 1, 1, *TestGame.PlayingPlayer)
	if err != nil {
		t.Errorf("Correct player cannot play when it could")
	}
	errs := TestGame.CurrentGameState.MoveCommand(1, 1, 1, 1, *TestGame.ObservingPlayer)
	if errs == nil {
		t.Errorf("Correct player played where it CANNOT")
	}

}

func TestFirstPlayerWins(t *testing.T) {
	TestGame := testutil.CreateBasicGame()
	movArray := testutil.MoveFirstPLayerWins

	fmt.Printf("Starting player: %s (%s)\n", TestGame.PlayingPlayer.Username, TestGame.PlayingPlayer.Uuid)

	for i, v := range movArray {
		currentPlayer := TestGame.PlayingPlayer
		fmt.Printf("Move %d: [%d,%d,%d,%d] - Player: %s (%s)\n",
			i, v[0], v[1], v[2], v[3], currentPlayer.Username, currentPlayer.Uuid)

		err := TestGame.CurrentGameState.MoveCommand(v[0], v[1], v[2], v[3], *currentPlayer)
		if err != nil {
			fmt.Printf("ERROR at move %d: %v\n", i, err)
			fmt.Printf("Expected: %s, PlayingPlayer after: %s\n",
				currentPlayer.Uuid, TestGame.PlayingPlayer.Uuid)
			t.Fatalf("Move %d failed: %v | Move data: %v", i, err, v)
		}

		fmt.Printf("After move %d - Next player: %s (%s)\n\n",
			i, TestGame.PlayingPlayer.Username, TestGame.PlayingPlayer.Uuid)
		TestGame.MainBoard.Print()
	}

	//TestGame.MainBoard.Print()

	if TestGame.Winner != nil {
		fmt.Printf("Winner: %s (%s)\n", TestGame.Winner.Username, TestGame.Winner.Uuid)
	}

}

func TestNoMoveBeforeOpponent(t *testing.T) {
	TestGame := testutil.CreateGameWaitingOpponent()
	err := TestGame.CurrentGameState.MoveCommand(1, 1, 1, 1, *TestGame.PlayingPlayer)
	if err == nil {
		t.Errorf("Payer allowed to play when it should not")
	}

}
