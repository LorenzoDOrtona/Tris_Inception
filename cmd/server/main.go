package main

import (
	"fmt"
	"bufio"
	"strconv"
	"strings"
	"os"
	"github.com/LorenzoDOrtona/Tris_Inception/internal/model/game"
	"github.com/LorenzoDOrtona/Tris_Inception/internal/model/positionable"
	"github.com/google/uuid"
)

func main() {
	fmt.Println("Ciao gente")
	m:=selectMode()
	if m==1{

	pl1 := game.Player{
		Uuid:     uuid.New(),
		Username: "pippo",
		MarkS:   positionable.Mark{Marktype: 1},
	}
	pl2 := game.Player{
		Uuid:     uuid.New(),
		Username: "peppa",
		MarkS:   positionable.Mark{Marktype: 2},
	}

	g:=game.New(pl1,pl2)
	g.Init()
	
	for {
		for ! g.Finished{
			fmt.Println("Move: ")
			reader := bufio.NewReader(os.Stdin)
			move, _ := reader.ReadString('\n')
    		move = strings.TrimSpace(move)
			p,err:=validMove(move)
			if err==nil{

				err=g.CurrentGameState.MoveCommand(p[0],p[1],p[2],p[3],g.CurrentPlaying)
				fmt.Println("curr played: ",g.CurrentPlaying.Username)
			}else{
				fmt.Println("Error:",err)
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
	fmt.Println("Seleziona la modalità: \n")
	fmt.Println("1- 1v1 local \n")
	fmt.Scanln(&m)
	if m=="1"{
		valid=true
		i=1
	}
	}	
	return i
}
func validMove(s string) ([]int,error){
	sideLenght:=9
	//array of the integer positions
	var p []int
	
	n:=strings.Split(s," ")
	
	if len(n)!=4{
		fmt.Println("lunghezza", len(n))
		return nil,fmt.Errorf("Input is too long!")
		
	}
	
	//conversion from string to int
	for u:=0;u<4;u++{
		num,err :=strconv.Atoi(n[u])
		if (err!=nil){
			return nil,fmt.Errorf("You didn't insert integers!")
		}
		if(num>(sideLenght/3)-1){
			return nil,fmt.Errorf("Input of move out of bound!")
		}
		p=append(p,num)
		
	}
	return p,nil
	
}