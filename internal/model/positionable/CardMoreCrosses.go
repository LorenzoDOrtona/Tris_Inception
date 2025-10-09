package positionable
import (
	"github.com/google/uuid")
type CardMoreCrosses struct{
	Card
}

func (c CardMoreCrosses) String() string{
	var out string
	out="C"
	return out
}
func (c CardMoreCrosses) ImCard() bool{
	return true
}
func (c CardMoreCrosses) ImMark() bool{
	return false
}
func (c CardMoreCrosses) ImEmpty() bool{
	return true
}
func (c CardMoreCrosses) Effect(player uuid.UUID) string{
	return ""
}
func (c CardMoreCrosses) Selected(player uuid.UUID, m Mark) {
	c.Effect(player)
}