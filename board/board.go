package board

import (
	"fmt"
	"math/rand"
	"strings"
)

type Board [][]Cell

func NewEmptyBoardOfSize(rows uint, columns uint) Board {
	var newBoard Board = make(Board, rows)

	for row := range newBoard {
		newBoard[row] = make([]Cell, columns)
	}

	return newBoard
}

func NewEmptyBoardAsBigAs(board *Board) Board {
	return NewEmptyBoardOfSize(board.NumberOfRows(), board.NumberOfColumns())
}

func NewRandomBoard(rows uint, columns uint, seed int64) Board {
	var newBoard Board = make(Board, rows)

	random := rand.New(rand.NewSource(seed))

	for row := range newBoard {
		newBoard[row] = make([]Cell, columns)
		for column := range newBoard[row] {
			randomValue := random.Float64()
			if randomValue >= 0.5 {
				newBoard[row][column] = ALIVE
			} else {
				newBoard[row][column] = DEAD
			}
		}
	}

	return newBoard
}

func (board *Board) ToString() string {
	strBuilder := strings.Builder{}
	for _, row := range *board {
		for _, cell := range row {
			strBuilder.WriteString(fmt.Sprintf("%s", cell.toCharacter()))
		}
		strBuilder.WriteString(fmt.Sprintln())
	}

	return strBuilder.String()
}

func (board *Board) NumberOfRows() uint {
	return uint(len(*board))
}

func (board *Board) NumberOfColumns() uint {
	return uint(len((*board)[0]))
}

func (board *Board) LowerBound(index int) int {
	if index == 0 {
		return 0
	} else {
		return index - 1
	}
}

func (board *Board) UpperBound(index int, maxLength uint) int {
	if index >= int(maxLength-1) {
		return index
	} else {
		return index + 1
	}
}
