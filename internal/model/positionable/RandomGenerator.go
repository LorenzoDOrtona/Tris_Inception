package model
import (
    "fmt"
    "math/rand"
    "time"
	"internal/model/positionable/Positionable"
)
//config
var cardChance uint:=10 //one in 10 should be a card
var M unit:=100
func randomGen() res int{
	/*
	generates a random number in 0,M range
	*/
	 // 1) Semina il generatore con un valore che cambia
	 rand.Seed(time.Now().UnixNano())
	 x := rand.Intn(M)         // 0 <= x < M
	return x
}

func isThereSpecialPostionable() bool{
	/*
	Generates a random number between 0 and M (example 100)
	If the result x (49) is more than M/x (10), we are NOT having 
	a special positionable, function returns false
	Otherwise we have true
	*/
	var x int :=randomGen()
	var cutOff float:=M/cardChance
	if x<=cutOff{
		return true
	}else{
		return false
	}
}
func getPositionable() Positionable{
	/*
	returns a random positionable, based on chances,
	of each type
	*/
	return
}