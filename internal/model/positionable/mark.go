package positionable

type Mark struct {
	marktype uint8
}

func (m Mark) GetType() (num int) {
	return 0
}
func (c Mark) String() string{
	if (c.marktype==0){
		return "0"
	}else {
		return "X"
	}
	
}
