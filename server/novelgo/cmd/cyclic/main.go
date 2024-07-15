package main

import (
	"errors"
	"fmt"
	"novelgo/internal/pkg/cyclic"
)

func main() {
	fmt.Println("cyclic go game")
	err := run()
	if err != nil {
		fmt.Printf("Program exited with error: %v\n", err)
	}
}

func run() error {
	fmt.Print("Enter board height and width, separated by space: ")
	var h, w int
	fmt.Scan(&h, &w)
	fmt.Printf("heigth = %d, width = %d\n", h, w)

	if h <= 0 || w <= 0 {
		return errors.New("Invalid board size")
	}
	b := cyclic.NewBoard(h, w)

	round := 0
	for {
		var r, c int
		var color cyclic.GridPointState
		var e error
		for {
			fmt.Print("Enter row and col coordinates, separated by space")
			fmt.Println("(enter any negative value to quit):")
			fmt.Scan(&r, &c)
			if r < 0 || c < 0 {
				return errors.New("Invalid coordinate")
			}
			fmt.Printf("row = %d, col = %d\n", r, c)
			rr := r % h
			cr := c % w
			fmt.Printf("row rounded = %d, col rounded = %d\n", rr, cr)
			if round%2 == 0 {
				color = cyclic.Black
			} else {
				color = cyclic.White
			}
			e = b.Put(rr, cr, color)
			if e == nil {
				break
			} else {
				fmt.Printf("Error putting on board: %v\n", e)
			}
		}
		fmt.Printf("board:\n")
		b.Print()
		round++
	}
}
