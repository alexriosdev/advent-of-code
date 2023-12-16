package main

import (
	"advent-of-code/utils"
	"fmt"
)

func main() {
	lines, _ := utils.ReadLines("input.txt")
	fmt.Println("2023 Day 15 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines))
}

func part1(lines []string) int {
	result := 0
	for _, line := range lines {
		val := 0
		for _, c := range line {
			if c == ',' {
				result += val
				val = 0
				continue
			}
			val = applyHash(val, c)
		}
		result += val
	}
	return result
}

func applyHash(val int, c rune) int {
	return ((val + int(c)) * 17) % 256
}
