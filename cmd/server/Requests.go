package server
import(	
	"github.com/google/uuid"
)
type req_Token struct{
	PlayerName string `json:"player_name"`
	UserType string `json:"user_type"`//real or bot
}

type req_CreateOrJoinGame struct{
	Token string `json:"token"` //authorization token
	GameMode string `json:"game_mode"` //no cards, yes cards
	SideMeasure int `json:"side_measure"`  //9 if 9x9 (3x 3x3)
}
type req_joinGame struct{
    Token string `json:"token"`
	IdGame uuid.UUID `json:"id_game"`
}

type req_move struct{
	Token string `json:"token"`
	//--	
	IdGame uuid.UUID `json:"id_game"`
	//--	
	x int `json:"x"`
	y int `json:"y"`
	i int `json:"i"`
	j int `json:"j"`

}
