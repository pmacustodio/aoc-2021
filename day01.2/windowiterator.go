package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type WindowIterator struct {
	file       *os.File
	scanner    *bufio.Scanner
	window     []int
	windowSize int
}

func NewWindowIterator(windowSize int, inputFileName string) *WindowIterator {
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	wi := WindowIterator{
		file:       file,
		scanner:    bufio.NewScanner(file),
		windowSize: windowSize,
	}

	return &wi
}

func (wi *WindowIterator) Next() (value int64, eof bool) {
	if len(wi.window) < 1 {
		for len(wi.window) < wi.windowSize {
			if !wi.scanner.Scan() {
				return 0, true
			}
			singleValue, err := strconv.Atoi(wi.scanner.Text())
			if err != nil {
				log.Fatal(err)
			}
			wi.window = append(wi.window, singleValue)
		}
		return wi.sum(), false
	}

	if !wi.scanner.Scan() {
		return 0, true
	}
	singleValue, err := strconv.Atoi(wi.scanner.Text())
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < wi.windowSize-1; i++ {
		wi.window[i] = wi.window[i+1]
	}
	wi.window[wi.windowSize-1] = singleValue

	return wi.sum(), false
}

func (wi *WindowIterator) sum() int64 {
	s := int64(0)
	for _, v := range wi.window {
		s += int64(v)
	}
	return s
}

func (wi *WindowIterator) Close() {
	err := wi.file.Close()
	if err != nil {
		log.Fatal(err)
	}
}
