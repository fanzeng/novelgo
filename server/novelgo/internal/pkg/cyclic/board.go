package cyclic

import "fmt"

type Board struct {
	Height          int
	Width           int
	GridPointStates [][]GridPointState
}

func NewBoard(h, w int) *Board {
	b := make([][]GridPointState, h)
	for i := range h {
		b[i] = make([]GridPointState, w)
	}
	return &Board{h, w, b}
}

type GridPointState int

const (
	Undefined GridPointState = iota
	Black
	White
)

type GridPoint struct {
	State GridPointState
}

func (b Board) Print() {
	b.GridPointStates[0][0] = Black
	b.GridPointStates[0][1] = White
	for r := 0; r < b.Height; r++ {
		for c := 0; c < b.Width; c++ {
			str := "+ "
			switch b.GridPointStates[r][c] {
			case Black:
				str = "\u23FA "
			case White:
				str = "\u25EF "
			}
			fmt.Printf("%s", str)
		}
		fmt.Println()
	}
	fmt.Println()
}
