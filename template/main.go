package main

import (
	"advent-of-code/utils"
	"fmt"
)

func main() {
	lines, _ := utils.ReadLines("202#/day##/input.txt")
	fmt.Println("202# Day ## Solution")
	fmt.Printf("Part 1: %v\n", part1(lines))
	fmt.Printf("Part 2: %v\n", part2(lines))
}

func part1(lines []string) int {
	return len(lines)
}

func part2(lines []string) int {
	return len(lines)
}
