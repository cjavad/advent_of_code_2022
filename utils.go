package utils

// Global utility functions for Advent of Code 2022 written in Go

import (
	"bufio"
	"os"
)

// ReadInput reads the input file and returns a slice of strings
func ReadInput(inputfile string) []string {
	// Read the input file
	input, err := os.Open(inputfile)
	if err != nil {
		panic(err)
	}
	defer input.Close()
	var lines []string
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func Sum(input []int64) int64 {
	sum := int64(0)
	for _, item := range input {
		sum += item
	}
	return sum
}
