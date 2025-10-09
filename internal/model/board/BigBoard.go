package board

import (
	"fmt"
	"github.com/LorenzoDOrtona/Tris_Inception/internal/model/positionable"
)

type BigBoard struct {
	mainBoard [3][3]Board
}

func (B *BigBoard) SetupBigBoard() {

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			B.mainBoard[i][j].SetupBoard()
		}
	}

}
func (BB BigBoard) String() string {
	var out string
	for i := 0; i < 9; i++ {
		if i%3 == 0 && i/3>0 {
			out+="---------------------\n"
		}
		for j := 0; j < 9; j++ {
			if j%3 == 0 && j/3>0 {
				out+="| "
			}
			out = out + BB.mainBoard[i/3][j/3].board[i%3][j%3].String() + " "
		}
		out = out + "\n"
		
	}
	return out
}
func (BB *BigBoard) GetCell(x,y,z,k int) positionable.Positionable{
	return BB.mainBoard[x][y].board[z][k]
}
//func update sector
func (BB *BigBoard) InsertMark(m positionable.Mark,i,j,x,y int) {
	BB.mainBoard[i][j].board[x][y]=m
}
func (BB BigBoard) Print() {
	fmt.Println(BB.String())
}
