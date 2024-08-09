package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("cyclic go game")
	input_file := flag.String("input_file", "", "path to input file")
	flag.Parse()
	var err error
	if len(*input_file) > 0 {
		fmt.Printf("Using input file: %s\n", *input_file)
		err = run(input_file)
	} else {
		err = runInteractive()
	}
	if err != nil {
		fmt.Printf("Program exited with error: %v\n", err)
	}
}
