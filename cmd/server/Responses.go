package server
import(	
	"github.com/google/uuid"
)
type resp_GameCreated struct{
	Token string `json:"token"` //authorization token
	ID_Game uuid.UUID `json:"id_game"`
}
type resp_Error struct{
	ErrorMessage string `json:"error_message"`
}