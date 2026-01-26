package internal

import (
	"github.com/LorenzoDOrtona/Tris_Inception/cmd/server/api"
	"github.com/LorenzoDOrtona/Tris_Inception/internal/model/game"
	"github.com/google/uuid"
)

type GameController struct {
	ConnectedPlayers  map[game.Player]uuid.UUID
	GameIdToGameState map[uuid.UUID]game.Game
	//games with player waiting for opponents
	ReadyGames   []game.Game
	OnGoingGames []game.Game
}

func (GC *GameController) CreateGame(reqStruct *api.ReqCreateOrJoinGame) {

}
