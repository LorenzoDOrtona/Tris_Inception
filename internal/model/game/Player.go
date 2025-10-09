package game

import (
	"github.com/google/uuid"
	"github.com/LorenzoDOrtona/Tris_Inception/internal/model/positionable"
)


type Player struct{
	Uuid uuid.UUID
	Username string
	MarkS positionable.Mark
}