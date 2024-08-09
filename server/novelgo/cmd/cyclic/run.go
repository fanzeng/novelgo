package main

import (
	"bufio"
	"errors"
	"fmt"
	"novelgo/internal/pkg/cyclic"
	"os"
	"strconv"
	"strings"
)

func run(input_file *string) error {
	file, err := os.Open(*input_file)
	if err != nil {
		return nil
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	if !sc.Scan() {
		return errors.New("Fail to read version of input file")
	}
	version := sc.Text()
	fmt.Printf("Version %s\n", version)
	if !sc.Scan() {
		return errors.New("Fail to board size")
	}
	boardSize := sc.Text()
	fmt.Printf("Board size %s\n", boardSize)

	h, w, err := parseLine(boardSize)
	b := cyclic.NewBoard(h, w)

	round := 0
	for sc.Scan() {
		var color cyclic.GridPointState
		line := sc.Text()
		if len(line) <= 0 {
			continue
		}
		r, c, e := parseLine(line)
		if e != nil {
			return e
		}
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
		if e != nil {
			fmt.Printf("Error putting on board: %v\n", e)
		}
		b.Print()
		round++
	}
	return nil
}

func runInteractive() error {
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

func parseLine(line string) (int, int, error) {
	parts := strings.Split(line, " ")
	if len(parts) != 2 {
		return -1, -1, errors.New("Expect 2 integer numbers on each line")
	}
	a, err := strconv.Atoi(parts[0])
	if err != nil {
		return -1, -1, fmt.Errorf("Fail to parse 1st int: %s", parts[0])
	}
	b, err := strconv.Atoi(parts[1])
	if err != nil {
		return -1, -1, fmt.Errorf("Fail to parse 2nd int: %s", parts[1])
	}
	return a, b, nil
}
