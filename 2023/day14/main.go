package main

import (
	"advent-of-code/utils"
	"fmt"
)

func main() {
	lines, _ := utils.ReadLines("input.txt")
	fmt.Println("2023 Day 14 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines))
}

func part1(lines []string) int {
	grid := linesToGrid(lines)
	rows, cols := len(grid), len(grid[0])
	tiltGrid(&grid, rows, cols)
	return getRockSum(&grid, rows)
}

func tiltGrid(grid *[][]rune, rows, cols int) {
	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			for k := 0; k < rows; k++ {
				if (*grid)[k][i] == 'O' && k > 0 && (*grid)[k-1][i] == '.' {
					(*grid)[k][i] = '.'
					(*grid)[k-1][i] = 'O'
				}
			}
		}
	}
	return
}

func getRockSum(grid *[][]rune, rows int) int {
	sum := 0
	for i, row := range *grid {
		for _, c := range row {
			if c == 'O' {
				sum += rows - i
			}
		}
	}
	return sum
}

func linesToGrid(lines []string) [][]rune {
	grid := [][]rune{}
	for _, line := range lines {
		grid = append(grid, []rune(line))
	}
	return grid
}
