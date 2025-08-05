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
	names:=[...]string{"None","MoreCrosses","LessCrosses","SwitchBoards","ChooseRemove","ChoseAdd"}
	if int(c) < 0 || int(c) >= len(names) {
		return "Unknown"
	}
	return names[c]
}