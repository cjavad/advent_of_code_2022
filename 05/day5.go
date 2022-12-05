package main

import (
	"fmt"
	"strings"

	utils "github.com/cjavad/advent_of_code_2022"
)

func main() {
	s1 := part_x(1)
	s2 := part_x(2)
	fmt.Println("Part 1:", get_top_row(s1))
	pretty_print_stack(s1)
	fmt.Println("Part 2:", get_top_row(s2))
	pretty_print_stack(s2)
}

func part_x(x int) map[int32][]int32 {
	lines := utils.ReadInput("input.txt")
	// Find index of empty line
	index := 0
	for i, line := range lines {
		if line == "" {
			index = i
			break
		}
	}
	// Parse stack
	stack := parse_stack(lines[:index])
	pretty_print_stack(stack)
	// Parse moves
	moves := parse_moves(lines[index+1:])
	// Execute moves
	for _, move := range moves {
		fmt.Printf("move %d from %d to %d\n", move[0], move[1], move[2])
		stack = execute_move(stack, move, x == 2)
		pretty_print_stack(stack)
		fmt.Println()
	}
	return stack
}

func get_top_row(stack map[int32][]int32) string {
	r := make([]string, len(stack))
	for key, value := range stack {
		r[int(key)-1] = string(value[0])
	}
	return strings.Join(r, "")
}

func pretty_print_stack(stack map[int32][]int32) {
	// Output each row as [char(c)]
	lines := make([]string, len(stack))
	for key, value := range stack {
		line := fmt.Sprintf("%d", key)
		for _, v := range value {
			line += fmt.Sprintf(" [%c] ", v)
		}
		// Order lines by key
		lines[key-1] = line
	}
	// Print lines
	for _, line := range lines {
		fmt.Println(line)
	}
}

func execute_move(stack map[int32][]int32, move []int32, keep_order bool) map[int32][]int32 {
	if keep_order {
		stack[move[2]] = append(append([]int32(nil), stack[move[1]][:int(move[0])]...), stack[move[2]]...)
		stack[move[1]] = stack[move[1]][int(move[0]):]
	} else {
		// Move move[0] elements from stack[move[1]] to stack[move[2]]
		for i := int32(0); i < move[0]; i++ {
			if len(stack[move[1]]) < 1 {
				continue
			}
			stack[move[2]] = append([]int32{stack[move[1]][0]}, stack[move[2]]...)
			// Remove element from stack[move[1]]
			stack[move[1]] = stack[move[1]][1:]
			// Prepend element to stack[move[2]]
		}
	}
	return stack
}

func parse_stack(input []string) map[int32][]rune {
	r := make(map[int32][]rune)
	for _, line := range input {
		arr := []rune(line)
		for j, char := range arr {
			index := int32(j/4) + 1
			if char == '[' && arr[j+2] == ']' {
				r[index] = append(r[index], arr[j+1])
				index++
			}
		}
	}
	return r
}

func parse_moves(input []string) [][]int32 {
	r := make([][]int32, 0)
	for _, line := range input {
		r = append(r, parse_ints([]rune(line)))
	}
	return r
}

func parse_ints(input []int32) []int32 {
	r := make([]int32, 3)
	i := 0
	for j, char := range input {
		if char_is_digit(char) {
			if i == 0 {
				r = append(r, 0)
			}
			r[i] = r[i]*10 + char - '0'
		} else if j > 0 && char_is_digit(input[j-1]) {
			i++
		}
	}
	return r
}

func char_is_digit(char int32) bool {
	return char >= '0' && char <= '9'
}
