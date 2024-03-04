package main

import (
	"reflect"
	"testing"

	"github.com/qjuanp/gameoflife/board"
)

func detailedErrorResult(t *testing.T, boardResult board.Board, boardExpected board.Board) {
	t.Errorf("Board result:\n%s\n", boardResult.ToString())
	t.Errorf("Board expected:\n%s\n", boardExpected.ToString())

	for rowIndex, row := range boardResult {
		for columnIndex, cell := range row {
			if cell != boardExpected[rowIndex][columnIndex] {
				t.Errorf("At(row=%d,column%d) Value=%t | Expected=%t", rowIndex, columnIndex, cell, boardExpected[rowIndex][columnIndex])
			}
		}
	}
}

func TestCreateBoard(t *testing.T) {
	const columnsExpected uint = 5
	const rowsExpected uint = 5
	boardExpected := board.Board{
		{board.DEAD, board.DEAD, board.DEAD, board.DEAD, board.DEAD},
		{board.DEAD, board.DEAD, board.DEAD, board.DEAD, board.DEAD},
		{board.DEAD, board.DEAD, board.DEAD, board.DEAD, board.DEAD},
		{board.DEAD, board.DEAD, board.DEAD, board.DEAD, board.DEAD},
		{board.DEAD, board.DEAD, board.DEAD, board.DEAD, board.DEAD},
	}

	boardResult := board.NewEmptyBoardOfSize(rowsExpected, columnsExpected)
	if !reflect.DeepEqual(boardResult, boardExpected) {
		t.Errorf("Unexpected created board")
		detailedErrorResult(t, boardResult, boardExpected)
	}
}

func TestRandomInitializationSize(t *testing.T) {
	const columnsExpected uint = 5
	const rowsExpected uint = 5
	seed := int64(1)

	boardResult := board.NewRandomBoard(rowsExpected, columnsExpected, seed)
	rowsResult := boardResult.NumberOfColumns()
	columnsResult := boardResult.NumberOfRows()

	if rowsResult != rowsExpected {
		t.Errorf("Rows: %d, wanted %d", rowsResult, rowsExpected)
	}

	if columnsResult != columnsExpected {
		t.Errorf("Columns: %d, wanted %d", columnsResult, columnsExpected)
	}
}

func TestRandomInitializationBoard(t *testing.T) {
	const (
		columnsExpected uint = 5
		rowsExpected    uint = 5
	)
	seed := int64(1)
	boardExpected := board.Board{
		{board.ALIVE, board.ALIVE, board.ALIVE, board.DEAD, board.DEAD},
		{board.ALIVE, board.DEAD, board.DEAD, board.DEAD, board.DEAD},
		{board.ALIVE, board.ALIVE, board.DEAD, board.DEAD, board.DEAD},
		{board.DEAD, board.DEAD, board.DEAD, board.ALIVE, board.DEAD},
		{board.DEAD, board.DEAD, board.ALIVE, board.ALIVE, board.DEAD},
	}

	boardResult := board.NewRandomBoard(rowsExpected, columnsExpected, seed)

	if !reflect.DeepEqual(boardResult, boardExpected) {
		t.Errorf("Unexpected initialized board")
		detailedErrorResult(t, boardResult, boardExpected)
	}
}

// func TestLowerBoundCentreCell(t *testing.T) {
// 	position := 3
// 	expectedLowerBound := 2

// 	result := lowerBound(position)

// 	if result != expectedLowerBound {
// 		t.Errorf("Lower bound mismatch result=%d | expected=%d", result, expectedLowerBound)
// 	}
// }

// func TestLowerBoundCornerCell(t *testing.T) {
// 	position := board.DEAD
// 	expectedLowerBound := board.DEAD

// 	result := lowerBound(position)

// 	if result != expectedLowerBound {
// 		t.Errorf("Lower bound mismatch result=%d | expected=%d", result, expectedLowerBound)
// 	}
// }

// func TestUpperBoundCentreCell(t *testing.T) {
// 	position := 3
// 	length := 5
// 	expectedLowerBound := 4

// 	result := upperBound(position, length)

// 	if result != expectedLowerBound {
// 		t.Errorf("Lower bound mismatch result=%d | expected=%d", result, expectedLowerBound)
// 	}
// }

// func TestUpperBoundCornerCell(t *testing.T) {
// 	position := 4
// 	length := 5
// 	expectedLowerBound := 4

// 	result := upperBound(position, length)

