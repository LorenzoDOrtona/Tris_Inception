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
	if m==1{
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
	}else if m == 2 {
		// modalità test automatico
		if err := RunAutomatedTest(g, GetDefaultTestMoves()); err != nil {
			fmt.Println("Errore durante test automatico:", err)
			return
		}
	}
}
func selectMode() int{
	var valid bool= false
	var m string
	var i int=0
	for !valid{
	fmt.Println("Seleziona la modalità:")
	fmt.Println("1- 1v1 local ")
	m= "2"
	if m=="1"{
		valid=true
		i=1
	}
	if m=="2"{
		valid=true
		i=2
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
///TEST
// RunAutomatedTest esegue una sequenza di mosse (stringhe "i j x y")
// contro il gioco passato. Usa validMove() e MoveCommand come nel main.
// Ritorna nil se la partita finisce correttamente; altrimenti ritorna errore con motivo.
func RunAutomatedTest(g *game.Game, moves []string) error {
	idx := 0

	for !g.Finished{
		if idx >= len(moves) {
			return fmt.Errorf("test moves esaurite ma la partita non è finita")
		}

		move := moves[idx]
		idx++
		fmt.Println("Auto move:", move)

		p, err := validMove(move)
		if err != nil {
			return fmt.Errorf("mossa non valida '%s': %w", move, err)
		}

		// Esegui la mossa usando l'API del gioco
		err = g.CurrentGameState.MoveCommand(p[0], p[1], p[2], p[3], g.CurrentPlaying)
		if err != nil {
			return fmt.Errorf("MoveCommand fallita sulla mossa '%s': %w", move, err)
		}

		fmt.Println("curr played:", g.CurrentPlaying.Username)
	}

	fmt.Println("Partita terminata. Stato Finished =", g.Finished)
	return nil
}

// GetDefaultTestMoves ritorna la lista di mosse cooperative (0-based)
// che ti ho preparato — incollala qui così la chiami facilmente.
func GetDefaultTestMoves() []string {
	return []string{
		"1 1 1 1",
		"1 1 0 0",
		"0 0 0 1",
		"0 1 2 2",
		"2 2 0 0",
		"0 0 1 2",
		"1 2 0 0",
		"0 0 2 1",
		"2 1 0 0",
		"0 0 0 2",
		"0 2 1 1",
		"1 1 2 0",
		"2 0 1 1",
		"1 1 0 2",
		"0 2 2 2",
		"2 2 1 2",
		"1 2 2 0",
		"2 0 2 1",
		"2 1 1 0",
		"1 0 2 2",
		"2 2 2 0",
		"2 0 0 1",
		"0 1 1 0",
		"1 0 0 0",
		"0 0 2 0",
		"2 0 0 2",
		"0 2 0 0",
	}
}
