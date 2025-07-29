package model
type CardType int
const(
	None=iota
	moreCrosses
	lessCrosses
	switchBoards
	chooseRemove
	choseAdd
)
func (c CardType)String()string{
	return[...]string{"None","moreCrosses","lessCrosses","switchBoards","chooseRemove","choseAdd"}
}