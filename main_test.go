package main

import (
	"testing"
)

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

	for rowIndex, row := range boardResult {
		for columnIndex, cell := range row {
			if cell != boardExpected[rowIndex][columnIndex] {
				t.Errorf("At(row=%d,column%d) Value=%d | Expected=%d", rowIndex, columnIndex, cell, boardExpected[rowIndex][columnIndex])
			}
		}
	}
}
