package main

import (
	"fmt"

	utils "github.com/cjavad/advent_of_code_2022"
	"golang.org/x/exp/slices"
)

func main() {
	p1 := part1()
	p2 := part2()
	fmt.Println("Part 1: ", p1)
	fmt.Println("Part 2: ", p2)
}

func part1() int {
	lines := utils.ReadInput("input.txt")
	sum := 0
	for _, line := range lines {
		sum += score_from_string(line)
	}
	return sum
}

func part2() int {
	lines := utils.ReadInput("input.txt")
	sum := 0
	for _, line := range lines {
		sum += reverse_score_from_string(line)
	}
	return sum
}

func print_output(x int) {
	switch x {
	case 1:
		fmt.Print("Rock")
	case 2:
		fmt.Print("Paper")
	case 3:
		fmt.Print("Scissors")
	}
}

func reverse_score_from_string(s string) int {
	lookup := [][]int{
		{0, -1, 1},
		{1, 0, -1},
		{-1, 1, 0},
	}
	r := []rune(s)
	e, c := int(rune(r[0]))-64, int(rune(r[2]))-87
	idx := int(0)
	p := 3 * (c - 1)
	print_output(e)

	switch c {
	case 1:

		fmt.Print(" needs to lose against ")
		idx = slices.Index(lookup[e-1], 1)
	case 2:
		p = 3
		fmt.Print(" needs to tie against ")
		idx = slices.Index(lookup[e-1], 0)
	case 3:
		fmt.Print(" needs to win against ")
		idx = slices.Index(lookup[e-1], -1)
	default:
		panic("Invalid input")
	}

	p += idx + 1

	print_output(idx + 1)
	fmt.Printf(" for %d points\n", p)

	return p
}

func score_from_string(s string) int {
	lookup := [][]int{
		{0, -1, 1},
		{1, 0, -1},
		{-1, 1, 0},
	}

	r := []rune(s)
	e, y := int(rune(r[0]))-64, int(rune(r[2]))-87
	p := y

	print_output(e)

	switch lookup[y-1][e-1] {
	case 1:
		p += 6
		fmt.Print(" beats ")
	case -1:
		fmt.Print(" loses to ")
	case 0:
		p += 3
		fmt.Print(" ties ")
	}

	print_output(y)
	fmt.Printf(" for %d points\n", p)

	return p
}
