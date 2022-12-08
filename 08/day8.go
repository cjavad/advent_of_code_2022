package main

import (
	"fmt"

	utils "github.com/cjavad/advent_of_code_2022"
)

func main() {
	tree_map := parse_tree_map(utils.ReadInput("input.txt"))
	score_map := make([][]int, len(tree_map))
	for i := range score_map {
		score_map[i] = make([]int, len(tree_map[i]))
	}
	trees_visible := 0

	for i := 0; i < len(tree_map); i++ {
		for j := 0; j < len(tree_map[i]); j++ {
			t, l, r, b := true, true, true, true
			td, ld, rd, bd := 0, 0, 0, 0

			for k := i - 1; k >= 0; k-- {
				td++
				if tree_map[k][j] >= tree_map[i][j] {
					t = false
					break
				}
			}

			for k := j - 1; k >= 0; k-- {
				ld++
				if tree_map[i][k] >= tree_map[i][j] {
					l = false
					break
				}
			}

			for k := j + 1; k < len(tree_map[i]); k++ {
				rd++
				if tree_map[i][k] >= tree_map[i][j] {
					r = false
					break
				}
			}

			for k := i + 1; k < len(tree_map); k++ {
				bd++
				if tree_map[k][j] >= tree_map[i][j] {
					b = false
					break
				}
			}

			score_map[i][j] = td * ld * rd * bd

			if t || l || r || b {
				fmt.Printf("Tree at (%d, %d) is visible\n", i+1, j+1)
				trees_visible++
			} else {
				fmt.Printf("Tree at (%d, %d) is not visible\n", i+1, j+1)
			}
		}
	}

	fmt.Println("Part 1: ", trees_visible)

	// Find highest score
	max_score := 0
	for i := 0; i < len(score_map); i++ {
		for j := 0; j < len(score_map[i]); j++ {
			if score_map[i][j] > max_score {
				max_score = score_map[i][j]
			}
		}
	}
	fmt.Printf("Part 2: %d\n", max_score)
}

func parse_tree_map(input []string) [][]int {
	arr := make([][]int, len(input))
	for i, line := range input {
		arr[i] = make([]int, len(line))
		for j, c := range line {
			arr[i][j] = int(c - '0')
		}
	}
	return arr
}

/*
Unused function for boolean approach, potentially faster but you cannot extract score after processing
*/

func directional_tree_map(arr [][]int) [][][]bool {
	visible := make([][][]bool, len(arr))
	for i := range visible {
		visible[i] = make([][]bool, len(arr[i]))
		for j := range visible[i] {
			visible[i][j] = is_visible(arr, i, j)
		}
	}
	return visible
}

// Top, Left, Right, Bottom
func is_visible(arr [][]int, i, j int) []bool {
	height := arr[i][j]
	t, l, r, b := true, true, true, true
	if i-1 > 0 {
		t = arr[i-1][j] <= height
	}

	if j-1 > 0 {
		l = arr[i][j-1] <= height
	}

	if j+1 < len(arr[i]) {
		r = arr[i][j+1] <= height
	}

	if i+1 < len(arr) {
		b = arr[i+1][j] <= height
	}

	return []bool{t, l, r, b}
}
