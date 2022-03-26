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

/*
- underpopulation
- right conditions
- overpopulation
- regeneration
*/

func TestDeadCellsWithNoNeighbors(t *testing.T) {
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
