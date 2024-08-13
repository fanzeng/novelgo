package cyclic

import (
	"errors"
	"fmt"
)

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
	Empty
	Black
	White
)

const colorRed = "\033[0;31m"
const colorNone = "\033[0m"

type GridPoint struct {
	State GridPointState
}

func (b Board) Print() {
	fmt.Print("  ")
	for c := 0; c < b.Width; c++ {
		fmt.Printf("%d ", c)
	}
	fmt.Println()
	for r := 0; r < b.Height; r++ {
		fmt.Printf("%d ", r)
		for c := 0; c < b.Width; c++ {
			str := colorRed + "+ " + colorNone
			switch b.GridPointStates[r][c] {
			case Black:
				str = "\u23FA "
			case White:
				str = "\u25EF "
			}
			fmt.Printf("%s", str)
		}
		fmt.Printf("%d\n", r)
	}
	fmt.Print("  ")
	for c := 0; c < b.Width; c++ {
		fmt.Printf("%d ", c)
	}
	fmt.Println()
}

func (b Board) Put(r, c int, color GridPointState) error {
	if b.GridPointStates[r][c] == Undefined || b.GridPointStates[r][c] == Empty {
		b.GridPointStates[r][c] = color
	} else {
		return errors.New("Attempt to put on occupied grid point")
	}
	return b.update(r, c, color)
}

func (b Board) getOpponentColor(color GridPointState) GridPointState {
	if color == Black {
		return White
	}
	if color == White {
		return Black
	}
	return color
}

func (b Board) getNeighbors(r, c int) [][]int {
	var coords [][]int
	if r > 0 { // TODO: make circular
		coords = append(coords, []int{r - 1, c})
	}
	if c > 0 {
		coords = append(coords, []int{r, c - 1})
	}
	if r+1 < b.Height {
		coords = append(coords, []int{r + 1, c})
	}
	if c+1 < b.Width {
		coords = append(coords, []int{r, c + 1})
	}
	return coords
}

func (b Board) update(r, c int, color GridPointState) error {
	neighbors := b.getNeighbors(r, c)
	fmt.Printf("neighbors = %v\n", neighbors)
	// TODO: check suicide
	for _, neighbor := range neighbors {
		nr := neighbor[0]
		nc := neighbor[1]
		if b.GridPointStates[nr][nc] == b.getOpponentColor(color) {
			fmt.Printf("opponent = %v", neighbor)
			alive, err := b.checkMarkAlive(r, c)
			if err != nil {
				return err
			}
			fmt.Printf("alive = %v\n", alive)
		}
	}
	return nil
}

func (b Board) checkMarkAlive(r, c int) (bool, error) {
	return true, nil
}
