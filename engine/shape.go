package main

// Shape represents a section of a World board. Ideally, it would be a construct
// that does something interesting, e.g. a glider.
type Shape struct {
	Board [][]byte
}

// NewShapeFromSeed creates a new shape.
func NewShapeFromSeed(seed [][]byte) *Shape {
	return &Shape{
		Board: seed,
	}
}

// Rotate rotates the shape 90 degrees clockwise.
func (s *Shape) Rotate() {
	// make a new Board
	nbh := len(s.Board[0])
	nbw := len(s.Board)
	nb := make([][]byte, nbh)
	for i := range nb {
		nb[i] = make([]byte, nbw)
	}
	// copy over
	for h := range s.Board {
		for w := range s.Board[h] {
			nb[w][nbw-h-1] = s.Board[h][w]
		}
	}
	s.Board = nb
}
