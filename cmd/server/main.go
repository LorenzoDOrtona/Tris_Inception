package main

import (
	"fmt"
	"github.com/LorenzoDOrtona/Tris_Inception/internal/model/board"
)
func main()  {
	fmt.Println("Ciao gente")
	var BB board.BigBoard
	BB.SetupBigBoard()
	BB.Print()
	
}
