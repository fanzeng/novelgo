package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("cyclic go game")
	inputFile := flag.String("input_file", "", "path to input file")
	pauseStep := flag.Bool("pause_step", false, "whether to pause between each step")
	cyclicLogic := flag.Bool("cyclic_logic", true, "whether the logic is cyclic")
	flag.Parse()
	var err error
	if len(*inputFile) > 0 {
		fmt.Printf("Using input file: %s\n", *inputFile)
		err = run(inputFile, *pauseStep, *cyclicLogic)
	} else {
		err = runInteractive(*cyclicLogic)
	}
	if err != nil {
		fmt.Printf("Program exited with error: %v\n", err)
	}
}
