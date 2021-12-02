package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func run() error {
	// Read input from file
	f, err := os.Open("d02/input.txt")
	if err != nil {
		return fmt.Errorf("open file: %w", err)
	}
	defer f.Close()

	// Scan each line
	var horizontal, depth int
	s := bufio.NewScanner(f)
	for s.Scan() {
		// Parse command and value
		line := strings.Fields(s.Text())
		if len(line) != 2 {
			return fmt.Errorf("parse line: invalid format: %#v", line)
		}
		cmd := line[0]
		val, err := strconv.Atoi(line[1])
		if err != nil {
			return fmt.Errorf("parse input: %#v: %w", val, err)
		}

		// Increment/decrement position counters
		switch cmd {
		case "forward":
			horizontal += val
		case "down":
			depth += val
		case "up":
			depth -= val
		default:
			return fmt.Errorf("unexpected command: %v", cmd)
		}
	}
	if err := s.Err(); err != nil {
		return fmt.Errorf("scan input: %w", err)
	}

	fmt.Println(horizontal * depth)
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}
