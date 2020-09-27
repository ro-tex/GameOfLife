package main

import (
	"time"

	"github.com/ro-tex/GameOfLife/engine/shape"

	"github.com/ro-tex/GameOfLife/engine/world"
)

/*
TODO
	- tests :)
	- performance test on NextGen. keep previous results on disk, so you can compare
	- add support for the standard format of Golly
	- figure type that can rotate, combine, store, load Golly files.
	- Web page has API that can init the world from base64. It can also load a standard file.
*/

func main() {
	w := world.NewFromSeed(shape.GliderSquare.Board)
	w.Pad()
	w.Pad()
	w.Pad()
	for i := 0; i < 1000; i++ {
		w.Print()
		w.NextGen()
		time.Sleep(200 * time.Millisecond)
	}
}
