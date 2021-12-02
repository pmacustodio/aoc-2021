package main

import (
	"fmt"
)

func main() {
	wi := NewWindowIterator(3, "input.txt")
	defer wi.Close()
	result := day012Puzzle(wi)
	fmt.Printf("Found %d increments", result)
}

func day012Puzzle(wi *WindowIterator) int {
	lastValueRead, eof := wi.Next()
	if eof {
		return 0
	}

	increments := 0
	for {
		currentValueRead, eof := wi.Next()
		if eof {
			break
		}

		if currentValueRead > lastValueRead {
			increments++
		}
		lastValueRead = currentValueRead
	}
	return increments
}
