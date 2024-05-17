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
	fmt.Printf("board:\n")
	b.Print()
}
