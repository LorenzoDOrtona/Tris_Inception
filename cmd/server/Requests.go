package server
import(	
	"github.com/google/uuid"
)
type req_CreateGame struct{
	UserType string `json:"user_type"`//real or bot
	PlayerName string `json:"player_name"`
	GameMode string `json:"game_mode"` //no cards, yes cards
	SideMeasure int `json:"side_measure"`  //9 if 9x9 (3x 3x3)
}
type req_joinGame struct{
    Token string `json:"token"`
	UserType string `json:"user_type"`//real or bot
	GameMode string `json:"game_mode"` //no cards, yes cards
	SideMeasure int `json:"side_measure"`  //9 if 9x9 (3x 3x3)
}

type req_joinOldGame struct{
	Token string `json:"token"`
	UserType string `json:"user_type"`//real or bot
	ID_Game uuid.UUID `json:"id_game"`
}

type req_joinFriendGame struct{
    Token string `json:"token"`
	UserType string `json:"user_type"`//real or bot
	GameMode string `json:"game_mode"` //no cards, yes cards
	SideMeasure int `json:"side_measure"`  //9 if 9x9 (3x 3x3)
	WantedEnemy uuid.UUID `json:"wanted_enemy"` //or None
}
type req_move struct{
	Token string `json:"token"`
	UserType string `json:"user_type"`//real or bot
	//--	
	ID_Game uuid.UUID `json:"id_game"`
	//--	
	x int `json:"x"`
	y int `json:"y"`
	i int `json:"i"`
	j int `json:"j"`

}
