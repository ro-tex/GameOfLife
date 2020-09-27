package world

import (
	"fmt"
)

/*
This is the main logic of the game.

The code is not supposed to be multi-threaded at this point, so there is no locking.
*/

// padding is a region around the board that's not displayed. Its goal is to
// allow moving shapes to "fly off the screen"
const padding = 10

// World is a complete representation of the game's state and logic.
// The game is Conway's Game of Life.
type World struct {
	// The board is the World itself. The less significant index denotes width
	// and the more significant denotes height, i.e. board[height][width].
	// TODO This can very much be represented by bits, reducing the needed memory 8 times.
	board [][]byte
	sweep bool // every other gen we wipe the edge of the world
}

// New creates a new empty World with a starting height and a max height.
func New(h, w int) World {
	wo := World{
		board: make([][]byte, h+2*padding),
	}
	for h := range wo.board {
		wo.board[h] = make([]byte, w+2*padding)
	}
	return wo
}

// NewFromSeed sets the state of the World.
func NewFromSeed(seed [][]byte) World {
	if len(seed) == 0 {
		// can't start with an empty world...
		seed = [][]byte{{0}}
	}
	wo := World{
		board: seed,
	}
	wo.pad()
	return wo
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
	// Create the next generation board.
	newBoard := make([][]byte, len(wo.board))
	woHeight := len(wo.board)
	woWidth := len(wo.board[0])
	for i := range newBoard {
		newBoard[i] = make([]byte, woWidth)
	}
	// Calculate the next generation.
	for h := range wo.board {
		for w := range wo.board[h] {
			n := wo.neighbours(h, w)
			// Live cell survives, dead cell springs to life.
			if n == 3 {
				newBoard[h][w] = 1
				continue
			}
			// Live cell survives.
			if wo.board[h][w] > 0 && n == 2 {
				newBoard[h][w] = 1
			}
		}
	}
	wo.sweep = !wo.sweep
	if wo.sweep {
		for h := 0; h < woHeight; h++ {
			newBoard[h][0] = 0
			newBoard[h][woWidth-1] = 0
		}
		for w := 0; w < woWidth; w++ {
			newBoard[0][w] = 0
			newBoard[woHeight-1][w] = 0
		}
	}
	// Swap the next generation into the world.
	wo.board = newBoard
}

// TODO either remove this method or make it respect the padding
// State returns the state of the World.
func (wo *World) State() [][]byte {
	return wo.board
}

// neighbours counts the neighbours the given cell has in the World.
func (wo *World) neighbours(h, w int) int {
	n := 0
	woHeight := len(wo.board)
	woWidth := len(wo.board[0])
	// Upstairs neighbours.
	if h > 0 {
		if w > 0 && wo.board[h-1][w-1] > 0 {
			n++
		}
		if wo.board[h-1][w] > 0 {
			n++
		}
		if w < woWidth-1 && wo.board[h-1][w+1] > 0 {
			n++
		}
	}
	// Own floor neighbours.
	if w > 0 && wo.board[h][w-1] > 0 {
		n++
	}
	if w < woWidth-1 && wo.board[h][w+1] > 0 {
		n++
	}
	// Downstairs neighbours.
	if h < woHeight-1 {
		if w > 0 && wo.board[h+1][w-1] > 0 {
			n++
		}
		if wo.board[h+1][w] > 0 {
			n++
		}
		if w < woWidth-1 && wo.board[h+1][w+1] > 0 {
			n++
		}
	}
	return n
}

// pad adds a layer of cells around the world that won't be displayed
func (wo *World) pad() {
	// pad horizontally
	hPad := make([]byte, padding, padding)
	for i := 0; i < len(wo.board); i++ {
		wo.board[i] = append(hPad, append(wo.board[i], hPad...)...)
	}
	// pad vertically
	vPad := make([][]byte, padding, padding)
	row := make([]byte, len(wo.board[0]), len(wo.board[0]))
	for i := 0; i < padding; i++ {
		vPad[i] = row
	}
	wo.board = append(vPad, append(wo.board, vPad...)...)
}

func (wo *World) Print() {
	for h := padding; h < len(wo.board)-padding; h++ {
		for w := padding; w < len(wo.board[h])-padding; w++ {
			if wo.board[h][w] > 0 {
				fmt.Print(" ⌘ ")
			} else {
				fmt.Print("   ")
			}
		}
		fmt.Println()
	}
}
