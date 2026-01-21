package server

import (
	"github.com/LorenzoDOrtona/Tris_Inception/cmd/server/api"
	"github.com/google/uuid"
)

type RespToken struct {
	Token string `json:"token"`
}
type RespGameCreatedOrFound struct {
	Token  string    `json:"token"` //authorization token
	IdGame uuid.UUID `json:"id_game"`
	Found  bool      `json:"found"` //true if joined existing game, false if created new game
}
type RespMove struct {
	GameState api.GameStateDTO `json:"game_state"`
}
type RespError struct {
	ErrorMessage string `json:"error_message"`
}
