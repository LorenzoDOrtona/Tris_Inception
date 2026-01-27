package internal

import (
	"github.com/LorenzoDOrtona/Tris_Inception/cmd/server/api"
	"github.com/LorenzoDOrtona/Tris_Inception/internal/model/game"
	"github.com/google/uuid"
)

type GameController struct {
	ConnectedPlayers  map[string]uuid.UUID
	GameIdToGameState map[uuid.UUID]game.Game
	//games with player waiting for opponents
	ReadyGames   []game.Game
	OnGoingGames []game.Game
}

func NewGameController() *GameController {
	return &GameController{
		ConnectedPlayers:  make(map[string]uuid.UUID),
		GameIdToGameState: make(map[uuid.UUID]game.Game),
		ReadyGames:        make([]game.Game, 0),
		OnGoingGames:      make([]game.Game, 0),
	}
}
func (GC *GameController) CreatePlayerToken(reqStruct *api.ReqToken) api.RespToken {
	newID := uuid.New()
	return api.RespToken{
		Token: newID.String(),
	}
}
func (GC *GameController) CreateGame(reqStruct *api.ReqCreateOrJoinGame) {

}
