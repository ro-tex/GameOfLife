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
	- tests :)
	- unexport stuff
	- proper constructor
	- sane defaults (seed with a blinker)
*/

func main() {
	w := world.World{
		Glider,
		20,
		20,
	}
	for i := 0; i < 1000; i++ {
		fmt.Println("\n\n\n\n\n\n\n\n\n\n\n")
		w.Print()
		w.NextGen()
		fmt.Println("Board height", len(w.Board), ", width", len(w.Board[0]))
		time.Sleep(200 * time.Millisecond)
	}
}
