package positionable
import (
	"github.com/google/uuid")
type Positionable interface {
	GetType() int
	String() string
	ImCard()bool
	ImMark() bool
	ImEmpty() bool
	ToInt() int
	Selected(player uuid.UUID, m Mark) 
}

