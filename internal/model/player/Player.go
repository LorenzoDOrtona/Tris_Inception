package player

import (
	"github.com/google/uuid"
	"github.com/LorenzoDOrtona/Tris_Inception/internal/model/positionable"
)


type Player struct{
	uuid uuid.UUID
	username String
	mark Mark
}