package positionable

type CardMoreCrosses struct {
	Card
	Positionable
}

func (c CardMoreCrosses) String() string {
	var out string
	out = "C"
	return out
}
func (c CardMoreCrosses) ImCard() bool {
	return true
}
func (c CardMoreCrosses) ImMark() bool {
	return false
}
func (c CardMoreCrosses) ImEmpty() bool {
	return true
}
func (c CardMoreCrosses) Effect(player string) string {
	return ""
}
func (c CardMoreCrosses) Selected(player string, m Mark) {
	c.Effect(player)
}
