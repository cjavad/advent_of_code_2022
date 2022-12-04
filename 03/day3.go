package main

import (
	"fmt"

	utils "github.com/cjavad/advent_of_code_2022"
	"golang.org/x/exp/slices"
)

func main() {
	part1()
	part2()
}

func part1() {
	rucksack := parse_rucksack(utils.ReadInput("input.txt"))
	var sum int64

	for _, compartments := range rucksack {
		intersection := unique_intersect(compartments[0], compartments[1])
		sum += utils.Sum(intersection)
	}

	fmt.Println(sum)
}

func part2() {
	rucksack := parse_rucksack(utils.ReadInput("input.txt"))

	// Three way unique intersection
	var sum int64
	for i := 0; i < len(rucksack); i++ {
		if (i+1)%3 == 0 && i > 0 {
			// For each third combination, find the intersection of the three
			sum += utils.Sum(three_way_unique_intersection(rucksack[i], rucksack[i-1], rucksack[i-2]))
		}
	}

	fmt.Println(sum)
}

func normalize_runes(runes []rune) []int64 {
	var result []int64
	for _, r := range runes {
		if r > 90 {
			result = append(result, int64(r)-96)
		} else {
			result = append(result, int64(r)-38)
		}
	}
	return result
}

func parse_rucksack(input []string) [][2][]int64 {
	rucksack := make([][2][]int64, 0)
	for _, line := range input {
		if line == "" {
			continue
		}

		// Convert line to runes
		runes_line := []rune(line)
		// Split the line into two parts
		compartments := [2][]int64{normalize_runes(runes_line[:len(runes_line)/2]), normalize_runes(runes_line[len(runes_line)/2:])}
		// Add the compartments to the rucksack
		rucksack = append(rucksack, compartments)
	}
	return rucksack
}

func unique_intersect(a, b []int64) []int64 {
	m := make(map[int64]bool)
	for _, v := range a {
		m[v] = true
	}

	var r []int64
	for _, v := range b {
		if m[v] && !slices.Contains(r, v) {
			r = append(r, v)
		}
	}
	return r
}

func three_way_unique_intersection(a, b, c [2][]int64) []int64 {
	m := unique_intersect(combine_rucksack(a), combine_rucksack(b))
	n := unique_intersect(m, combine_rucksack(c))
	return n
}

func combine_rucksack(c [2][]int64) []int64 {
	return append(c[0], c[1]...)
}
