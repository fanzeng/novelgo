package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"novelgo/internal/pkg/cyclic"
	"os"
	"strings"
	"time"
)

var version = "v0.1.0"

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
	v := sc.Text()
	fmt.Printf("gameplay version %s\n", v)
	fmt.Printf("current version %s\n", version)
	if !sc.Scan() {
		return errors.New("Fail to read board size")
	}
	boardSize := sc.Text()
	fmt.Printf("Board size %s\n", boardSize)

	h, w, err := parseLine(boardSize)
	if err != nil {
		return err
	}
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
	currentTime := time.Now().Format("20060102_150405")
	filename := fmt.Sprintf("%s.txt", currentTime)

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error creating gameplay file:", err)
		return err
	}
	defer file.Close()
	file.WriteString(version + "\n")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter board height and width, separated by space: ")
	var h, w int
	line, e := reader.ReadString('\n')
	if e != nil {
		return e
	}
	fmt.Sscanf(line, "%d %d", &h, &w)
	fmt.Printf("heigth = %d, width = %d\n", h, w)

	if h <= 0 || w <= 0 {
		return errors.New("Invalid board size")
	}
	line = fmt.Sprintf("%d %d\n", h, w)
	if _, e = file.WriteString(line); e != nil {
		fmt.Println("Error writing to file:", e)
		return e
	}
	b := cyclic.NewBoard(h, w)
	fmt.Printf("board:\n")
	b.Print()

	round := 0
	for {
		var color cyclic.GridPointState
		for {
			fmt.Print("Enter row and col coordinates, separated by space")
			fmt.Println("(enter any negative value to quit):")

			line, e := reader.ReadString('\n')
			if e != nil {
				return e
			}
			fmt.Print(line)
			r, c, e := parseLine(line)
			if errors.Is(e, io.EOF) {
				fmt.Println("Did you press [ENTER] too soon?")
				continue
			} else if e != nil {
				return e
			}
			if r < 0 || c < 0 {
				fmt.Println("bye~")
				return nil
			}
			fmt.Printf("row = %d, col = %d\n", r, c)
			line = fmt.Sprintf("%d %d\n", r, c)
			if _, e = file.WriteString(line); e != nil {
				fmt.Println("Error writing to file:", e)
				return e
			}
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
	var a, b int
	var extra string
	line = strings.TrimSpace(line)
	n, err := fmt.Sscanf(line, "%d %d", &a, &b)
	if err != nil {
		return -1, -1, err
	}
	n, err = fmt.Sscanf(line, "%d %d%s", &a, &b, &extra)
	if n != 2 {
		return -1, -1, errors.New("Expect 2 integer numbers on each line")
	}
	return a, b, nil
}
