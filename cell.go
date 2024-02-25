package main

const ALIVE_CHARACTER string = "\u2588"
const DEAD_CHARACTER string = "\u0020"

type Cell uint8

const ALIVE Cell = Cell(1)
const DEAD Cell = Cell(0)

func (cell Cell) toCharacter() string {
	if cell == ALIVE {
		return ALIVE_CHARACTER
	} else {
		return DEAD_CHARACTER
	}
}

func (cell Cell) newCellState(quantityOfAliveNeighbors uint8) Cell {
	// Overpopulation
	if quantityOfAliveNeighbors > 3 {
		return DEAD
	}

	// Revive
	if quantityOfAliveNeighbors == 3 {
		return ALIVE
	}

	// Right conditions
	if cell == ALIVE && quantityOfAliveNeighbors >= 2 && quantityOfAliveNeighbors <= 3 {
		return ALIVE
	}

	if quantityOfAliveNeighbors <= 1 {
		return DEAD
	}

	// no rule applied
	return cell
}
