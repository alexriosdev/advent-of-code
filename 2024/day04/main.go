package main

import (
	"advent-of-code/utils"
	"fmt"
)

func main() {
	lines, _ := utils.ReadLines("input.txt")
	fmt.Println("2024 Day 04 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines))
}

func part1(lines []string) int {
	grid := convertToGrid(lines)
	dirs := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}}
	count := 0
	for i := range grid {
		for j := range grid[0] {
			for _, dir := range dirs {
				if isValid(grid, i, j, dir) {
					count++
				}
			}
		}
	}
	return count
}

func isValid(grid [][]rune, i, j int, dir []int) bool {
	for _, c := range "XMAS" {
		if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || c != grid[i][j] {
			return false
		}
		i += dir[0]
		j += dir[1]
	}
	return true
}

func convertToGrid(lines []string) [][]rune {
	grid := [][]rune{}
	for _, line := range lines {
		grid = append(grid, []rune(line))
	}
	return grid
}
