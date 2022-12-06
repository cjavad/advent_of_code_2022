package main

import (
	"fmt"

	utils "github.com/cjavad/advent_of_code_2022"
)

func main() {
	p1 := part_n(4)
	fmt.Print("Part 1:")
	for _, v := range p1 {
		fmt.Printf(" %d", v)
	}
	fmt.Println()

	p2 := part_n(14)
	fmt.Print("Part 1:")
	for _, v := range p2 {
		fmt.Printf(" %d", v)
	}
	fmt.Println()
}

func part_n(n int) map[int]int {
	r := make(map[int]int)
	lines := utils.ReadInput("input.txt")
	for j, line := range lines {
		for i, _ := range line {
			if is_unique_sequence([]rune(line)[i:i+n]) && len(line) > i+n {
				r[j] = i + n
				break
			}
		}
	}
	return r
}

func is_unique_sequence(seq []int32) bool {
	// Check if seq contains dubplicates
	// Return true if it does not
	seen := make(map[int32]bool)
	for _, val := range seq {
		if seen[val] {
			return false
		}
		seen[val] = true
	}
	return true
}
