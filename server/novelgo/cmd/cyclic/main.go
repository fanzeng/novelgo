package main

import (
	"fmt"
	"novelgo/internal/pkg/cyclic"
)

func main() {
	fmt.Println("cyclic go game")
	run()
}

func run() {
	b := cyclic.NewBoard(5, 5)
	fmt.Print("Enter row and col coordinates, separated by space: ")
	var r, c int
	fmt.Scan(&r, &c)
	fmt.Printf("row = %d, col = %d\n", r, c)
	fmt.Printf("board:\n")
	b.Print()
}
