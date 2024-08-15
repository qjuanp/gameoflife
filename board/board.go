package board

import (
	"fmt"
	"math/rand"
	"strings"
)

type Board [][]bool

func NewEmptyBoardOfSize(rows uint, columns uint) Board {
	var newBoard Board = make(Board, rows)

	for row := range newBoard {
		newBoard[row] = make([]bool, columns)
	}

	return newBoard
}

func NewEmptyBoardAsBigAs(board *Board) Board {
	return NewEmptyBoardOfSize(board.CountRows(), board.CountColumns())
}

func NewRandomBoard(rows uint, columns uint, seed int64) Board {
	var newBoard Board = make(Board, rows)

	random := rand.New(rand.NewSource(seed))

	for row := range newBoard {
		newBoard[row] = make([]bool, columns)
		for column := range newBoard[row] {
			randomValue := random.Float64()
			if randomValue >= 0.5 {
				newBoard[row][column] = true
			} else {
				newBoard[row][column] = false
			}
		}
	}

	return newBoard
}

func (board *Board) ToString() string {
	strBuilder := strings.Builder{}
	character := map[bool]string{
		true:  "\u2588",
		false: "\u0020",
	}
	for _, row := range *board {
		for _, cell := range row {
			strBuilder.WriteString(fmt.Sprintf("%s", character[cell]))
		}
		strBuilder.WriteString(fmt.Sprintln())
	}

	return strBuilder.String()
}

func (board *Board) CountRows() uint {
	return uint(len(*board))
}

func (board *Board) CountColumns() uint {
	return uint(len((*board)[0]))
}

func (board *Board) LowerBound(index int) int {
	if index == 0 {
		return 0
	} else {
		return index - 1
	}
}

func (board *Board) RowsUpperBoundFor(index int) int {
	return upperBound(index, len(*board))
}

func (board *Board) ColumnsUpperBoundFor(index int) int {
	return upperBound(index, len((*board)[0]))
}

func upperBound(index int, maxLength int) int {
	if index >= maxLength-1 {
		return index
	} else {
		return index + 1
	}
}
