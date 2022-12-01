package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Part 1: ", part1())
	fmt.Println("Part 2: ", part2(3))
}

func part1() int64 {
	highestTotal := int64(0)
	inventories := parseElvesInventory(readInput("input.txt"))
	sumOfInventories := getSumArray(inventories)

	for _, total := range sumOfInventories {
		if total > highestTotal {
			highestTotal = int64(total)
		}
	}

	return highestTotal
}

func part2(n int) int64 {
	// Finds n highest numbers in the input
	inventories := parseElvesInventory(readInput("input.txt"))
	sumOfInventories := getSumArray(inventories)
	// Sort sumOfInventories
	// Return the first n numbers

	for i := 0; i < len(sumOfInventories); i++ {
		for j := 0; j < len(sumOfInventories); j++ {
			if sumOfInventories[i] > sumOfInventories[j] {
				sumOfInventories[i], sumOfInventories[j] = sumOfInventories[j], sumOfInventories[i]
			}
		}
	}

	return sum(sumOfInventories[:n])
}

func readInput(inputfile string) []string {
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

func parseElvesInventory(input []string) map[int][]int64 {
	elvesInventory := make(map[int][]int64)
	inventoryIndex := 0
	for _, line := range input {
		if elvesInventory[inventoryIndex] == nil {
			elvesInventory[inventoryIndex] = make([]int64, 0)
		}

		if line == "" {
			inventoryIndex++
			continue
		}

		// Add the new item to the inventory
		item, err := strconv.ParseInt(line, 10, 64)

		if err != nil {
			continue
		}

		elvesInventory[inventoryIndex] = append(elvesInventory[inventoryIndex], item)

	}
	return elvesInventory
}

func getSumArray(input map[int][]int64) []int64 {
	sumArray := make([]int64, len(input))
	for i := 0; i < len(input); i++ {
		sumArray[i] = sum(input[i])
	}
	return sumArray
}

func sum(input []int64) int64 {
	sum := int64(0)
	for _, item := range input {
		sum += item
	}
	return sum
}
