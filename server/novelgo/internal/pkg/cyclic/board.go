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
	// TODO: check suicide including whether the new stone can be alive itself
	for _, n := range neighbors {
		if b.GridPointStates[n[0]][n[1]] == b.getOpponentColor(color) {
			fmt.Printf("opponent = %v", n)
			alive, err := b.checkMarkAlive(n[0], n[1])
			if err != nil {
				return err
			}
			fmt.Printf("alive = %v\n", alive)
		}
	}
	return nil
}

func (b Board) checkMarkAlive(r, c int) (bool, error) {
	visited := make([][]bool, b.Height)
	for i := range b.Height {
		visited[i] = make([]bool, b.Width)
	}
	cluster := b.getCluster(r, c, visited)
	fmt.Printf("cluster = %v\n", cluster)
	// opponentColor := b.getOpponentColor(b.GridPointStates[r][c])
	// fmt.Printf("opponentColor = %v\n", opponentColor)
	isAlive := false
	for _, p := range cluster {
		for _, n := range b.getNeighbors(p[0], p[1]) {
			if b.GridPointStates[n[0]][n[1]] == Undefined || b.GridPointStates[n[0]][n[1]] == Empty {
				isAlive = true
				break
			}
		}
		if isAlive {
			break
		}
	}
	if !isAlive {
		for _, p := range cluster {
			b.GridPointStates[p[0]][p[1]] = Empty
		}
	}
	return isAlive, nil
}

func (b Board) getCluster(r, c int, visited [][]bool) [][]int {
	var cluster [][]int
	cluster = append(cluster, []int{r, c})
	visited[r][c] = true
	color := b.GridPointStates[r][c]
	for _, n := range b.getNeighbors(r, c) {
		if visited[n[0]][n[1]] {
			continue
		}
		if b.GridPointStates[n[0]][n[1]] == color {
			cluster = append(cluster, n)
			visited[n[0]][n[1]] = true
			nn := b.getCluster(n[0], n[1], visited)
			cluster = append(cluster, nn...)
		}
	}
	return cluster
}
