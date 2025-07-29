package model
type MarkType int
const(
	circle=iota
	cross
)
func (p posType) String() s String{
	return [...]string{"circle","cross"}
}