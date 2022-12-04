package main

import (
	"fmt"
	"strconv"
	"strings"

	utils "github.com/cjavad/advent_of_code_2022"
)

func main() {
	part1()
	part2()
}

func part1() {
	sum := 0
	lines := utils.ReadInput("input.txt")
	pair_ranges := parse_pair_ranges(lines)
	for _, pair_range := range pair_ranges {
		if range_is_contained(pair_range[0], pair_range[1]) {
			sum += 1
		}
	}

	fmt.Println(sum)
}

func part2() {
	sum := 0
	lines := utils.ReadInput("input.txt")
	pair_ranges := parse_pair_ranges(lines)
	for _, pair_range := range pair_ranges {
		if range_is_overlapping(pair_range[0], pair_range[1]) {
			sum += 1
		}
	}

	fmt.Println(sum)
}

func parse_pair_ranges(input []string) [][][2]int {
	pair_ranges := make([][][2]int, 0)
	for _, line := range input {
		if line == "" {
			continue
		}

		pairs := strings.Split(line, ",")
		pair_range := make([][2]int, 0)
		for _, pair := range pairs {
			pair_range = append(pair_range, parse_range(pair))
		}
		pair_ranges = append(pair_ranges, pair_range)
	}
	return pair_ranges
}

func parse_range(input string) [2]int {
	input = strings.Trim(input, " ")
	pair := strings.Split(input, "-")
	lower, _ := strconv.Atoi(pair[0])
	upper, _ := strconv.Atoi(pair[1])
	return [2]int{lower, upper}
}

func range_is_contained(a, b [2]int) bool {
	return a[0] >= b[0] && a[1] <= b[1] || b[0] >= a[0] && b[1] <= a[1]
}

func range_is_overlapping(a, b [2]int) bool {
	return a[0] <= b[0] && a[1] >= b[0] || b[0] <= a[0] && b[1] >= a[0]
}
