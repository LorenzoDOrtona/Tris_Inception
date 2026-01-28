package api

import (
	"github.com/google/uuid"
)

type RespToken struct {
	Token string `json:"token"`
}
type RespGameCreatedOrFound struct {
	IdGame      uuid.UUID `json:"id_game"`
	CreatorName string    `json:"creator_name"`
	//iff found there is a creator name
	Found bool `json:"found"` //true if joined existing game, false if created new game
}
type RespMove struct {
	GameState GameStateDTO `json:"game_state"`
}
type RespError struct {
	ErrorMessage string `json:"error_message"`
}
