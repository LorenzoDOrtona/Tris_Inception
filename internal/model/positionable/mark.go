package positionable

type Mark struct {
	Marktype uint8
	Positionable
}

func (m Mark) GetType() (num int) {
	return 0
}
func (c Mark) String() string{
	if (c.Marktype==0){
		return "_"
	}
	if (c.Marktype==1){
		return "X"
	}
	return "O"
	
	
}
func (m Mark) ImCard() bool{
	return false
}
func (m Mark)ImMark() bool{
	return true
}
func (m Mark) ImEmpty() bool{
	if (m.Marktype==0){
		return true
	}else {
		return false
	}
}
