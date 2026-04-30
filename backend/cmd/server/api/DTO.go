package api

type GameStateDTO struct {
	// Add backticks `json:"..."` to tell Go how to name fields in JSON
	Started         bool    `json:"started"`
	Board           [][]int `json:"board"`
	AvailableBoards [][]int `json:"available_boards"`
	AvailableMoves  [][]int `json:"available_moves"`
	IsComplete      bool    `json:"is_complete"`
	Winner          string  `json:"winner"`
	LastTimestamp   int     `json:"last_timestamp"` // Important: lowercase for consistency with JS
}
