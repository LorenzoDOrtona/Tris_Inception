package errors

import "errors"

var (
	ErrGameNotFound   = errors.New("game not found")
	ErrNotYourTurn    = errors.New("not your turn")
	ErrInvalidMove    = errors.New("invalid move")
	ErrPlayerNotFound = errors.New("Player not found")
)
