package positionable
import (
	"github.com/google/uuid")
type Card interface {
	Positionable
	Effect(player uuid.UUID) string//EffResponse
}



