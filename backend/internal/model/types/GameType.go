package types
type GameType int
const(
	NoGame=iota
	offlineVsComputer
	offlineVsFriend
	onlineVsComputer
	onlineVsFriend
)
func (c GameType)String() string{
	names:=[...]string{"None",
						"offlineVsComputer",
						"offlineVsFriend",
						"onlineVsComputer",
						"onlineVsFriend"}
	if int(c) < 0 || int(c) >= len(names) {
		return "Unknown"
	}
	return names[c]
}