package comunication
import (
	"github.com/LorenzoDOrtona/Tris_Inception/internal/model/game"
)
type CommandHandler struct {
	GameEngine Game
	EventBusInstance EventBus
}