package api

type GameStateDTO struct {
	Board [][]int `json:"board"`

	AvailableBoards [][]int `json:"available_boards"`
	AvailableMoves  [][]int `json:"available_moves"`

	IsComplete    bool   `json:"is_complete"`
	Winner        string `json:"winner,omitempty"`
	LastTimestamp int    `json:"LastTimestamp"`
}
