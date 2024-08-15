package main

import "github.com/qjuanp/gameoflife/board"

type GameOfLife struct {
	board.Board
}

const ALIVE bool = true
const DEAD bool = false

func NewGameOfSize(rows uint, columns uint) GameOfLife {
	return GameOfLife{board.NewRandomBoard(rows, columns, 1)}
}

func (game *GameOfLife) next() GameOfLife {
	newBoardState := board.NewEmptyBoardAsBigAs(&game.Board)

	for rowIndex, row := range game.Board {
		for columnIndex, cell := range row {
			aliveNeighbors := game.countAliveNeighboards(rowIndex, columnIndex)
			newBoardState[rowIndex][columnIndex] = newCellState(cell, aliveNeighbors)
		}
	}

	return GameOfLife{newBoardState}
}

func (game *GameOfLife) countAliveNeighboards(row int, column int) uint8 {
	var aliveNeighbors uint8 = 0

	for r := game.LowerBound(row); r <= game.RowsUpperBoundFor(row); r++ {
		for c := game.LowerBound(column); c <= game.ColumnsUpperBoundFor(column); c++ {
			// fmt.Printf("on(%d,%d)=%d", r, c, currentBoard[r][c])
			// fmt.Println()
			if (r != row || c != column) && game.Board[r][c] == ALIVE {
				// fmt.Printf("Counted on(%d,%d)=%d", r, c, currentBoard[r][c])
				// fmt.Println()
				aliveNeighbors++
			}
		}
	}

	return aliveNeighbors
}

func newCellState(cell bool, quantityOfAliveNeighbors uint8) bool {
	// Overpopulation
	if quantityOfAliveNeighbors > 3 {
		return DEAD
	}

	// Revive
	if quantityOfAliveNeighbors == 3 {
		return ALIVE
	}

	// Right conditions
	if cell == ALIVE && quantityOfAliveNeighbors >= 2 && quantityOfAliveNeighbors <= 3 {
		return ALIVE
	}

	if quantityOfAliveNeighbors <= 1 {
		return DEAD
	}

	// no rule applied
	return cell
}
