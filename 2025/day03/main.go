package main

import (
	"advent-of-code/utils"
	"fmt"
)

func main() {
	lines, _ := utils.ReadLines("2025/day03/input.txt")
	fmt.Println("2025 Day 03 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines))
}

func part1(lines []string) int {
	sum := 0
	for _, line := range lines {
		maxVal := 0
		for i := 0; i < len(line); i++ {
			for j := i + 1; j < len(line); j++ {
				a := int(line[i] - '0')
				b := int(line[j] - '0')
				maxVal = max(maxVal, (a*10)+b)
			}
		}
		sum += maxVal
	}
	return sum
}
