package model
import(
	"fmt"
	"internal/model/positionable/Board"
)
type BigBoard struct {
	mainBoard:=[3][3]Board

}
func setupBigBoard(){
	var boardInit BigBoard
	for i==0;i<3;i++{
		for j==0;j<3;j++{
			boardInit[i][j]=Board.setupBoard()
		}
	}
	return boardInit
}
func (BB BigBoard) String() string{
	var out string
	for i==0;i<3;i++{
		for j==0;j<3;j++{
			out= out+mainBoard[i][j].String()+" "
		}
		out=out+"\n"
	}
}
func (BB BigBoard) print(){
	fmt.Println(BB.String())
}