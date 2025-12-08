package main

import (
	"advent-of-code/utils"
	"fmt"
)

func main() {
	lines, _ := utils.ReadLines("2025/day07/input.txt")
	fmt.Println("2025 Day 07 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines))
	fmt.Printf("Part 2: %v\n", part2(lines))
}

func part1(lines []string) int {
	grid := convertToGrid(lines)
	count := 0
	for r, row := range grid {
		for c, curr := range row {
			prev := '.'
			if r > 0 {
				prev = grid[r-1][c]
			}
			switch {
			case curr == 'S':
				grid[r+1][c] = '|'
			case curr == '.' && prev == '|':
				grid[r][c] = '|'
			case curr == '^' && prev == '|':
				count++
				grid[r][c-1] = '|'
				grid[r][c+1] = '|'
			}
		}
	}
	return count
}

func part2(lines []string) int {
	return len(lines)
}

func displayGrid(grid *[][]rune) {
	for _, row := range *grid {
		fmt.Println(string(row))
	}
	fmt.Println()
}

func convertToGrid(lines []string) [][]rune {
	grid := [][]rune{}
	for _, line := range lines {
		grid = append(grid, []rune(line))
	}
	return grid
}
