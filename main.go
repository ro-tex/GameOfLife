package main

import (
	"fmt"
	"github.com/ro-tex/GameOfLife/world"
	"time"
)

var (
	Blinker = [][]byte{
		{0, 0, 0},
		{1, 1, 1},
		{0, 0, 0},
	}

	Glider = [][]byte{
		{1, 0, 0},
		{0, 1, 1},
		{1, 1, 0},
	}
)

/*
TODO
	- gliders should fly off the world, not get crumpled into the wall of it
	- tests :)
	- unexport stuff
	- proper constructor
	- sane defaults (seed with a blinker)
*/

func main() {
	max := 13
	w := world.World{
		Glider,
		max,
		max,
	}
	// Maximise:
	w.Board = append(w.Board, make([][]byte, max-len(w.Board))...)
	for i := range w.Board {
		w.Board[i] = append(w.Board[i], make([]byte, max-len(w.Board[i]))...)
	}
	for i := 0; i < 1000; i++ {
		fmt.Println("\n\n\n\n\n\n\n\n\n\n\n")
		w.Print()
		w.NextGen()
		fmt.Println("Board height", len(w.Board), ", width", len(w.Board[0]))
		time.Sleep(200 * time.Millisecond)
	}
}
