package main

import (
	"advent-of-code/utils"
	"fmt"
	"unicode"
)

func main() {
	lines, _ := utils.ReadLines("input.txt")
	fmt.Println("2023 Day 01 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines))
}

func part1(lines []string) int {
	result := 0
	for _, s := range lines {
		first, last := 0, 0
		for _, c := range s {
			if unicode.IsDigit(c) {
				if first == 0 {
					first = int(c - '0')
				}
				last = int(c - '0')
			}
		}
		result += (first * 10) + last
	}
	return result
}
