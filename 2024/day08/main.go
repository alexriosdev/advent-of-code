package main

import (
	"advent-of-code/utils"
	"fmt"
)

func main() {
	lines, _ := utils.ReadLines("2024/day08/input.txt")
	fmt.Println("2024 Day 08 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines))
	fmt.Printf("Part 2: %v\n", part2(lines))
}

func part1(lines []string) int {
	grid := convertToGrid(lines)
	frequenciesMap := getFrequenciesMap(grid)
	antinodes := getAntinodes(grid, frequenciesMap, false)
	return len(antinodes)
}

func part2(lines []string) int {
	grid := convertToGrid(lines)
	frequenciesMap := getFrequenciesMap(grid)
	antinodes := getAntinodes(grid, frequenciesMap, true)
	newAntinodesCount := 0
	for i, row := range grid {
		for j, c := range row {
			if !antinodes[coordinate{i, j}] && c != '.' {
				newAntinodesCount++
			}
		}
	}
	return len(antinodes) + newAntinodesCount
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

func getFrequenciesMap(grid [][]rune) map[rune][]coordinate {
	frequenciesMap := map[rune][]coordinate{}
	for i, row := range grid {
		for j, c := range row {
			if c != '.' {
				frequenciesMap[c] = append(frequenciesMap[c], coordinate{i, j})
			}
		}
	}
	return frequenciesMap
}

func getAntinodes(grid [][]rune, frequenciesMap map[rune][]coordinate, useRecursion bool) map[coordinate]bool {
	antinodes := map[coordinate]bool{}
	for _, frequencies := range frequenciesMap {
		for _, a := range frequencies {
			for _, b := range frequencies {
				if a == b {
					continue
				}
				getAntinodesRecursive(grid, a, b, &antinodes, useRecursion)
			}
		}
	}
	return antinodes
}

func getAntinodesRecursive(grid [][]rune, a, b coordinate, antinodes *map[coordinate]bool, useRecursion bool) {
	dx := 2*a.x - b.x
	dy := 2*a.y - b.y
	if isRange(grid, dx, dy) {
		freq := coordinate{dx, dy}
		(*antinodes)[freq] = true
		if useRecursion {
			getAntinodesRecursive(grid, freq, a, antinodes, useRecursion)
		}
	}
}

func isRange(grid [][]rune, i, j int) bool {
	return 0 <= i && i < len(grid) && 0 <= j && j < len(grid[0])
}
