package board

import (
	"fmt"
	"github.com/LorenzoDOrtona/Tris_Inception/internal/model/positionable"
	"github.com/google/uuid"
)

type BigBoard struct {
	mainBoard [3][3]Board
	SideSize int
	AvailableBoards map[[2]int]bool
	AvailableMoves map[[4]int]bool
	IsComplete bool
	PlayerWhoCompleted uuid.UUID
}

func (B *BigBoard) SetupBigBoard() {
	//initialize the map
	B.SideSize=9
	//at the beginning the bigBoard is obv not complete
	B.IsComplete=false
	//Setting the boards availability
	B.AvailableBoards=make(map[[2]int]bool)
	B.AvailableMoves=make(map[[4]int]bool)
	//Setupping available moves, all are available
	for i := 0; i < B.SideSize/3; i++ {
		for j := 0; j < B.SideSize/3; j++ {
			for x:= 0; x < B.SideSize/3; x++ {
				for y := 0; y < B.SideSize/3; y++ {
					B.AvailableMoves[[4]int{i,j,x,y}]=true
				}}}}
	for i := 0; i < B.SideSize/3; i++ {
		for j := 0; j < B.SideSize/3; j++ {
			B.mainBoard[i][j].SetupBoard()
			//setting the board accessible
			B.AvailableBoards[[2]int{i,j}]=true
		}
	}


}
func (BB BigBoard) String() string {
	var out string
	for i := 0; i < BB.SideSize; i++ {
		if i%3 == 0 && i/3>0 {
			out+="---------------------\n"
		}
		for j := 0; j < BB.SideSize; j++ {
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
func (BB *BigBoard)UpdateAvailableMoves(){
	
}
func (BB *BigBoard)MakeAllMovesUnavailable(){
	for k,val := BB.AvailableMoves{
		BB.AvailableMoves[k]=false
	}
	
}
func (BB *BigBoard)AllowNewCorrectMoves(i,j,x,y){
	//if x, y are the cordinates of a complete little board
	//Now there are available moves across all the bigBoard
	if(BB.mainBoard[x][y].IsComplete){
		for k,v:=BB.AvailableBoards{
			for u:=0;u<3;u++{
				for o:=0;o<3;o++{
					if BB.mainBoard[k[0]][k[1]].Board[u][o].ImEmpty{
						BB.AvailableMoves[[4]int{k[0],k[1],u,o}]=true
					}else{
						BB.AvailableMoves[[4]int{k[0],k[1],u,o}]=false
					}
				}}
		}

	}else{
		for u:=0;u<3;u++{
			for o:=0;o<3;o++{
				if BB.mainBoard[x][y].Board[u][o].ImEmpty{
					BB.AvailableMoves[[4]int{x,y,u,o}]=true
				}else{
					BB.AvailableMoves[[4]int{x,y,u,o}]=false
				}
			}}
	}
	
	
}
func (BB *BigBoard)ChangeBoardAvailability(i,j,x,y int){
	//Im making every board unavailable now
	for u:=0;u<3;u++{
		for o:=0;o<3;o++{
			BB.AvailableBoards[[2]int{u,o}]=false
		}
	}
	
	//but one (descrived by where I put the mark in the little board) is now available
	//if its not complete
	if (BB.mainBoard[x][y].IsComplete==false){
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
	}
	//And also every move is now unavailable
	BB.MakeAllMovesUnavailable()
	//and only correct moves are available
	BB.AllowNewCorrectMoves(i,j,x,y)
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
func (BB *BigBoard)CheckWin() bool {
	return false
}
func (BB *BigBoard)CheckWin(m positionable.Positionable,player uuid.UUID) bool {
	
	
	for x:=0;x<3;x++{
		mini_tris_done:=true
		for y:=0;y<3;y++{
			little_tris:=&BB.mainBoard[x][y]
			if little_tris.IsComplete && little_tris.MarkCompleted!=m{
				mini_tris_done=false
			}
		}
		if mini_tris_done==true{
		// ho fatto un tris verticale
			BB.IsComplete=true
			BB.PlayerWhoCompleted=player
			return true
		}
	}
	//if I'm here I didn't do a vertical tris
	for y:=0;y<3;y++{
		mini_tris_done:=true
		for x:=0;x<3;x++{
			little_tris:=&BB.mainBoard[x][y]
			if little_tris.IsComplete && little_tris.MarkCompleted!=m{
				mini_tris_done=false
			}
		}
		if mini_tris_done==true{
		// ho fatto un tris verticale
			BB.IsComplete=true
			BB.PlayerWhoCompleted=player
			return true
		}
	}
	//if I'm here I didn't do an horizontal tris
	mini_tris_done:=true
	for x:=0;x<3;x++{
		little_tris:=&BB.mainBoard[x][x]
		if little_tris.IsComplete && little_tris.MarkCompleted!=m{
			mini_tris_done=false
		}
	}
	if mini_tris_done==true{
		BB.IsComplete=true
		BB.PlayerWhoCompleted=player
		return true
	}
	mini_tris_done=true
	for x:=0;x<3;x++{
		little_tris:=&BB.mainBoard[2-x][x]
		if little_tris.IsComplete && little_tris.MarkCompleted!=m{
			mini_tris_done=false
		}
		
	}
	if mini_tris_done==true{
		BB.IsComplete=true
		BB.PlayerWhoCompleted=player
		return true
	}
	//at this point I checked the cross section and found nothing
	BB.IsComplete=false
	return false
}