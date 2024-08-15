package board

import (
	"reflect"
	"testing"
)

func detailedErrorResult(t *testing.T, boardResult Board, boardExpected Board) {
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
	boardExpected := Board{
		{false, false, false, false, false},
		{false, false, false, false, false},
		{false, false, false, false, false},
		{false, false, false, false, false},
		{false, false, false, false, false},
	}

	boardResult := NewEmptyBoardOfSize(rowsExpected, columnsExpected)
	if !reflect.DeepEqual(boardResult, boardExpected) {
		t.Errorf("Unexpected created board")
		detailedErrorResult(t, boardResult, boardExpected)
	}
}

func TestRandomInitializationSize(t *testing.T) {
	const columnsExpected uint = 5
	const rowsExpected uint = 5
	seed := int64(1)

	boardResult := NewRandomBoard(rowsExpected, columnsExpected, seed)
	rowsResult := boardResult.CountColumns()
	columnsResult := boardResult.CountRows()

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
	boardExpected := Board{
		{true, true, true, false, false},
		{true, false, false, false, false},
		{true, true, false, false, false},
		{false, false, false, true, false},
		{false, false, true, true, false},
	}

	boardResult := NewRandomBoard(rowsExpected, columnsExpected, seed)

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
// 	position := false
// 	expectedLowerBound := false

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
