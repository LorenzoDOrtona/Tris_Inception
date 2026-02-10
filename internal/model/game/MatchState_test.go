package game_test

import (
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

}
func TestFirstPlayerWins(t *testing.T) {
	TestGame := testutil.CreateBasicGame()
	movArray := testutil.MoveFirstPLayerWins
	for _, v := range movArray {
		TestGame.CurrentGameState.MoveCommand(v[0], v[1], v[2], v[3], *TestGame.PlayingPlayer)
	}
}
