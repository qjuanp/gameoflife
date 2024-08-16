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

func TestIterateNeighborsOfRow(t *testing.T) {
	boardExpected := Board{
		{false},
		{true},
		{false},
	}

	// first element, which is the prev neighgord
	val, hasNext, idx, it := boardExpected.IterateNeighborsOfRow(1)

	if val[0] {
		t.Error("Wrong `val`. Expected `false` in position 0")
	}

	if idx != 0 {
		t.Errorf("Wrong `idx=%d`. Expected position 0", idx)
	}

	if !hasNext {
		t.Error("Iterator `it` calculated wrong `hasNext`")
	}

	// second element, which is the row reviewed
	val, hasNext, idx = it()

	if !val[0] {
		t.Error("Wrong `val`. Expected `true` in position 1")
	}

	if idx != 1 {
		t.Errorf("Wrong `idx=%d`. Expected position 1", idx)
	}

	if !hasNext {
		t.Error("Iterator `it` calculated wrong `hasNext`")
	}

	// third element, which is the next row
	val, hasNext, idx = it()

	if val[0] {
		t.Error("Wrong `val`. Expected `false` in position 2")
	}

	if idx != 2 {
		t.Errorf("Wrong `idx=%d`. Expected position 2", idx)
	}

	if !hasNext {
		t.Error("Iterator `it` calculated wrong `hasNext`")
	}
}

func TestIterateNeighborsOfRowLowerBound(t *testing.T) {
	boardExpected := Board{
		{false},
		{true},
		{false},
	}

	// first element, which is the prev neighgord
	val, hasNext, idx, it := boardExpected.IterateNeighborsOfRow(0)

	if val[0] {
		t.Error("Wrong `val`. Expected `false` in position 2")
	}

	if idx != 2 {
		t.Errorf("Wrong `idx=%d`. Expected position 2", idx)
	}

	if !hasNext {
		t.Error("Iterator `it` calculated wrong `hasNext`")
	}

	// second element, which is the row reviewed
	val, hasNext, idx = it()

	if val[0] {
		t.Error("Wrong `val`. Expected `true` in position 1")
	}

	if idx != 0 {
		t.Errorf("Wrong `idx=%d`. Expected position 0", idx)
	}

	if !hasNext {
		t.Error("Iterator `it` calculated wrong `hasNext`")
	}

	// third element, which is the next row
	val, hasNext, idx = it()

	if !val[0] {
		t.Error("Wrong `val`. Expected `false` in position 2")
	}

	if idx != 1 {
		t.Errorf("Wrong `idx=%d`. Expected position 1", idx)
	}

	if !hasNext {
		t.Error("Iterator `it` calculated wrong `hasNext`")
	}
}

func TestIterateNeighborsOfRowLowerBound100x100(t *testing.T) {
	boardExpected := NewEmptyBoardOfSize(100, 100)

	row, hasNext, idx, it := boardExpected.IterateNeighborsOfRow(0)

	// first element, which is the prev neighgord

	if len(row) != 100 {
		t.Error("Wrong `row` length. Expected 100")
	}

	if idx != 99 {
		t.Errorf("Wrong `idx=%d`. Expected position 99", idx)
	}

	if !hasNext {
		t.Error("Iterator `it` calculated wrong `hasNext`")
	}

	// second element, which is the row reviewed
	row, hasNext, idx = it()

	if len(row) != 100 {
		t.Error("Wrong `row` length. Expected 100")
	}

	if idx != 0 {
		t.Errorf("Wrong `idx=%d`. Expected position 0", idx)
	}

	if !hasNext {
		t.Error("Iterator `it` calculated wrong `hasNext`")
	}

	// third element, which is the next row
	row, hasNext, idx = it()

	if len(row) != 100 {
		t.Error("Wrong `row` length. Expected 100")
	}

	if idx != 1 {
		t.Errorf("Wrong `idx=%d`. Expected position 1", idx)
	}

	if !hasNext {
		t.Error("Iterator `it` calculated wrong `hasNext`")
	}
}

