package cyclic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoardCyclicLogicFalse(t *testing.T) {
	b := NewBoard(3, 3, false)
	moves := [][]int{
		{0, 0}, {0, 1}, {1, 1}, {1, 0}, {2, 1}, {2, 2}, {2, 0}, {0, 2}, {1, 2},
	}
	color := Black
	for _, coord := range moves {
		b.Put(coord[0], coord[1], color)
		color = b.getOpponentColor(color)
		b.Print()
	}
	want := [][]GridPointState{
		{Empty, White, White},
		{White, Black, Black},
		{Black, Black, Empty},
	}
	assert.Equal(t, b.GridPointStates, want, "incorrect gameplay outcome")
}

func TestBoardCyclicLogicTrue(t *testing.T) {
	b := NewBoard(3, 3, true)
	moves := [][]int{
		{0, 0}, {0, 1}, {1, 1}, {1, 0}, {2, 1}, {2, 2}, {2, 0}, {0, 2}, {1, 2},
	}
	color := Black
	for _, coord := range moves {
		b.Put(coord[0], coord[1], color)
		color = b.getOpponentColor(color)
		b.Print()
	}
	want := [][]GridPointState{
		{Black, Empty, Empty},
		{Empty, Black, Black},
		{Black, Black, Empty},
	}
	assert.Equal(t, b.GridPointStates, want, "incorrect gameplay outcome")
}
