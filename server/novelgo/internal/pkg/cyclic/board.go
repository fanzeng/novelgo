package cyclic

import (
	"errors"
	"fmt"
)

type Board struct {
	Height          int
	Width           int
	GridPointStates [][]GridPointState
	CyclicLogic     bool
}

func NewBoard(h, w int, cyclicLogic bool) *Board {
	b := make([][]GridPointState, h)
	for i := range h {
		b[i] = make([]GridPointState, w)
	}
	return &Board{h, w, b, cyclicLogic}
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
	for c := range b.Width {
		fmt.Printf("%d ", c)
	}
	fmt.Println()
	for r := range b.Height {
		fmt.Printf("%d ", r)
		for c := range b.Width {
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
	for c := range b.Width {
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
	if r > 0 {
		coords = append(coords, []int{r - 1, c})
	} else if r == 0 && b.CyclicLogic && b.Height > 1 {
		coords = append(coords, []int{b.Height - 1, c})
	}
	if c > 0 {
		coords = append(coords, []int{r, c - 1})
	} else if c == 0 && b.CyclicLogic && b.Width > 1 {
		coords = append(coords, []int{r, b.Width - 1})
	}
	if r+1 < b.Height {
		coords = append(coords, []int{r + 1, c})
	} else if r+1 == b.Height && b.CyclicLogic && b.Height > 1 {
		coords = append(coords, []int{0, c})
	}
	if c+1 < b.Width {
		coords = append(coords, []int{r, c + 1})
	} else if c+1 == b.Width && b.CyclicLogic && b.Width > 1 {
		coords = append(coords, []int{r, 0})
	}
	return coords
}

func (b Board) update(r, c int, color GridPointState) error {
	neighbors := b.getNeighbors(r, c)
	fmt.Printf("neighbors = %v\n", neighbors)
	// TODO: check suicide including whether the new stone can be alive itself
	isKill := false
	for _, n := range neighbors {
		if b.GridPointStates[n[0]][n[1]] == b.getOpponentColor(color) {
			fmt.Printf("opponent = %v\n", n)
			alive, err := b.checkMarkAlive(n[0], n[1], false)
			if err != nil {
				return err
			}
			fmt.Printf("alive = %v\n", alive)
			if !alive {
				isKill = true
			}
		}
	}
	if !isKill {
		// Check for suicide if the move is not a kill
		if alive, _ := b.checkMarkAlive(r, c, true); alive == false {
			// Roll back the suicidal move
			b.GridPointStates[r][c] = Empty
			return errors.New("suicide not allowed")
		}
	}
	return nil
}

func (b Board) checkMarkAlive(r, c int, checkOnly bool) (bool, error) {
	visited := make([][]bool, b.Height)
	for i := range b.Height {
		visited[i] = make([]bool, b.Width)
	}
	cluster := b.getCluster(r, c, visited)
	fmt.Printf("cluster = %v\n", cluster)
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
	if !checkOnly && !isAlive {
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

func (b Board) GetGridPointsAsArray() []int {
	a := make([]int, b.Height*b.Width)
	for r := range b.Height {
		fmt.Printf("%d ", r)
		for c := range b.Width {
			a[r*b.Width + c] = int(b.GridPointStates[r][c])
		}
	}
	return a
}
