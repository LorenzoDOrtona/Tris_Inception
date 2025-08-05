package types

type MarkType int

const (
	Circle = iota
	Cross
)

func (p MarkType) String() string {
	names := [...]string{"Circle", "Cross"}
	if int(p) < 0 || int(p) >= len(names) {
		return "Unknown"
	}
	return names[p]
}
