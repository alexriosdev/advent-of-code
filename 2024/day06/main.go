package main

import (
	"advent-of-code/utils"
	"fmt"
)

func main() {
	lines, _ := utils.ReadLines("input.txt")
	fmt.Println("2024 Day 06 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines))
}

func part1(lines []string) int {
	grid := convertToGrid(lines)
	start := getStart(grid)
	visited := map[coordinate]bool{}
	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	i := 0
	for 0 <= start.x && start.x < len(grid) && 0 <= start.y && start.y < len(grid[0]) {
		if grid[start.x][start.y] != '#' {
			visited[start] = true
			start.x += dirs[i%4][0]
			start.y += dirs[i%4][1]
		} else {
			start.x -= dirs[i%4][0]
			start.y -= dirs[i%4][1]
			i++
		}
	}
	return len(visited)
}

type coordinate struct {
	x, y int
}

func convertToGrid(lines []string) [][]rune {
	grid := [][]rune{}
	for _, line := range lines {
		grid = append(grid, []rune(line))
	}
	return grid
}

func getStart(grid [][]rune) coordinate {
	for i, row := range grid {
		for j, cell := range row {
			if cell == '^' {
				return coordinate{i, j}
			}
		}
	}
	return coordinate{-1, -1}
}
