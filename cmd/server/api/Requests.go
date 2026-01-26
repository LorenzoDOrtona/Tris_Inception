package api

import (
	"github.com/google/uuid"
)

type ReqToken struct {
	PlayerName string `json:"player_name"`
	UserType   string `json:"user_type"` //real or bot
}

type ReqCreateOrJoinGame struct {
	Token       string `json:"token"`        //authorization token
	GameMode    string `json:"game_mode"`    //no cards, yes cards
	SideMeasure int    `json:"side_measure"` //9 if 9x9 (3x 3x3)
}
type ReqJoinGame struct {
	Token  string    `json:"token"`
	IdGame uuid.UUID `json:"id_game"`
}

type ReqMove struct {
	Token string `json:"token"`
	//--
	IdGame uuid.UUID `json:"id_game"`
	//--
	X int `json:"x"`
	Y int `json:"y"`
	I int `json:"i"`
	J int `json:"j"`
}
