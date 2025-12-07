package main

import (
	"advent-of-code/utils"
	"fmt"
	"strings"
)

func main() {
	lines, _ := utils.ReadLines("2025/day06/input.txt")
	fmt.Println("2025 Day 06 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines))
	fmt.Printf("Part 2: %v\n", part2(lines))
}

func part1(lines []string) int {
	problems := [][]int{}
	for i := 0; i < len(lines)-1; i++ {
		nums := []int{}
		for _, s := range strings.Fields(lines[i]) {
			nums = append(nums, utils.StrToInt(s))
		}
		problems = append(problems, nums)
	}
	operands := strings.Fields(lines[len(lines)-1])
	rows, cols := len(problems), len(problems[0])
	result := 0
	for c := 0; c < cols; c++ {
		mult, sum := 1, 0
		for r := 0; r < rows; r++ {
			if operands[c] == "*" {
				mult *= problems[r][c]
			} else {
				sum += problems[r][c]
			}
		}
		if mult > 1 {
			result += mult
		}
		result += sum
	}
	return result
}

func part2(lines []string) int {
	return len(lines)
}
