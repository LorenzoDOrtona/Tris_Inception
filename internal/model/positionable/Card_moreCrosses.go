package model
import(
	"fmt"
)
type Card_moreCrosses struct{
	Card
}
func effect(){

}
func (c Positionable) GetType() (num int) {
	return 0
}
func (c Card_moreCrosses) String() string{
	var out string
	out="C"
	return out
}