package board
import (
	"github.com/LorenzoDOrtona/Tris_Inception/internal/model/positionable"
)
type Board struct {
	board [3][3]positionable.Positionable

}
func (B * Board)SetupBoard() {
	/*
	Creates the board, filling it with positionables
	*/
	for i:=0;i<3;i++{
		for j:=0;j<3;j++{
			B.board[i][j] = positionable.GetPositionable()
		}
	}
	
}

