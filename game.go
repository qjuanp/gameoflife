package main

type GameOfLife struct {
	Board
}

func NewGameOfSize(rows uint, columns uint) GameOfLife {
	return GameOfLife{NewRandomBoard(rows, columns, 1)}
}

func (game *GameOfLife) next() GameOfLife {
	newBoardState := NewEmptyBoardAsBigAs(&game.Board)

	for rowIndex, row := range game.Board {
		for columnIndex, cell := range row {
			aliveNeighbors := game.countAliveNeighboards(rowIndex, columnIndex)
			newBoardState[rowIndex][columnIndex] = cell.newCellState(aliveNeighbors)
		}
	}

	return GameOfLife{newBoardState}
}

func (game *GameOfLife) countAliveNeighboards(row int, column int) uint8 {
	aliveNeighbors := 0

	for r := game.lowerBound(row); r <= game.upperBound(row, game.numberOfRows()); r++ {
		for c := game.lowerBound(column); c <= game.upperBound(column, game.numberOfColumns()); c++ {
			// fmt.Printf("on(%d,%d)=%d", r, c, currentBoard[r][c])
			// fmt.Println()
			if (r != row || c != column) && game.Board[r][c] == ALIVE {
				// fmt.Printf("Counted on(%d,%d)=%d", r, c, currentBoard[r][c])
				// fmt.Println()
				aliveNeighbors++
			}
		}
	}

	return uint8(aliveNeighbors)
}
