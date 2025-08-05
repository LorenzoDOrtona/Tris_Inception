package types
type CardType int
const(
	None=iota
	MoreCrosses
	LessCrosses
	SwitchBoards
	ChooseRemove
	ChoseAdd
)
func (c CardType)String() string{
	return[...]string{"None","MoreCrosses","LessCrosses","SwitchBoards","ChooseRemove","ChoseAdd"}
}