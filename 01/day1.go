package main

import (
	"fmt"
	"strconv"

	utils "github.com/cjavad/advent_of_code_2022"
)

func main() {
	fmt.Println("Part 1: ", part1())
	fmt.Println("Part 2: ", part2(3))
}

func part1() int64 {
	highestTotal := int64(0)
	inventories := parseElvesInventory(utils.ReadInput("input.txt"))
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
	inventories := parseElvesInventory(utils.ReadInput("input.txt"))
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
		sumArray[i] = utils.Sum(input[i])
	}
	return sumArray
}
