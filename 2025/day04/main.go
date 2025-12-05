package main

import (
	"advent-of-code/utils"
	"fmt"
)

var dirs = [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}}

func main() {
	lines, _ := utils.ReadLines("2025/day04/input.txt")
	fmt.Println("2025 Day 04 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines))
	fmt.Printf("Part 2: %v\n", part2(lines))
}

func part1(lines []string) int {
	grid := linesToGrid(lines)
	count, _ := getCountAndCollectVisited(grid)
	return count
}

func part2(lines []string) int {
	grid := linesToGrid(lines)
	sum := 0
	for {
		count, visited := getCountAndCollectVisited(grid)
		if count == 0 {
			break
		}
		sum += count
		for _, cell := range visited {
			i, j := cell[0], cell[1]
			grid[i][j] = 'x'
		}
	}
	return sum
}

func getCountAndCollectVisited(grid [][]rune) (int, [][]int) {
	count := 0
	visited := [][]int{}
	for i, row := range grid {
		for j := range row {
			if grid[i][j] == '@' && checkNeighbors(grid, i, j) {
				visited = append(visited, []int{i, j})
				count++
			}
		}
	}
	return count, visited
}

func checkNeighbors(grid [][]rune, i, j int) bool {
	count := 0
	for _, dir := range dirs {
		r := i + dir[0]
		c := j + dir[1]
		if r >= 0 && r < len(grid) && c >= 0 && c < len(grid[i]) && grid[r][c] == '@' {
			count++
		}
		if count > 3 {
			return false
		}
	}
	return true
}

func linesToGrid(lines []string) [][]rune {
	grid := [][]rune{}
	for _, line := range lines {
		grid = append(grid, []rune(line))
	}
	return grid
}
