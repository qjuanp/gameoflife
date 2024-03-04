package main

import (
	"fmt"
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

func BenchmarkIntialization(b *testing.B) {
	for i := 1; i < 10; i++ {
		boardSize := uint(i * 1000)
		b.Run(fmt.Sprintf("Intialization size=%d", boardSize), func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				NewGameOfSize(boardSize, boardSize)
			}
		})
	}
}

func BenchmarkNextStateCalculation(b *testing.B) {
	for i := 1; i < 10; i++ {
		boardSize := uint(i * 1000)
		b.Run(fmt.Sprintf("Next Move on board of size size=%d", boardSize), func(b *testing.B) {
			game := NewGameOfSize(boardSize, boardSize)
			for j := 0; j < b.N; j++ {
				game.next()
			}
		})
	}
}
