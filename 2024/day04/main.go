package main

import (
	"advent-of-code/utils"
	"fmt"
)

func main() {
	lines, _ := utils.ReadLines("input.txt")
	fmt.Println("2024 Day 04 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines))
	fmt.Printf("Part 2: %v\n", part2(lines))
}

func part1(lines []string) int {
	grid := convertToGrid(lines)
	dirs := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}}
	count := 0
	for i := range grid {
		for j := range grid[0] {
			for _, dir := range dirs {
				if isValidPattern("XMAS", grid, i, j, dir) {
					count++
				}
			}
		}
	}
	return count
}

func part2(lines []string) int {
	grid := convertToGrid(lines)
	count := 0
	for i := range grid {
		for j := range grid[0] {
			d1 := isValidPattern("MAS", grid, i, j, []int{1, 1})
			d2 := isValidPattern("MAS", grid, i+2, j+2, []int{-1, -1})
			d3 := isValidPattern("MAS", grid, i+2, j, []int{-1, 1})
			d4 := isValidPattern("MAS", grid, i, j+2, []int{1, -1})
			if (d1 || d2) && (d3 || d4) {
				count++
			}
		}
	}
	return count
}

func isValidPattern(s string, grid [][]rune, i, j int, dir []int) bool {
	for _, c := range s {
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
