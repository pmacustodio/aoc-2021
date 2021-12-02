package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	inputFileName := "input.txt"
	result := day011Puzzle(inputFileName)
	fmt.Printf("Found %d increments", result)
}

func day011Puzzle(inputFileName string) int {
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		log.Fatal("empty file")
	}

	lastValueRead, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	increments := 0
	for scanner.Scan() {
		currentValueRead, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		if currentValueRead > lastValueRead {
			increments++
		}
		lastValueRead = currentValueRead
	}

	return increments
}
