package types

type MarkType int

const (
	Circle = iota
	Cross
)

func (p MarkType) String() string {
	return [...]string{"Circle", "Cross"}
}
