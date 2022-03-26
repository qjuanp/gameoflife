package main

import (
	"reflect"
	"testing"
)

func detailedErrorResult(t *testing.T, boardResult [][]uint8, boardExpected [][]uint8) {
	t.Errorf("Board result:\n%s\n", boardToString(boardResult, toInt))
	t.Errorf("Board expected:\n%s\n", boardToString(boardExpected, toInt))

	for rowIndex, row := range boardResult {
		for columnIndex, cell := range row {
			if cell != boardExpected[rowIndex][columnIndex] {
				t.Errorf("At(row=%d,column%d) Value=%d | Expected=%d", rowIndex, columnIndex, cell, boardExpected[rowIndex][columnIndex])
			}
		}
	}
}

func TestCreateBoard(t *testing.T) {
	const columnsExpected int = 5
	const rowsExpected int = 5
	boardExpected := [][]uint8{
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}

	boardResult := createBoard(rowsExpected, columnsExpected)
	if !reflect.DeepEqual(boardResult, boardExpected) {
		t.Errorf("Unexpected created board")
		detailedErrorResult(t, boardResult, boardExpected)
	}
}

func TestRandomInitializationSize(t *testing.T) {
	const columnsExpected uint32 = 5
	const rowsExpected uint32 = 5
	seed := int64(1)

	boardResult := ramdomInitialization(rowsExpected, columnsExpected, seed)
	rowsResult := uint32(len(boardResult))
	columnsResult := uint32(len(boardResult[0]))

	if rowsResult != rowsExpected {
		t.Errorf("Rows: %d, wanted %d", rowsResult, rowsExpected)
	}

	if columnsResult != columnsExpected {
		t.Errorf("Columns: %d, wanted %d", columnsResult, columnsExpected)
	}
}

func TestRandomInitializationBoard(t *testing.T) {
	const (
		columnsExpected uint32 = 5
		rowsExpected    uint32 = 5
	)
	seed := int64(1)
	boardExpected := [][]uint8{
		{1, 1, 1, 0, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 1, 0},
		{0, 0, 1, 1, 0},
	}

	boardResult := ramdomInitialization(rowsExpected, columnsExpected, seed)

	if !reflect.DeepEqual(boardResult, boardExpected) {
		t.Errorf("Unexpected initialized board")
		detailedErrorResult(t, boardResult, boardExpected)
	}
}

func TestLowerBoundCentreCell(t *testing.T) {
	position := 3
	expectedLowerBound := 2

	result := lowerBound(position)

	if result != expectedLowerBound {
		t.Errorf("Lower bound mismatch result=%d | expected=%d", result, expectedLowerBound)
	}
}

func TestLowerBoundCornerCell(t *testing.T) {
	position := 0
	expectedLowerBound := 0

	result := lowerBound(position)

	if result != expectedLowerBound {
		t.Errorf("Lower bound mismatch result=%d | expected=%d", result, expectedLowerBound)
	}
}

func TestUpperBoundCentreCell(t *testing.T) {
	position := 3
	length := 5
	expectedLowerBound := 4

	result := upperBound(position, length)

	if result != expectedLowerBound {
		t.Errorf("Lower bound mismatch result=%d | expected=%d", result, expectedLowerBound)
	}
}

func TestUpperBoundCornerCell(t *testing.T) {
	position := 4
	length := 5
	expectedLowerBound := 4

	result := upperBound(position, length)

	if result != expectedLowerBound {
		t.Errorf("Lower bound mismatch result=%d | expected=%d", result, expectedLowerBound)
	}
}

func TestCountAliveNeighborsAllDead(t *testing.T) {
	currentBoard := [][]uint8{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	result := countAliveNeighboards(currentBoard, 1, 1)

	if result != 0 {
		t.Error("Miscalculation of alive neighbors when all dead")
	}
}

func TestCountAliveNeighbors2AliveFromCenter(t *testing.T) {
	currentBoard := [][]uint8{
		{0, 0, 1},
		{0, 0, 0},
		{1, 0, 0},
	}

	result := countAliveNeighboards(currentBoard, 1, 1)

	if result != 2 {
		t.Error("Miscalculation of alive neighbors when all dead")
	}
}

func TestCountAliveNeighborsAllAliveFromUpperLeftcorner(t *testing.T) {
	currentBoard := [][]uint8{
		{1, 1, 0},
		{1, 1, 0},
		{0, 0, 0},
	}

	result := countAliveNeighboards(currentBoard, 0, 0)

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
	currentBoard := [][]uint8{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	expectedNextGenerationBoard := [][]uint8{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	resultNextGenerationBoard := nextBoardState(currentBoard)

	if !reflect.DeepEqual(resultNextGenerationBoard, expectedNextGenerationBoard) {
		t.Errorf("Rule error: Dead cells with no neighbors")
		detailedErrorResult(t, resultNextGenerationBoard, expectedNextGenerationBoard)
	}
}

func TestDeadCellsWithExcatly3NeighborsShouldComeAlive(t *testing.T) {
	currentBoard := [][]uint8{
		{0, 0, 1},
		{0, 1, 1},
		{0, 0, 0},
	}

	expectedNextGenerationBoard := [][]uint8{
		{0, 1, 1},
		{0, 1, 1},
		{0, 0, 0},
	}

	resultNextGenerationBoard := nextBoardState(currentBoard)

	if !reflect.DeepEqual(resultNextGenerationBoard, expectedNextGenerationBoard) {
		t.Errorf("Rule error: Dead cells with no neighbors")
		detailedErrorResult(t, resultNextGenerationBoard, expectedNextGenerationBoard)
	}
}

func TestDeadCellsWithExcatly1NeighborShouldBeDead(t *testing.T) {
	currentBoard := [][]uint8{
		{0, 0, 0},
		{1, 1, 1},
		{0, 0, 0},
	}

	expectedNextGenerationBoard := [][]uint8{
		{0, 1, 0},
		{0, 1, 0},
		{0, 1, 0},
	}

	resultNextGenerationBoard := nextBoardState(currentBoard)

	if !reflect.DeepEqual(resultNextGenerationBoard, expectedNextGenerationBoard) {
		t.Errorf("Rule error: Dead cells with no neighbors")
		detailedErrorResult(t, resultNextGenerationBoard, expectedNextGenerationBoard)
	}
}
