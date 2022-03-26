package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Project setup done")
}

const ALIVE uint8 = 1
const DEAD uint8 = 0

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
