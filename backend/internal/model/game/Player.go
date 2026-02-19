package game

import (
	"github.com/LorenzoDOrtona/Tris_Inception/internal/model/positionable"
)

type Player struct {
	Uuid     string
	Username string
	MarkS    positionable.Mark
}

func NewPlayer(uuid string, name string) Player {
	return Player{
		Uuid:     uuid,
		Username: name,
		MarkS:    positionable.NewRandomMark(),
	}
}
