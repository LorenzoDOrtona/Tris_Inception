package server
import(	
	"github.com/google/uuid"
)
type GameStateDTO struct{
	BoardState [][]int `json:"board_state"` //2D array representing the board
	AvailableBoards [][2]int `json:"available_boards"` //list of available boards
	AvailableMoves [][4]int `json:"available_moves"` //list of available moves
	IsComplete bool `json:"is_complete"` //if the game is complete
	Winner uuid.UUID `json:"winner"` //id of the winning player, null if no winner yet
}