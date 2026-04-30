package api

import (
	"github.com/LorenzoDOrtona/Tris_Inception/internal/model/game"
)

func BigBoardToDTO(g *game.Game) GameStateDTO {
	size := g.MainBoard.SideSize
	boardState := make([][]int, size)
	for i := 0; i < size; i++ {
		boardState[i] = make([]int, size)
		for j := 0; j < size; j++ {
			boardState[i][j] = g.MainBoard.GetCell(i/3, j/3, i%3, j%3).ToInt()
		}
	}
	S := g.Opponent != nil
	winnerName := ""
	if g.Winner != nil {
		winnerName = g.Winner.Username
	}
	return GameStateDTO{
		Started:         S,
		Board:           boardState,
		AvailableBoards: mapAvailableBoards(g.MainBoard.AvailableBoards),
		AvailableMoves:  mapAvailableMoves(g.MainBoard.AvailableMoves),
		IsComplete:      g.MainBoard.IsComplete,
		Winner:          winnerName,
		LastTimestamp:   g.MainBoard.LastTimestamp,
	}
}

// Transforms the map into a 3x3 matrix where 1 = Active, 0 = Inactive
func mapAvailableBoards(m map[[2]int]bool) [][]int {
	// Create empty 3x3 matrix
	matrix := make([][]int, 3)
	for i := range matrix {
		matrix[i] = make([]int, 3)
	}

	// Populate with 1 where the board is available
	for k, v := range m {
		if v {
			matrix[k[0]][k[1]] = 1
		}
	}
	return matrix
}

// Note: AvailableMoves is a list of specific moves, it's fine to leave it as a list
// because the frontend doesn't use it to draw the grid, but for validation (if implemented).
// If you just want to draw, it's fine for now.
func mapAvailableMoves(m map[[4]int]bool) [][]int {
	res := [][]int{}
	for k, v := range m {
		if v {
			res = append(res, []int{k[0], k[1], k[2], k[3]})
		}
	}
	return res
}
