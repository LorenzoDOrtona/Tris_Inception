package model

type Mark struct {
	marktype uint8
}

func (m Positionable) GetType() (num int) {
	return 0
}
func (c Mark) String() string{
	if GetType()==0{
		return "0"
	}else {
		return "X"
	}
	
}
