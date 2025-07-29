package model
import (
    "fmt"
	"internal/model/positionable/RandomGenerator"
)
type Board struct {
	board:=[3][3]Positionable

}
func setupBoard(){
	/*
	Creates the board, filling it with positionables
	*/
	var boardInit Board
	for i==0;i<3;i++{
		for j==0;j<3;j++{
			boardInit[i][j]=RandomGenerator.getPositionable()
		}
	}
	return boardInit
}

