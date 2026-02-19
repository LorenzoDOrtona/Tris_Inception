package positionable

type Mark struct {
	Marktype uint8
	Positionable
}

func (m Mark) GetType() (num int) {
	return 0
}
func (c Mark) String() string {
	if c.Marktype == 0 {
		return "_"
	}
	if c.Marktype == 1 {
		return "X"
	}
	return "O"

}
func (c Mark) ToInt() int {
	if c.Marktype == 0 { //empty
		return 0
	}
	if c.Marktype == 1 { //X
		return 1
	}
	return 2 //O

}
func (m Mark) ImCard() bool {
	return false
}
func (m Mark) ImMark() bool {
	return true
}
func (m Mark) ImEmpty() bool {
	if m.Marktype == 0 {
		return true
	} else {
		return false
	}
}
