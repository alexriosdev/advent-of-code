package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("2023 Day 13 Solution")
	fmt.Printf("Part 1: %v\n", part1(input))
}

func part1(input []byte) int {
	result := 0
	for _, inputSplit := range strings.Split(strings.TrimSpace(string(input)), "\n\n") {
		lines := strings.Split(inputSplit, "\n")
		grid := linesToGrid(lines)
		rows, cols := len(grid), len(grid[0])
		result += getVerticalSum(grid, rows, cols)
		result += getHorizontalSum(grid, rows, cols)
	}
	return result
}

func getVerticalSum(grid [][]rune, rows, cols int) int {
	sum := 0
	for i := 0; i < cols-1; i++ {
		isVertical := true
		for j := 0; j < cols; j++ {
			if !isVertical {
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
					break
				}
			}
		}
		if isVertical {
			sum += i + 1
		}
	}
	return sum
}

func getHorizontalSum(grid [][]rune, rows, cols int) int {
	sum := 0
	for i := 0; i < rows-1; i++ {
		isHorizontal := true
		for j := 0; j < rows; j++ {
			if !isHorizontal {
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
					break
				}
			}
		}
		if isHorizontal {
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
