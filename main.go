package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	fmt.Println("Project setup done")
	board := ramdomInitialization(5, 5, 1)
	fmt.Print(boardToString(board, toInt))
}

type serializer func(uint8) string

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

func boardToString(board [][]uint8, serialize serializer) string {
	boardStringBuilder := strings.Builder{}
	for _, row := range board {
		for _, cell := range row {
			boardStringBuilder.WriteString(fmt.Sprintf("%s", serialize(cell)))
		}
		boardStringBuilder.WriteString(fmt.Sprintln())
	}

	return boardStringBuilder.String()
}

func toCharacter(bit uint8) string {
	if bit == ALIVE {
		return ALIVE_CHARACTER
	} else {
		return DEAD_CHARACTER
	}
}

func toInt(bit uint8) string {
	return fmt.Sprintf("%d", bit)
}
