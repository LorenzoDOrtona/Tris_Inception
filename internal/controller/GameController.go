package internal

import (
	"sync"
	"time"

	"github.com/LorenzoDOrtona/Tris_Inception/cmd/server/api"
	"github.com/LorenzoDOrtona/Tris_Inception/internal/errors"
	"github.com/LorenzoDOrtona/Tris_Inception/internal/model/game"
	"github.com/google/uuid"
)

type GameController struct {
	muPlayer             sync.Mutex
	ConnectedPlayersID   map[string]string
	ConnectedPlayersName map[string]string
	//games with player waiting for opponents
	muGames      sync.Mutex
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
	GC.muPlayer.Lock()
	Id, exist := GC.ConnectedPlayersID[reqStruct.PlayerName]
	if exist {
		//exists already the name
		GC.muPlayer.Unlock()
		return api.RespToken{
			Token: Id,
		}
	} else {
		newID := uuid.New().String()
		GC.ConnectedPlayersID[reqStruct.PlayerName] = newID
		GC.ConnectedPlayersName[newID] = reqStruct.PlayerName
		GC.muPlayer.Unlock()
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
	GC.muGames.Lock()
	for id, game := range GC.ReadyGames {
		if FeasibleGameFound(*game, reqStruct) {
			gId = id
			break
		}
	}
	//didnt find a good game
	if gId == "" {
		GC.muGames.Unlock()
		return nil, errors.ErrGameNotFound
	}
	//found one
	game := GC.ReadyGames[gId]
	GC.muGames.Unlock()
	return game, nil
}
func (GC *GameController) StartABot(g *game.Game) {
	GC.muGames.Lock()
	GC.OnGoingGames[g.GameUuid.String()] = g
	GC.muGames.Unlock()
	for {
		time.Sleep(500 * time.Millisecond)
		g.Mu.Lock()
		if g.Finished {
			g.Mu.Unlock()
			break
		}
		if g.PlayingPlayer.Uuid == "BOT-UUID" {
			move := g.MainBoard.MakeBotMove("")
			g.CurrentGameState.MoveCommand(move[0], move[1], move[2], move[3], *g.PlayingPlayer)
		}
		g.Mu.Unlock()

	}
	GC.muGames.Lock()
	delete(GC.OnGoingGames, g.GameUuid.String())
	GC.muGames.Unlock()
}
func (GC *GameController) CreateGame(reqStruct *api.ReqCreateOrJoinGame) (*api.RespGameCreatedOrFound, error) {
	//Check if a user( with that Token exists
	GC.muPlayer.Lock()
	name, exist := GC.ConnectedPlayersName[reqStruct.Token]
	GC.muPlayer.Unlock()
	//if player exits
	if exist {
		//bot request
		if reqStruct.GameMode == "BOT" {
			newCreator := game.NewPlayer(reqStruct.Token, name)
			newGame := game.New(&newCreator, nil)
			gameUUID := newGame.GameUuid
			creatorName := newGame.Creator.Username
			//creating oppenent player
			newGame.CurrentGameState.AddOpponent("BOT-UUID", "BOT", false)
			go GC.StartABot(newGame)
			return &api.RespGameCreatedOrFound{
				IdGame:      gameUUID,
				CreatorName: creatorName,
				Found:       true,
			}, nil

		} else {
			//real player request
			g, err := GC.checkForExistingGame(reqStruct)
			if err != nil {
				//no game matching expectation found
				//creating one new
				newCreator := game.NewPlayer(reqStruct.Token, name)
				newGame := game.New(&newCreator, nil)

				GC.muGames.Lock()
				GC.ReadyGames[newGame.GameUuid.String()] = newGame
				GC.muGames.Unlock()

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
				g.CurrentGameState.AddOpponent(reqStruct.Token, name, false)
				//removing from ready games
				GC.muGames.Lock()
				delete(GC.ReadyGames, g.GameUuid.String())
				//putting in Ongoing games
				GC.OnGoingGames[g.GameUuid.String()] = g
				GC.muGames.Unlock()
				return &api.RespGameCreatedOrFound{
					IdGame:      gameUUID,
					CreatorName: creatorName,
					Found:       true,
				}, nil

			}
		}
	} else {
		return nil, errors.ErrPlayerNotFound
	}
}
func (GC *GameController) CheckLastTimestamp(reqPool *api.ReqPooling) (respPool *api.RespGameState, err error) {
	GC.muGames.Lock()
	game, exists := GC.OnGoingGames[reqPool.IdGame]
	GC.muGames.Unlock()
	if !exists {
		//game not found
		GC.muGames.Lock()
		game, exists = GC.ReadyGames[reqPool.IdGame]
		GC.muGames.Unlock()
		if !exists {
			return nil, errors.ErrPlayerNotFound
		}
	}
	game.Mu.Lock()
	defer game.Mu.Unlock() //unlock when I exit from the function
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
	GC.muGames.Lock()
	game, exists := GC.OnGoingGames[reqMove.IdGame]
	GC.muGames.Unlock()
	if !exists {
		return nil, errors.ErrGameNotFound
	}
	game.Mu.Lock()
	defer game.Mu.Unlock() //unlock when I exit from the function
	// Make the move in the game's main board
	if game.PlayingPlayer.Uuid != reqMove.Token {
		return nil, errors.ErrNotYourTurn
	}
	err := game.CurrentGameState.MoveCommand(reqMove.X, reqMove.Y, reqMove.I, reqMove.J, *game.PlayingPlayer)
	if err != nil {
		return nil, errors.ErrInvalidMove
	}
	game.MainBoard.LastTimestamp = int(time.Now().UnixNano())
	return &api.RespGameState{
		GameState: api.BigBoardToDTO(game),
	}, nil
}
