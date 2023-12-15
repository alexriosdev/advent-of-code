package main

import (
	"advent-of-code/utils"
	"fmt"
	"strings"
)

func main() {
	lines, _ := utils.ReadLines("input.txt")
	fmt.Println("2023 Day 14 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines))
	fmt.Printf("Part 2: %v\n", part2(lines))
}

func part1(lines []string) int {
	grid := linesToGrid(lines)
	tiltGrid(&grid)
	return getRockSum(&grid, len(grid))
}

func part2(lines []string) int {
	vGrid := linesToGrid(lines)
	rows, cols := len(vGrid), len(vGrid[0])
	hGrid := createHorizontalGrid(rows, cols)
	visited := map[string]int{}
	for i := 0; i < 1000000000; i++ {
		for j := 0; j < 2; j++ {
			tiltGrid(&vGrid)
			rotateGrid(&vGrid, &hGrid)
			tiltGrid(&hGrid)
			rotateGrid(&hGrid, &vGrid)
		}
		key := gridToString(&vGrid)
		if val, ok := visited[key]; ok {
			cycle := i - val
			i += ((1000000000 - i) / cycle) * cycle
		}
		visited[key] = i
	}
	return getRockSum(&vGrid, rows)
}

func tiltGrid(grid *[][]rune) {
	rows, cols := len(*grid), len((*grid)[0])
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

func rotateGrid(gridA, gridB *[][]rune) {
	rows, cols := len(*gridA), len((*gridA)[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			(*gridB)[j][rows-i-1] = (*gridA)[i][j]
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

func createHorizontalGrid(rows, cols int) [][]rune {
	hGrid := make([][]rune, cols)
	for i := range hGrid {
		hGrid[i] = make([]rune, rows)
	}
	return hGrid
}

func gridToString(grid *[][]rune) string {
	sb := strings.Builder{}
	for _, row := range *grid {
		for _, c := range row {
			sb.WriteRune(c)
		}
	}
	return sb.String()
}

func linesToGrid(lines []string) [][]rune {
	grid := [][]rune{}
	for _, line := range lines {
		grid = append(grid, []rune(line))
	}
	return grid
}
