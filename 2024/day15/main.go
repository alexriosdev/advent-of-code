package main

import (
	"fmt"
	"os"
	"strings"
)

var UP = coordinate{-1, 0}
var RIGHT = coordinate{0, 1}
var DOWN = coordinate{1, 0}
var LEFT = coordinate{0, -1}

func main() {
	input, _ := os.ReadFile("2024/day15/input.txt")
	fmt.Println("2024 Day 15 Solution")
	fmt.Printf("Part 1: %v\n", part1(input))
}

func part1(input []byte) int {
	sections := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	grid := convertToGrid(strings.Split(sections[0], "\n"))
	dirs := getDirs(strings.Split(sections[1], "\n"))
	start := getStart(grid)
	moveRobot(&grid, start, dirs)
	result := 0
	for i, row := range grid {
		for j, c := range row {
			if c == 'O' {
				result += (100 * i) + j
			}
		}
	}
	return result
}

type coordinate struct {
	x, y int
}

func getStart(grid [][]rune) coordinate {
	for i, row := range grid {
		for j, cell := range row {
			if cell == '@' {
				return coordinate{i, j}
			}
		}
	}
	return coordinate{-1, -1}
}

func getDirs(lines []string) []coordinate {
	dirs := []coordinate{}
	for _, s := range lines {
		for _, c := range s {
			switch c {
			case '^':
				dirs = append(dirs, UP)
			case '>':
				dirs = append(dirs, RIGHT)
			case 'v':
				dirs = append(dirs, DOWN)
			case '<':
				dirs = append(dirs, LEFT)
			}
		}
	}
	return dirs
}

func moveRobot(grid *[][]rune, start coordinate, dirs []coordinate) {
	for _, dir := range dirs {
		next := coordinate{
			start.x + dir.x,
			start.y + dir.y,
		}
		if !isRange((*grid), next.x, next.y) || (*grid)[next.x][next.y] == '#' {
			continue
		}
		if (*grid)[next.x][next.y] == '.' {
			(*grid)[next.x][next.y] = '@'
			(*grid)[start.x][start.y] = '.'
			start = next
			continue
		}
		r, c := next.x+dir.x, next.y+dir.y
		for (*grid)[r][c] == 'O' && isRange((*grid), r, c) {
			r += dir.x
			c += dir.y
		}
		if (*grid)[r][c] == '#' && isRange((*grid), r, c) {
			continue
		}
		(*grid)[r][c] = 'O'
		(*grid)[next.x][next.y] = '@'
		(*grid)[start.x][start.y] = '.'
		start = next
	}
}

func convertToGrid(lines []string) [][]rune {
	grid := [][]rune{}
	for _, line := range lines {
		grid = append(grid, []rune(line))
	}
	return grid
}

func isRange(grid [][]rune, i, j int) bool {
	return 0 <= i && i < len(grid) && 0 <= j && j < len(grid[0])
}
