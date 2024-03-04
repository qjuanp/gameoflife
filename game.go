package main

import "github.com/qjuanp/gameoflife/board"

type GameOfLife struct {
	board.Board
}

func NewGameOfSize(rows uint, columns uint) GameOfLife {
	return GameOfLife{board.NewRandomBoard(rows, columns, 1)}
}

func (game *GameOfLife) next() GameOfLife {
	newBoardState := board.NewEmptyBoardAsBigAs(&game.Board)

	for rowIndex, row := range game.Board {
		for columnIndex, cell := range row {
			aliveNeighbors := game.countAliveNeighboards(rowIndex, columnIndex)
			newBoardState[rowIndex][columnIndex] = cell.NewCellState(aliveNeighbors)
		}
	}

	return GameOfLife{newBoardState}
}

func (game *GameOfLife) countAliveNeighboards(row int, column int) uint8 {
	aliveNeighbors := 0

	for r := game.LowerBound(row); r <= game.UpperBound(row, game.NumberOfRows()); r++ {
		for c := game.LowerBound(column); c <= game.UpperBound(column, game.NumberOfColumns()); c++ {
			// fmt.Printf("on(%d,%d)=%d", r, c, currentBoard[r][c])
			// fmt.Println()
			if (r != row || c != column) && game.Board[r][c] == board.ALIVE {
				// fmt.Printf("Counted on(%d,%d)=%d", r, c, currentBoard[r][c])
				// fmt.Println()
				aliveNeighbors++
			}
		}
	}

	return uint8(aliveNeighbors)
}
