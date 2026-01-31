package api

import (
	"github.com/LorenzoDOrtona/Tris_Inception/internal/model/board"
)

func BigBoardToDTO(bb *board.BigBoard) GameStateDTO {
	size := bb.SideSize
	boardState := make([][]int, size)
	for i := 0; i < size; i++ {
		boardState[i] = make([]int, size)
		for j := 0; j < size; j++ {
			boardState[i][j] = bb.GetCell(i/3, j/3, i%3, j%3).ToInt()
		}
	}

	return GameStateDTO{
		Board:           boardState,
		AvailableBoards: mapAvailableBoards(bb.AvailableBoards),
		AvailableMoves:  mapAvailableMoves(bb.AvailableMoves),
		IsComplete:      bb.IsComplete,
		Winner:          bb.PlayerWhoCompleted,
		LastTimestamp:   bb.LastTimestamp,
	}
}
func mapAvailableBoards(m map[[2]int]bool) [][]int {
	res := [][]int{}
	for k, v := range m {
		if v {
			res = append(res, []int{k[0], k[1]})
		}
	}
	return res
}

func mapAvailableMoves(m map[[4]int]bool) [][]int {
	res := [][]int{}
	for k, v := range m {
		if v {
			res = append(res, []int{k[0], k[1], k[2], k[3]})
		}
	}
	return res
}
