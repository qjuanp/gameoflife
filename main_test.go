package main

import (
	"testing"
	"time"
)

func TestRandomInitializing(t *testing.T) {
	const columnsExpected uint32 = 5
	const rowsExpected uint32 = 5
	seed := time.Now().UnixNano()

	initializedBoard := ramdomInitialization(rowsExpected, columnsExpected, seed)
	rowsResult := uint32(len(initializedBoard))
	columnsResult := uint32(len(initializedBoard[0]))

	if rowsResult != rowsExpected {
		t.Errorf("Rows: %d, wanted %d", rowsResult, rowsExpected)
	}

	if columnsResult != columnsExpected {
		t.Errorf("Columns: %d, wanted %d", columnsResult, columnsExpected)
	}
}
