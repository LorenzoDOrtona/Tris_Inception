package board

import (
	"fmt"
	"github.com/LorenzoDOrtona/Tris_Inception/internal/model/positionable"
)

type BigBoard struct {
	mainBoard [3][3]Board
	AvailableBoards map[[2]int]bool

}

func (B *BigBoard) SetupBigBoard() {
	//initialize the map
	B.AvailableBoards=make(map[[2]int]bool)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			
			B.mainBoard[i][j].SetupBoard()
			//setting the board accessible
			B.AvailableBoards[[2]int{i,j}]=true
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
			if ( BB.mainBoard[i/3][j/3].IsComplete==true){
				out = out + BB.mainBoard[i/3][j/3].MarkCompleted.String() + " "
			}else{
			out = out + BB.mainBoard[i/3][j/3].Board[i%3][j%3].String() + " "
		}}
		out = out + "\n"
		
	}
	return out
}
func (BB *BigBoard) GetCell(x,y,z,k int) positionable.Positionable{
	return BB.mainBoard[x][y].Board[z][k]
}
//func update sector
func (BB *BigBoard) InsertMark(m positionable.Mark,i,j,x,y int) {
	BB.mainBoard[i][j].Board[x][y]=m
}
func (BB *BigBoard)ChangeBoardAvailability(i,j,x,y int){
	//the Board where I put the mark is unavailable now
	for u:=0;u<3;u++{
		for o:=0;o<3;o++{
			BB.AvailableBoards[[2]int{u,o}]=false
		}
	}
	//but one descrived by where I put the mark in the little board is now available
	//even if is the same I just blocked
	if (BB.mainBoard[i][j].IsComplete==false){
		BB.AvailableBoards[[2]int{x, y}] = true
	}else 
	{
		for u:=0;u<3;u++{
			for o:=0;o<3;o++{
				if (BB.mainBoard[u][o].IsComplete==true) {
					BB.AvailableBoards[[2]int{u,o}]=false
				}else{
					BB.AvailableBoards[[2]int{u,o}]=true
				}
			}
		}
		//BB.AvailableBoards[[2]int{x, y}] = false
    
}
}
func (BB BigBoard) Print() {
	fmt.Println(BB.String())
}
func (BB *BigBoard)CheckSmallWin(m positionable.Mark,i,j int) bool{
	little_tris:=&BB.mainBoard[i][j]
	
	for x:=0;x<3;x++{
		mini_tris_done:=true
		for y:=0;y<3;y++{
			if little_tris.Board[x][y]!=m{
				mini_tris_done=false
			}
		}
		if mini_tris_done==true{
		// ho fatto un tris verticale
			BB.mainBoard[i][j].IsComplete=true
			BB.mainBoard[i][j].MarkCompleted=m
			return true
		}
	}
	//if I'm here I didn't do a vertical tris
	for y:=0;y<3;y++{
		mini_tris_done:=true
		for x:=0;x<3;x++{
			if little_tris.Board[x][y]!=m{
				mini_tris_done=false
			}
		}
		if mini_tris_done==true{
		// ho fatto un tris verticale
			BB.mainBoard[i][j].IsComplete=true
			BB.mainBoard[i][j].MarkCompleted=m
			return true
		}
	}
	//if I'm here I didn't do an horizontal tris
	mini_tris_done:=true
	for x:=0;x<3;x++{
		if little_tris.Board[x][x]!=m{
			mini_tris_done=false
		}
	}
	if mini_tris_done==true{
		BB.mainBoard[i][j].IsComplete=true
		BB.mainBoard[i][j].MarkCompleted=m
		return true
	}
	mini_tris_done=true
	for x:=0;x<3;x++{
		if little_tris.Board[2-x][x]!=m{
			mini_tris_done=false
		}
	}
	if mini_tris_done==true{
		BB.mainBoard[i][j].IsComplete=true
		BB.mainBoard[i][j].MarkCompleted=m
		return true
	}
	//at this point I checked the cross section and found nothing
	BB.mainBoard[i][j].IsComplete=false
	return false
}
