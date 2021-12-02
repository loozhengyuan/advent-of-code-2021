package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func run() error {
	// Read input from file
	f, err := os.Open("d01/input.txt")
	if err != nil {
		return fmt.Errorf("open file: %w", err)
	}
	defer f.Close()

	// Scan each line
	data := make([]int, 0)
	s := bufio.NewScanner(f)
	for s.Scan() {
		// Parse value into integer type
		val, err := strconv.Atoi(s.Text())
		if err != nil {
			return fmt.Errorf("parse input: %#v: %w", val, err)
		}
		data = append(data, val)
	}
	if err := s.Err(); err != nil {
		return fmt.Errorf("scan input: %w", err)
	}

	// Compute sliding windows
	var last, count int
	for i := 2; i < len(data); i++ {
		agg := data[i-2] + data[i-1] + data[i]
		if i > 2 && agg > last {
			count++
		}
		last = agg
	}

	fmt.Println(count)
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}
