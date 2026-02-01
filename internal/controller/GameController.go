package internal

import (
	"time"

	"github.com/LorenzoDOrtona/Tris_Inception/cmd/server/api"
	"github.com/LorenzoDOrtona/Tris_Inception/internal/errors"
	"github.com/LorenzoDOrtona/Tris_Inception/internal/model/game"
	"github.com/google/uuid"
)

type GameController struct {
	ConnectedPlayersID   map[string]string
	ConnectedPlayersName map[string]string
	//games with player waiting for opponents
	ReadyGames   map[string]*game.Game
	OnGoingGames map[string]*game.Game
}

func NewGameController() *GameController {
	return &GameController{
		ConnectedPlayersID:   make(map[string]string),
		ConnectedPlayersName: make(map[string]string),
		ReadyGames:           make(map[string]*game.Game),
		OnGoingGames:         make(map[string]*game.Game),
	}
}
func (GC *GameController) CreatePlayerToken(reqStruct *api.ReqToken) api.RespToken {
	Id, exist := GC.ConnectedPlayersID[reqStruct.PlayerName]
	if exist {
		//exists already the name
		return api.RespToken{
			Token: Id,
		}
	} else {

		newID := uuid.New().String()
		GC.ConnectedPlayersID[reqStruct.PlayerName] = newID
		GC.ConnectedPlayersName[newID] = reqStruct.PlayerName
		return api.RespToken{
			Token: newID,
		}
	}
}
func FeasibleGameFound(g game.Game, reqStruct *api.ReqCreateOrJoinGame) bool {
	return true
}
func (GC *GameController) checkForExistingGame(reqStruct *api.ReqCreateOrJoinGame) (*game.Game, error) {
	var gId string
	for id, game := range GC.ReadyGames {
		if FeasibleGameFound(*game, reqStruct) {
			gId = id
			break
		}
	}
	//didnt find a good game
	if gId == "" {
		return nil, errors.ErrGameNotFound
	}
	//found one
	return GC.ReadyGames[gId], nil
}
func (GC *GameController) CreateGame(reqStruct *api.ReqCreateOrJoinGame) (*api.RespGameCreatedOrFound, error) {
	//Check if a user with that Token exists
	name, exist := GC.ConnectedPlayersName[reqStruct.Token]
	if exist {
		g, err := GC.checkForExistingGame(reqStruct)
		if err != nil {
			//no game matching expectation found
			//creating one new
			newCreator := game.NewPlayer(reqStruct.Token, name)
			newGame := game.New(&newCreator, nil)
			GC.ReadyGames[newGame.GameUuid.String()] = newGame
			return &api.RespGameCreatedOrFound{
				IdGame:      newGame.GameUuid,
				CreatorName: newCreator.Username,
				Found:       false,
			}, nil
		} else {
			//joining existing game
			gameUUID := g.GameUuid
			creatorName := g.Creator.Username
			//creating oppenent player
			g.CurrentGameState.AddOpponent(reqStruct.Token, name)
			//removing from ready games
			delete(GC.ReadyGames, g.GameUuid.String())
			//putting in Ongoing games
			GC.OnGoingGames[g.GameUuid.String()] = g
			return &api.RespGameCreatedOrFound{
				IdGame:      gameUUID,
				CreatorName: creatorName,
				Found:       true,
			}, nil

		}
	} else {
		return nil, errors.ErrPlayerNotFound
	}
}
func (GC *GameController) CheckLastTimestamp(reqPool *api.ReqPooling) (respPool *api.RespGameState, err error) {
	game, exists := GC.OnGoingGames[reqPool.IdGame]
	if !exists {
		//game not found
		game, exists = GC.ReadyGames[reqPool.IdGame]
		if !exists {
			return nil, errors.ErrPlayerNotFound
		}
	}
	//check timestamp
	if reqPool.Timestamp < game.MainBoard.LastTimestamp {
		//there asre updates
		return &api.RespGameState{
			GameState: api.BigBoardToDTO(game),
		}, nil
	}
	//no updates
	return &api.RespGameState{
		GameState: api.GameStateDTO{},
	}, nil
}
func (GC *GameController) MakeMove(reqMove *api.ReqMove) (*api.RespGameState, error) {
	game, exists := GC.OnGoingGames[reqMove.IdGame]
	if !exists {
		return nil, errors.ErrGameNotFound
	}
	// Make the move in the game's main board
	if game.PlayingPlayer.Uuid != reqMove.Token {
		return nil, errors.ErrNotYourTurn
	}
	err := game.CurrentGameState.MoveCommand(reqMove.X, reqMove.Y, reqMove.I, reqMove.J, *game.PlayingPlayer)
	if err != nil {
		return nil, errors.ErrInvalidMove
	}
	game.MainBoard.LastTimestamp = int(time.Now().Unix())
	return &api.RespGameState{
		GameState: api.BigBoardToDTO(game),
	}, nil
}
