package main

import (
	"time"
)

var (
	Blinker = Shape{[][]byte{
		{0, 0, 0},
		{1, 1, 1},
		{0, 0, 0},
	}}

	GliderDR = Shape{[][]byte{
		{1, 0, 0},
		{0, 1, 1},
		{1, 1, 0},
	}}
	GliderDL = Shape{[][]byte{
		{1, 0, 1},
		{1, 1, 0},
		{0, 1, 0},
	}}
	GliderUL = Shape{[][]byte{
		{0, 1, 1},
		{1, 1, 0},
		{0, 0, 1},
	}}
	GliderUR = Shape{[][]byte{
		{0, 1, 0},
		{0, 1, 1},
		{1, 0, 1},
	}}
	GliderSquare = Shape{[][]byte{
		{0, 1, 1, 0, 0, 0, 0, 1, 0},
		{1, 1, 0, 0, 0, 0, 0, 1, 1},
		{0, 0, 1, 0, 0, 0, 1, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 0, 1, 0, 0, 0, 1, 0, 0},
		{1, 1, 0, 0, 0, 0, 0, 1, 1},
		{0, 1, 0, 0, 0, 0, 1, 1, 0},
	}}
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
	w := NewWorldFromSeed(GliderSquare.Board)
	w.Pad()
	w.Pad()
	w.Pad()
	for i := 0; i < 1000; i++ {
		w.Print()
		w.NextGen()
		time.Sleep(200 * time.Millisecond)
	}
}
