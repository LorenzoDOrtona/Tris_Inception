package internal

import (
	"errors"
	"time"

	"github.com/LorenzoDOrtona/Tris_Inception/cmd/server/api"
	"github.com/LorenzoDOrtona/Tris_Inception/internal/model/game"
	"github.com/google/uuid"
)

type GameController struct {
	ConnectedPlayersID   map[string]uuid.UUID
	ConnectedPlayersName map[string]string
	//games with player waiting for opponents
	ReadyGames   map[string]*game.Game
	OnGoingGames map[string]*game.Game
}

func NewGameController() *GameController {
	return &GameController{
		ConnectedPlayersID:   make(map[string]uuid.UUID),
		ConnectedPlayersName: make(map[string]string),
		ReadyGames:           make(map[string]*game.Game),
		OnGoingGames:         make(map[string]*game.Game),
	}
}
func (GC *GameController) CreatePlayerToken(reqStruct *api.ReqToken) api.RespToken {
	newID, exist := GC.ConnectedPlayersID[reqStruct.PlayerName]
	if exist {
		//exists already the name
		return api.RespToken{
			Token: newID.String(),
		}
	}
	newID = uuid.New()
	GC.ConnectedPlayersID[reqStruct.PlayerName] = newID
	GC.ConnectedPlayersName[newID.String()] = reqStruct.PlayerName
	return api.RespToken{
		Token: newID.String(),
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
		return nil, errors.New("no game")
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
			newGame.MainBoard.LastTimestamp = int(time.Now().Unix())
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
			oppPlayer := game.NewOpponent(reqStruct.Token, name, g.Creator.MarkS)
			//giving opponent player to the game object
			g.Opponent = &oppPlayer
			//update timestamp
			g.MainBoard.LastTimestamp = int(time.Now().Unix())
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
		return nil, errors.New("No player with that Token")
	}
}
func (GC *GameController) CheckLastTimestamp(reqPool *api.ReqPooling) (respPool *api.RespGameState, err error) {
	game, exists := GC.OnGoingGames[reqPool.IdGame]
	if !exists {
		//game not found
		return nil, errors.New("game not found")
	}
	//check timestamp
	if reqPool.Timestamp < game.MainBoard.LastTimestamp {
		//there are updates
		return &api.RespGameState{
			GameState: api.BigBoardToDTO(&game.MainBoard),
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
		return nil, errors.New("game not found")
	}
	// Make the move in the game's main board
	if game.PlayingPlayer.Uuid != reqMove.Token {
		return nil, errors.New("not your turn")
	}
	err := game.CurrentGameState.MoveCommand(reqMove.X, reqMove.Y, reqMove.I, reqMove.J, *game.PlayingPlayer)
	if err != nil {
		return nil, err
	}
	game.MainBoard.LastTimestamp = int(time.Now().Unix())
	return &api.RespGameState{
		GameState: api.BigBoardToDTO(&game.MainBoard),
	}, nil
}