// 	if result != expectedLowerBound {
// 		t.Errorf("Lower bound mismatch result=%d | expected=%d", result, expectedLowerBound)
// 	}
// }

func TestCountAliveNeighborsAllDead(t *testing.T) {
	currentBoard := board.Board{
		{board.DEAD, board.DEAD, board.DEAD},
		{board.DEAD, board.DEAD, board.DEAD},
		{board.DEAD, board.DEAD, board.DEAD},
	}

	game := GameOfLife{currentBoard}
	result := game.countAliveNeighboards(1, 1)

	if result != 0 {
		t.Error("Miscalculation of alive neighbors when all dead")
	}
}

func TestCountAliveNeighbors2AliveFromCenter(t *testing.T) {
	currentBoard := board.Board{
		{board.DEAD, board.DEAD, board.ALIVE},
		{board.DEAD, board.DEAD, board.DEAD},
		{board.ALIVE, board.DEAD, board.DEAD},
	}

	game := GameOfLife{currentBoard}
	result := game.countAliveNeighboards(1, 1)

	if result != 2 {
		t.Error("Miscalculation of alive neighbors when all dead")
	}
}

func TestCountAliveNeighborsAllAliveFromUpperLeftcorner(t *testing.T) {
	currentBoard := board.Board{
		{board.ALIVE, board.ALIVE, board.DEAD},
		{board.ALIVE, board.ALIVE, board.DEAD},
		{board.DEAD, board.DEAD, board.DEAD},
	}
	game := GameOfLife{currentBoard}
	result := game.countAliveNeighboards(1, 1)

	if result != 3 {
		t.Errorf("Miscalculation of alive neighbors when all dead | result=%d", result)
	}
}

/*
- underpopulation
- right conditions
- overpopulation
- regeneration
*/

func TestDeadCellsWithNoNeighborsShouldStayDead(t *testing.T) {
	currentBoard := board.Board{
		{board.DEAD, board.DEAD, board.DEAD},
		{board.DEAD, board.DEAD, board.DEAD},
		{board.DEAD, board.DEAD, board.DEAD},
	}

	expectedNextGenerationBoard := board.Board{
		{board.DEAD, board.DEAD, board.DEAD},
		{board.DEAD, board.DEAD, board.DEAD},
		{board.DEAD, board.DEAD, board.DEAD},
	}

	game := GameOfLife{currentBoard}

	nextGenerationGame := game.next()

	if !reflect.DeepEqual(nextGenerationGame.Board, expectedNextGenerationBoard) {
		t.Errorf("Rule error: Dead cells with no neighbors")
		detailedErrorResult(t, nextGenerationGame.Board, expectedNextGenerationBoard)
	}
}

func TestDeadCellsWithExcatly3NeighborsShouldComeAlive(t *testing.T) {
	currentBoard := board.Board{
		{board.DEAD, board.DEAD, board.ALIVE},
		{board.DEAD, board.ALIVE, board.ALIVE},
		{board.DEAD, board.DEAD, board.DEAD},
	}

	expectedNextGenerationBoard := board.Board{
		{board.DEAD, board.ALIVE, board.ALIVE},
		{board.DEAD, board.ALIVE, board.ALIVE},
		{board.DEAD, board.DEAD, board.DEAD},
	}

	game := GameOfLife{currentBoard}

	nextGenerationGame := game.next()

	if !reflect.DeepEqual(nextGenerationGame.Board, expectedNextGenerationBoard) {
		t.Errorf("Rule error: Dead cells with no neighbors")
		detailedErrorResult(t, nextGenerationGame.Board, expectedNextGenerationBoard)
	}
}

func TestDeadCellsWithExcatlyAliveNeighborShouldBeDead(t *testing.T) {
	currentBoard := board.Board{
		{board.DEAD, board.DEAD, board.DEAD},
		{board.ALIVE, board.ALIVE, board.ALIVE},
		{board.DEAD, board.DEAD, board.DEAD},
	}

	expectedNextGenerationBoard := board.Board{
		{board.DEAD, board.ALIVE, board.DEAD},
		{board.DEAD, board.ALIVE, board.DEAD},
		{board.DEAD, board.ALIVE, board.DEAD},
	}

	game := GameOfLife{currentBoard}
	nextGenerationGame := game.next()

	if !reflect.DeepEqual(nextGenerationGame.Board, expectedNextGenerationBoard) {
		t.Errorf("Rule error: Dead cells with no neighbors")
		detailedErrorResult(t, nextGenerationGame.Board, expectedNextGenerationBoard)
	}
}
