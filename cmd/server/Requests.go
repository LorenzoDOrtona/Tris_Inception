package main
import(	
	"encoding/json"
	"fmt"
	"log"
)
type req_base struct{
	Type string `json:"type"`
}
type req_joinGame struct{
	Token string `json:"token"`
	User_Type string `json:"user_type"`//real or bot
	Req_Type string `json:"req_type"`
	//--	
	Old_Game bool  `json:"type"`
	ID_Game UUID.uuid `json:"type"`
	//--
	Game_Mode string `json:"type"` //no cards, yes cards
	
	Side_Measure int `json:"type"`  //9 if 9x9 (3x 3x3)
	Wanted_Enemy UUID.uuid `json:"type"` //or None

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
