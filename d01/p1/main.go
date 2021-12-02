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
	var last, count int
	s := bufio.NewScanner(f)
	for s.Scan() {
		// Parse value into integer type
		val, err := strconv.Atoi(s.Text())
		if err != nil {
			return fmt.Errorf("parse input: %#v: %w", val, err)
		}

		// Increment counter if not first
		// NOTE: Assumes that input will never be 0
		if last != 0 && val > last {
			count++
		}
		last = val
	}
	if err := s.Err(); err != nil {
		return fmt.Errorf("scan input: %w", err)
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
