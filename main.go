package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	fmt.Println("Project setup done")
	board := ramdomInitialization(5, 5, 1)
	fmt.Print(boardToString(board, toInt))
}

type serializer func(uint8) string

const ALIVE uint8 = 1
const ALIVE_CHARACTER string = "\u2588"

const DEAD uint8 = 0
const DEAD_CHARACTER string = "\u0020"

func createBoard(rows int, columns int) [][]uint8 {
	var board [][]uint8 = make([][]uint8, rows)

	for row := range board {
		board[row] = make([]uint8, columns)
	}

	return board
}

func ramdomInitialization(rows uint32, columns uint32, seed int64) [][]uint8 {
	var board [][]uint8 = make([][]uint8, rows)
	rand.Seed(seed)

	for row := range board {
		board[row] = make([]uint8, columns)
		for column := range board[row] {
			randomValue := rand.Float64()
			if randomValue >= 0.5 {
				board[row][column] = ALIVE
			} else {
				board[row][column] = DEAD
			}
		}
	}

	return board
}

func boardToString(board [][]uint8, serialize serializer) string {
	boardStringBuilder := strings.Builder{}
	for _, row := range board {
		for _, cell := range row {
			boardStringBuilder.WriteString(fmt.Sprintf("%s", serialize(cell)))
		}
		boardStringBuilder.WriteString(fmt.Sprintln())
	}

	return boardStringBuilder.String()
}

func toCharacter(bit uint8) string {
	if bit == ALIVE {
		return ALIVE_CHARACTER
	} else {
		return DEAD_CHARACTER
	}
}

func toInt(bit uint8) string {
	return fmt.Sprintf("%d", bit)
}

func nextBoardState(currentBoard [][]uint8) [][]uint8 {
	nextState := createBoard(len(currentBoard), len(currentBoard[0]))
	for rowIndex, row := range currentBoard {
		for columnIndex, cell := range row {
			aliveNeighbors := countAliveNeighboards(currentBoard, rowIndex, columnIndex)
			nextState[rowIndex][columnIndex] = newCellState(cell, aliveNeighbors)
		}
	}
	return nextState
}

func countAliveNeighboards(currentBoard [][]uint8, row int, column int) uint8 {
	aliveNeighbors := 0

	for r := lowerBound(row); r <= upperBound(row, len(currentBoard)); r++ {
		for c := lowerBound(column); c <= upperBound(column, len(currentBoard[r])); c++ {
			// fmt.Printf("on(%d,%d)=%d", r, c, currentBoard[r][c])
			// fmt.Println()
			if (r != row || c != column) && currentBoard[r][c] == ALIVE {
				// fmt.Printf("Counted on(%d,%d)=%d", r, c, currentBoard[r][c])
				// fmt.Println()
				aliveNeighbors++
			}
		}
	}

	return uint8(aliveNeighbors)
}

func lowerBound(index int) int {
	if index == 0 {
		return 0
	} else {
		return index - 1
	}
}

func upperBound(index int, maxLength int) int {
	if index >= (maxLength - 1) {
		return index
	} else {
		return index + 1
	}
}

func newCellState(cell uint8, quantityOfAliveNeighbors uint8) uint8 {
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
