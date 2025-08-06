package main

import (
	"fmt"

	"github.com/LorenzoDOrtona/Tris_Inception/internal/model/game"
)

func main() {
	fmt.Println("Ciao gente")
	m:=selectMode()
	if m==1{

	g:= game.Game{}
	g.Initial()
	
	for {
		var move string=""
		for ! g.Finished{
			fmt.Println("Move: ")
			fmt.Scanln(&move)
			err:=g.CurrentGameState.MoveCommand(1,1,1,1,g.CurrentPlaying)
			if err==nil{

			}
		}
	}
	}
}
func selectMode() int{
	var valid bool= false
	var m string
	var i int=0
	for !valid{
	fmt.Println("Seleziona la modalità: ")
	fmt.Scanln(&m)
	if m=="1"{
		valid=true
		i=1
	}
	}	
	return i
}
