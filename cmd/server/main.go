package main

import (
	"fmt"

	"github.com/LorenzoDOrtona/Tris_Inception/internal/model/game"
)

func main() {
	fmt.Println("Ciao gente")

	g := CreateGame()
	g.Init()

}
func CreateGame() game.Game {
	mainGame := game.Game{}
	return mainGame

}
