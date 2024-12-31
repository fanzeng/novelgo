package cyclic

import (
	"fmt"
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
	assert.Equal(t, want, b.GridPointStates, "incorrect gameplay outcome")
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
	assert.Equal(t, want, b.GridPointStates, "incorrect gameplay outcome")
}

func TestBoardSuicideDetection(t *testing.T) {
	b := NewBoard(4, 4, true)
	moves := [][]int{
		{1, 1}, {0, 1}, {3, 0}, {0, 2}, {3, 1}, {1, 0}, {3, 2}, {1, 3}, {3, 3}, {2, 1}, {3, 3}, {2, 2},
	}
	color := Black
	for _, coord := range moves {
		b.Put(coord[0], coord[1], color)
		color = b.getOpponentColor(color)
		b.Print()
	}
	err := b.Put(1, 2, color)
	assert.NotNil(t, err, "expected an error but got nil")
	assert.EqualError(t, err, "suicide not allowed", err.Error())
	fmt.Print(err)
}

func TestBoardSuicideWithKillDetection(t *testing.T) {
	b := NewBoard(4, 4, true)
	moves := [][]int{
		{1, 1}, {0, 1}, {3, 0}, {0, 2}, {3, 1}, {1, 0}, {3, 2}, {1, 3}, {3, 3}, {2, 1}, {3, 3}, {2, 2}, {2, 0},
	}
	color := Black
	for _, coord := range moves {
		b.Put(coord[0], coord[1], color)
		color = b.getOpponentColor(color)
		b.Print()
	}
	err := b.Put(1, 2, color)
	b.Print()
	assert.Nil(t, err, "expected no error but got one")
	// want := [][]GridPointState{
	// 	{Empty, White, White, Empty},
	// 	{White, Empty, White, White},
	// 	{Black, White, White, Empty},
	// 	{Black, Black, Black, Black},
	// }
	// b.Print()
	// assert.Equal(t, want, b.GridPointStates, "incorrect gameplay outcome")
}
