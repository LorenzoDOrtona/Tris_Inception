package main
import(	
	"encoding/json"
	"fmt"
	"log"
)
type req_base struct{
	Type string
}
type req_joinGame struct{
	Token string
	User_Type string //real or bot
	Req_Type string
	//--	
	Old_Game bool
	ID_Game UUID.uuid	
	//--
	Game_Mode string //no cards, yes cards
	Side_Measure int //9 if 9x9 (3x 3x3)
	Wanted_Enemy UUID.uuid //or None


}
type req_move struct{
	Token string
	User_Type string
	Req_Type string
	//--	
	ID_Game UUID.uuid
	//--	
	x int
	y int
	i int
	j int

}
