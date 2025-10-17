package board
import (
	"github.com/LorenzoDOrtona/Tris_Inception/internal/model/positionable"
)
type Board struct {
	Board [3][3]positionable.Positionable
	IsComplete bool

}
func (B * Board)SetupBoard() {
	/*
	Creates the board, filling it with positionables
	*/
	for i:=0;i<3;i++{
		for j:=0;j<3;j++{
			
			B.Board[i][j] = positionable.GetPositionable()
		}
	}
	B.IsComplete=false
	
}

