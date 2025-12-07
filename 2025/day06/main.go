package main

import (
	"advent-of-code/utils"
	"fmt"
	"slices"
	"strings"
	"unicode"
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
		replacer := strings.NewReplacer("x", "")
		for _, s := range strings.Fields(replacer.Replace(lines[i])) {
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
	grid := convertToGrid(lines)
	rows, cols := len(grid), len(grid[0])
	isMult, isSum := false, false
	mult, sum := 1, 0
	result := 0
	for c := cols - 1; c >= 0; c-- {
		runes := []rune{}
		for r := rows - 1; r >= 0; r-- {
			val := grid[r][c]
			if unicode.IsNumber(val) {
				runes = append(runes, val)
			}
			if val == '*' {
				isMult = true
				isSum = false
			}
			if val == '+' {
				isSum = true
				isMult = false
			}
		}
		if len(runes) == 0 {
			continue
		}
		slices.Reverse(runes)
		num := utils.StrToInt(string(runes))
		mult *= num
		sum += num
		if isMult && mult > 1 {
			result += mult
			mult, sum = 1, 0
			isMult = false
		}
		if isSum {
			result += sum
			mult, sum = 1, 0
			isSum = false
		}
	}
	return result
}

func convertToGrid(lines []string) [][]rune {
	grid := [][]rune{}
	for _, line := range lines {
		grid = append(grid, []rune(line))
	}
	return grid
}
