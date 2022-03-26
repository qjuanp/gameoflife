package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Project setup done")
	board := ramdomInitialization(5, 5, time.Now().UnixNano())
	printBoard(board)
}

const ALIVE uint8 = 1
const ALIVE_CHARACTER string = "\u2588"

const DEAD uint8 = 0
const DEAD_CHARACTER string = "\u0020"

func ramdomInitialization(rows uint32, columns uint32, seed int64) [][]uint8 {
	var board [][]uint8 = make([][]uint8, rows)
	rand.Seed(seed)

	for row := range board {
		board[row] = make([]uint8, columns)
		for column := range board[row] {
			randomValue := rand.Float64()
			if randomValue >= 0.5 {
				board[row][column] = ALIVE
			} else {
				board[row][column] = DEAD
			}
		}
	}

	return board
}

func printBoard(board [][]uint8) {
	for _, row := range board {
		for _, cell := range row {
			fmt.Printf("%s", toCharacter(cell))
		}
		fmt.Println()
	}
}

func toCharacter(bit uint8) string {
	if bit == ALIVE {
		return ALIVE_CHARACTER
	} else {
		return DEAD_CHARACTER
	}
}