func TestIterateNeighborsOfRowUpperBound(t *testing.T) {
	boardExpected := Board{
		{false},
		{true},
		{false},
	}

	// first element, which is the prev neighgord
	val, hasNext, idx, it := boardExpected.IterateNeighborsOfRow(2)

	if !val[0] {
		t.Error("Wrong `val`. Expected `false` in position 1")
	}

	if idx != 1 {
		t.Errorf("Wrong `idx=%d`. Expected position 1", idx)
	}

	if !hasNext {
		t.Error("Iterator `it` calculated wrong `hasNext`")
	}

	// second element, which is the row reviewed
	val, hasNext, idx = it()

	if val[0] {
		t.Error("Wrong `val`. Expected `true` in position 2")
	}

	if idx != 2 {
		t.Errorf("Wrong `idx=%d`. Expected position 2", idx)
	}

	if !hasNext {
		t.Error("Iterator `it` calculated wrong `hasNext`")
	}

	// third element, which is the next row
	val, hasNext, idx = it()

	if val[0] {
		t.Error("Wrong `val`. Expected `false` in position 0")
	}

	if idx != 0 {
		t.Errorf("Wrong `idx=%d`. Expected position 0", idx)
	}

	if !hasNext {
		t.Error("Iterator `it` calculated wrong `hasNext`")
	}
}

func TestIterateNeighborsOfColumn(t *testing.T) {
	boardExpected := Board{
		{false, true, false},
	}

	// first element, which is the prev neighgord
	val, hasNext, idx, it := boardExpected.IterateNightborsOf(boardExpected[0], 1)

	if val {
		t.Error("Wrong `val`. Expected `false` in position 0")
	}

	if idx != 0 {
		t.Errorf("Wrong `idx=%d`. Expected position 0", idx)
	}

	if !hasNext {
		t.Error("Iterator `it` calculated wrong `hasNext`")
	}

	// second element, which is the row reviewed
	val, hasNext, idx = it()

	if !val {
		t.Error("Wrong `val`. Expected `true` in position 1")
	}

	if idx != 1 {
		t.Errorf("Wrong `idx=%d`. Expected position 1", idx)
	}

	if !hasNext {
		t.Error("Iterator `it` calculated wrong `hasNext`")
	}

	// third element, which is the next row
	val, hasNext, idx = it()

	if val {
		t.Error("Wrong `val`. Expected `false` in position 2")
	}

	if idx != 2 {
		t.Errorf("Wrong `idx=%d`. Expected position 2", idx)
	}

	if !hasNext {
		t.Error("Iterator `it` calculated wrong `hasNext`")
	}
}

func TestIterateNeighborsOfColumnsLowerBound(t *testing.T) {
	boardExpected := Board{
		{false, true, false},
	}

	// first element, which is the prev neighgord
	val, hasNext, idx, it := boardExpected.IterateNightborsOf(boardExpected[0], 0)

	if val {
		t.Error("Wrong `val`. Expected `false` in position 2")
	}

	if idx != 2 {
		t.Errorf("Wrong `idx=%d`. Expected position 2", idx)
	}

	if !hasNext {
		t.Error("Iterator `it` calculated wrong `hasNext`")
	}

	// second element, which is the row reviewed
	val, hasNext, idx = it()

	if val {
		t.Error("Wrong `val`. Expected `true` in position 1")
	}

	if idx != 0 {
		t.Errorf("Wrong `idx=%d`. Expected position 0", idx)
	}

	if !hasNext {
		t.Error("Iterator `it` calculated wrong `hasNext`")
	}

	// third element, which is the next row
	val, hasNext, idx = it()

	if !val {
		t.Error("Wrong `val`. Expected `false` in position 2")
	}

	if idx != 1 {
		t.Errorf("Wrong `idx=%d`. Expected position 1", idx)
	}

	if !hasNext {
		t.Error("Iterator `it` calculated wrong `hasNext`")
	}
}

func TestIterateNeighborsOfColumnsUpperBound(t *testing.T) {
	boardExpected := Board{
		{false, true, false},
	}

	// first element, which is the prev neighgord
	val, hasNext, idx, it := boardExpected.IterateNightborsOf(boardExpected[0], 2)

	if !val {
		t.Error("Wrong `val`. Expected `false` in position 1")
	}

	if idx != 1 {
		t.Errorf("Wrong `idx=%d`. Expected position 1", idx)
	}

	if !hasNext {
		t.Error("Iterator `it` calculated wrong `hasNext`")
	}

	// second element, which is the row reviewed
	val, hasNext, idx = it()

	if val {
		t.Error("Wrong `val`. Expected `true` in position 2")
	}

	if idx != 2 {
		t.Errorf("Wrong `idx=%d`. Expected position 2", idx)
	}

	if !hasNext {
		t.Error("Iterator `it` calculated wrong `hasNext`")
	}

	// third element, which is the next row
	val, hasNext, idx = it()

	if val {
		t.Error("Wrong `val`. Expected `false` in position 0")
	}

	if idx != 0 {
		t.Errorf("Wrong `idx=%d`. Expected position 0", idx)
	}

	if !hasNext {
		t.Error("Iterator `it` calculated wrong `hasNext`")
	}
}
