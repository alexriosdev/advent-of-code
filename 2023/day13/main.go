package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("2023/day13/input.txt")
	fmt.Println("2023 Day 13 Solution")
	fmt.Printf("Part 1: %v\n", part1(input))
	fmt.Printf("Part 2: %v\n", part2(input))
}

func part1(input []byte) int {
	result := 0
	for _, inputSplit := range strings.Split(strings.TrimSpace(string(input)), "\n\n") {
		lines := strings.Split(inputSplit, "\n")
		grid := linesToGrid(lines)
		rows, cols := len(grid), len(grid[0])
		result += getVerticalSum(grid, rows, cols, false)
		result += getHorizontalSum(grid, rows, cols, false)
	}
	return result
}

func part2(input []byte) int {
	result := 0
	for _, inputSplit := range strings.Split(strings.TrimSpace(string(input)), "\n\n") {
		lines := strings.Split(inputSplit, "\n")
		grid := linesToGrid(lines)
		rows, cols := len(grid), len(grid[0])
		result += getVerticalSum(grid, rows, cols, true)
		result += getHorizontalSum(grid, rows, cols, true)
	}
	return result
}

func getVerticalSum(grid [][]rune, rows, cols int, hasSmudge bool) int {
	sum := 0
	for i := 0; i < cols-1; i++ {
		isVertical := true
		smudge := 0
		for j := 0; j < cols; j++ {
			if !isVertical && !hasSmudge {
				break
			}
			left := i - j
			right := i + j + 1
			if !(0 <= left && right < cols) {
				continue
			}
			for k := 0; k < rows; k++ {
				if grid[k][left] != grid[k][right] {
					isVertical = false
					smudge++
				}
				if !isVertical && !hasSmudge {
					break
				}
			}
		}
		if (isVertical && !hasSmudge) || (hasSmudge && smudge == 1) {
			sum += i + 1
		}
	}
	return sum
}

func getHorizontalSum(grid [][]rune, rows, cols int, hasSmudge bool) int {
	sum := 0
	for i := 0; i < rows-1; i++ {
		isHorizontal := true
		smudge := 0
		for j := 0; j < rows; j++ {
			if !isHorizontal && !hasSmudge {
				break
			}
			up := i - j
			down := i + j + 1
			if !(0 <= up && down < rows) {
				continue
			}
			for k := 0; k < cols; k++ {
				if grid[up][k] != grid[down][k] {
					isHorizontal = false
					smudge++
				}
				if !isHorizontal && !hasSmudge {
					break
				}
			}
		}
		if (isHorizontal && !hasSmudge) || (hasSmudge && smudge == 1) {
			sum += (i + 1) * 100
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
