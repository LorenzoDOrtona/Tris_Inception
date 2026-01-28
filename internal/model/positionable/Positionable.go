package positionable

import (
	"math/rand"
)

type Positionable interface {
	GetType() int
	String() string
	ImCard() bool
	ImMark() bool
	ImEmpty() bool
	ToInt() int
	Selected(player string, m Mark)
}

func NewRandomMark() Mark {
	i := rand.Int()
	var k uint8 = 1
	var l uint8 = 2

	if i%2 == 0 {
		return Mark{Marktype: k}
	}
	return Mark{Marktype: l}
}
func OppositeMark(m Mark) Mark {
	if m.Marktype == 1 {
		return Mark{Marktype: 2}
	}
	return Mark{Marktype: 1}
}
