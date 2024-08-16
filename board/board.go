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
	return board.ToStringWith(map[bool]string{
		true:  "\u2588",
		false: "\u0020",
	})
}

func (board *Board) ToNumbers() string {
	return board.ToStringWith(map[bool]string{
		true:  "1",
		false: "0",
	})
}

func (board *Board) ToStringWith(character map[bool]string) string {
	strBuilder := strings.Builder{}
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

// TODO: Make borderless optional
func (board *Board) IterateNeighborsOfRow(row int) (val []bool, hasNext bool, idx int, it func() (val []bool, hasNext bool, idx int)) {
	length := len(*board)
	var currentIndex int
	count := 1

	if row == 0 {
		currentIndex = int(board.CountRows()) - 1
	} else if row >= length {
		currentIndex = length - 2
	} else {
		currentIndex = row - 1
	}
	return (*board)[currentIndex], (count <= 3), currentIndex, func() (val []bool, hasNext bool, idx int) {
		currentIndex++
		count++

		if currentIndex >= length {
			currentIndex = 0
		}

		currentRow := (*board)[currentIndex]

		if len(currentRow) == 0 {
			println("Here's an error")
		}

		return (*board)[currentIndex], (count <= 3), currentIndex

	}
}

func (board *Board) IterateNightborsOf(row []bool, column int) (cell bool, hasNext bool, idx int, it func() (cell bool, hasNext bool, idx int)) {
	length := len(row)
	var currentIndex int
	count := 1

	if column == 0 {
		currentIndex = int(board.CountColumns()) - 1
	} else if column >= length {
		currentIndex = length - 2
	} else {
		currentIndex = column - 1
	}

	return row[currentIndex], (count <= 3), currentIndex, func() (cell bool, hasNext bool, idx int) {
		currentIndex++
		count++

		if currentIndex >= length {
			currentIndex = 0
		}

		if length == 0 {
			println("here's an error")
		}

		return row[currentIndex], (count <= 3), currentIndex
	}
}
