package api

type GameStateDTO struct {
	// Aggiungi i backtick `json:"..."` per dire a Go come chiamare i campi nel JSON
	Started         bool    `json:"started"`
	Board           [][]int `json:"board"`
	AvailableBoards [][]int `json:"available_boards"`
	AvailableMoves  [][]int `json:"available_moves"`
	IsComplete      bool    `json:"is_complete"`
	Winner          string  `json:"winner"`
	LastTimestamp   int     `json:"last_timestamp"` // Importante: minuscolo per coerenza col JS
}
