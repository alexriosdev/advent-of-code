package main

import (
	"advent-of-code/utils"
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("2025/day05/input.txt")
	fmt.Println("2025 Day 05 Solution")
	fmt.Printf("Part 1: %v\n", part1(input))
	fmt.Printf("Part 2: %v\n", part2(input))
}

func part1(input []byte) int {
	sections := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	ingredients := []int{}
	for _, s := range strings.Split(sections[1], "\n") {
		ingredients = append(ingredients, utils.StrToInt(s))
	}
	count := 0
	for _, ingredient := range ingredients {
		for _, s := range strings.Split(sections[0], "\n") {
			split := strings.Split(s, "-")
			start, end := utils.StrToInt(split[0]), utils.StrToInt(split[1])
			if start <= ingredient && ingredient <= end {
				count++
				break
			}
		}
	}
	return count
}

func part2(input []byte) int {
	return len(input)
}
