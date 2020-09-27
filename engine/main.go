package main

import (
	"fmt"
	"time"

	"github.com/ro-tex/GameOfLife/engine/world"
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
	w := world.NewFromSeed(Glider)
	for i := 0; i < 1000; i++ {
		fmt.Println("\n\n\n\n\n\n\n\n\n\n\n")
		w.Print()
		w.NextGen()
		time.Sleep(200 * time.Millisecond)
	}
}
