package server
import(	
	"github.com/google/uuid"
)
type resp_Token struct{
	Token string `json:"token"`
}
type resp_GameCreatedOrFound struct{
	Token string `json:"token"` //authorization token
	IdGame uuid.UUID `json:"id_game"`
	Found bool `json:"found"` //true if joined existing game, false if created new game
}
type resp_Move struct{
	BoardState [][]int `json:"board_state"` //2D array representing the board
}
type resp_Error struct{
	ErrorMessage string `json:"error_message"`
}