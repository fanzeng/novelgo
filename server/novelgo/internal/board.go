package wrapped_around_go

import "fmt"

type Board struct {
	Height int
	Width  int
}

func NewBoard() *Board {
	return &Board{2, 2}
}

type GridPoint struct {
	State int
}

func (b Board) Print() {
	for r := 0; r < b.Height; r++ {
		for c := 0; c < b.Width; c++ {
			fmt.Print("+")
		}
	}
}
