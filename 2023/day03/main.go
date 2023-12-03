package main

import (
	"advent-of-code/utils"
	"fmt"
	"unicode"
)

func main() {
	lines, _ := utils.ReadLines("input.txt")
	fmt.Println("2023 Day 03 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines))
}

func part1(lines []string) int {
	result := 0
	grid := convertToGrid(lines)
	for i, row := range grid {
		num := 0
		isValid := false
		for j := range row {
			c := grid[i][j]
			if unicode.IsDigit(c) {
				num = (num * 10) + int(c-'0')
				if isValid == false {
					isValid = checkNeighbors(grid, i, j)
				}
				continue
			}
			if isValid && num > 0 {
				result += num
				isValid = false
			}
			num = 0
		}
		if isValid && num > 0 {
			result += num
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

func checkNeighbors(grid [][]rune, i, j int) bool {
	isSymbol := false
	dirs := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}}
	for _, dir := range dirs {
		r := i + dir[0]
		c := j + dir[1]
		if c >= 0 && c < len(grid[i]) && r >= 0 && r < len(grid) {
			char := grid[r][c]
			if char != '.' && !unicode.IsDigit(char) {
				isSymbol = true
			}
		}
	}
	return isSymbol
}
