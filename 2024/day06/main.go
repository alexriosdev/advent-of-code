package main

import (
	"advent-of-code/utils"
	"fmt"
)

func main() {
	lines, _ := utils.ReadLines("2024/day06/input.txt")
	fmt.Println("2024 Day 06 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines))
	fmt.Printf("Part 2: %v\n", part2(lines))
}

func part1(lines []string) int {
	grid := convertToGrid(lines)
	start := getStart(grid)
	visited := getVisited(grid, start)
	return len(visited)
}

func part2(lines []string) int {
	grid := convertToGrid(lines)
	start := getStart(grid)
	visited := getVisited(grid, start)
	result := 0
	for cell := range visited {
		grid[cell.x][cell.y] = '#'
		if isCycle(grid, start) {
			result++
		}
		grid[cell.x][cell.y] = '.'
	}
	return result
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

func getVisited(grid [][]rune, start coordinate) map[coordinate]bool {
	visited := map[coordinate]bool{}
	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	i := 0
	for isRange(grid, start.x, start.y) {
		visited[start] = true
		nextX := start.x + dirs[i%4][0]
		nextY := start.y + dirs[i%4][1]
		if isRange(grid, nextX, nextY) && grid[nextX][nextY] == '#' {
			i++
		} else {
			start.x = nextX
			start.y = nextY
		}
	}
	return visited
}

func isCycle(grid [][]rune, start coordinate) bool {
	visited := map[coordinate]int{}
	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	i := 0
	for isRange(grid, start.x, start.y) {
		if visited[start] > len(dirs) {
			return true
		}
		visited[start]++
		nextX := start.x + dirs[i%4][0]
		nextY := start.y + dirs[i%4][1]
		if isRange(grid, nextX, nextY) && grid[nextX][nextY] == '#' {
			i++
		} else {
			start.x = nextX
			start.y = nextY
		}
	}
	return false
}

func isRange(grid [][]rune, i, j int) bool {
	return 0 <= i && i < len(grid) && 0 <= j && j < len(grid[0])
}
