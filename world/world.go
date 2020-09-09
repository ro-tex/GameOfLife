package world

import "fmt"

/*
This is the main logic of the game.

The code is not supposed to be multi-threaded at this point, so there is no locking.
*/

// World is a complete representation of the game's state and logic.
// The game is Conway's Game of Life.
type World struct {
	// The Board is the World itself. The less significant index denotes width
	// and the more significant denotes height, i.e. Board[height][width].
	// TODO This can very much be represented by bits, reducing the needed memory 8 times.
	Board [][]byte // TODO unexport again

	// Maximal dimensions of the World. Anything beyond that is considered dead.
	MaxHeight int // TODO unexport again
	MaxWidth  int // TODO unexport again
}

// NextGen calculates the next generation of the World, growing it if necessary.
//
// Rules:
// Any live cell with fewer than two live neighbours dies, as if by underpopulation.
// Any live cell with two or three live neighbours lives on to the next generation.
// Any live cell with more than three live neighbours dies, as if by overpopulation.
// Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
//
// These rules, which compare the behavior of the automaton to real life, can be condensed into the following:
// Any live cell with two live neighbours survives.
// Any cell with three live neighbours stays/becomes a live cell.
// All other live cells die in the next generation. Similarly, all other dead cells stay dead.
func (wo *World) NextGen() {
	wo.growIfNeeded()
	// Create the next generation Board.
	newBoard := make([][]byte, len(wo.Board))
	woWidth := len(wo.Board[0])
	for i := range newBoard {
		newBoard[i] = make([]byte, woWidth)
	}
	// Calculate the next generation.
	for h := range wo.Board {
		for w := range wo.Board[h] {
			n := wo.neighbours(h, w)
			// Live cell survives, dead cell springs to life.
			if n == 3 {
				newBoard[h][w] = 1
				continue
			}
			// Live cell survives.
			if wo.Board[h][w] > 0 && n == 2 {
				newBoard[h][w] = 1
			}
		}
	}
	// Swap the next generation into the world.
	wo.Board = newBoard
}

// Seed sets the state of the World.
func (wo *World) Seed(seed [][]byte) {
	wo.Board = seed
}

// State returns the state of the World.
func (wo *World) State() [][]byte {
	return wo.Board
}

// growIfNeeded adds a single new row to each side of the world where a cell is
// living on the edge, unless the world has reached its maximum size.
// The order is bottom, top, right, left.
func (wo *World) growIfNeeded() {
	rowHasLiveCells := func(h int) bool {
		for i := range wo.Board[h] {
			if wo.Board[h][i] > 0 {
				return true
			}
		}
		return false
	}
	colHasLiveCells := func(w int) bool {
		for i := range wo.Board {
			if wo.Board[i][w] > 0 {
				return true
			}
		}
		return false
	}
	if len(wo.Board) < wo.MaxHeight {
		if rowHasLiveCells(0) {
			wo.Board = append([][]byte{make([]byte, len(wo.Board[0]))}, wo.Board...)
		}
	}
	if len(wo.Board) < wo.MaxHeight {
		if rowHasLiveCells(len(wo.Board) - 1) {
			wo.Board = append(wo.Board, make([]byte, len(wo.Board[0])))
		}
	}
	if len(wo.Board[0]) < wo.MaxWidth {
		if colHasLiveCells(0) {
			for i := range wo.Board {
				wo.Board[i] = append([]byte{0}, wo.Board[i]...)
			}
		}
	}
	if len(wo.Board[0]) < wo.MaxWidth {
		if colHasLiveCells(len(wo.Board[0]) - 1) {
			for i := range wo.Board {
				wo.Board[i] = append(wo.Board[i], 0)
			}
		}
	}
}

// neighbours counts the neighbours the given cell has in the World.
func (wo *World) neighbours(h, w int) int {
	n := 0
	woHeight := len(wo.Board)
	woWidth := len(wo.Board[0])
	// Upstairs neighbours.
	if h > 0 {
		if w > 0 && wo.Board[h-1][w-1] > 0 {
			n++
		}
		if wo.Board[h-1][w] > 0 {
			n++
		}
		if w < woWidth-1 && wo.Board[h-1][w+1] > 0 {
			n++
		}
	}
	// Own floor neighbours.
	if w > 0 && wo.Board[h][w-1] > 0 {
		n++
	}
	if w < woWidth-1 && wo.Board[h][w+1] > 0 {
		n++
	}
	// Downstairs neighbours.
	if h < woHeight-1 {
		if w > 0 && wo.Board[h+1][w-1] > 0 {
			n++
		}
		if wo.Board[h+1][w] > 0 {
			n++
		}
		if w < woWidth-1 && wo.Board[h+1][w+1] > 0 {
			n++
		}
	}
	return n
}

func (wo *World) Print() {
	for h := range wo.Board {
		for w := range wo.Board[h] {
			if wo.Board[h][w] > 0 {
				fmt.Print(" ⌘ ")
			} else {
				fmt.Print("   ")
			}
		}
		fmt.Println()
	}
}
