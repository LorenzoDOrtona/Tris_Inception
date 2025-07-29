package model

type Card interface {
	Positionable
}

func effect() {

}
func (c Card) String() string{
	var out string
	out="C"
	return out
}
