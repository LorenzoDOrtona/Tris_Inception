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

	return GameStateDTO{
		Started:         S,
		Board:           boardState,
		AvailableBoards: mapAvailableBoards(g.MainBoard.AvailableBoards),
		AvailableMoves:  mapAvailableMoves(g.MainBoard.AvailableMoves),
		IsComplete:      g.MainBoard.IsComplete,
		Winner:          g.MainBoard.PlayerWhoCompleted,
		LastTimestamp:   g.MainBoard.LastTimestamp,
	}
}

// Trasforma la mappa in una matrice 3x3 dove 1 = Attivo, 0 = Inattivo
func mapAvailableBoards(m map[[2]int]bool) [][]int {
	// Crea matrice 3x3 vuota
	matrix := make([][]int, 3)
	for i := range matrix {
		matrix[i] = make([]int, 3)
	}

	// Popola con 1 dove la board è disponibile
	for k, v := range m {
		if v {
			matrix[k[0]][k[1]] = 1
		}
	}
	return matrix
}

// Nota: AvailableMoves è una lista di mosse specifiche, questa va bene lasciarla come lista
// perché il frontend non la usa per disegnare la griglia, ma per validare (se lo implementerai).
// Se però vuoi solo disegnare, per ora va bene anche così.
func mapAvailableMoves(m map[[4]int]bool) [][]int {
	res := [][]int{}
	for k, v := range m {
		if v {
			res = append(res, []int{k[0], k[1], k[2], k[3]})
		}
	}
	return res
}
