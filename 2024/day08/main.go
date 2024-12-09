package main

import (
	"advent-of-code/utils"
	"fmt"
)

func main() {
	lines, _ := utils.ReadLines("input.txt")
	fmt.Println("2024 Day 08 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines))
}

func part1(lines []string) int {
	grid := convertToGrid(lines)
	frequenciesMap := getFrequenciesMap(grid)
	antinodes := getAntinodes(grid, frequenciesMap)
	return len(antinodes)
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

func getAntinodes(grid [][]rune, frequenciesMap map[rune][]coordinate) map[coordinate]bool {
	antinodes := map[coordinate]bool{}
	for _, frequencies := range frequenciesMap {
		for _, a := range frequencies {
			for _, b := range frequencies {
				if a == b {
					continue
				}
				dx := 2*a.x - b.x
				dy := 2*a.y - b.y
				if isRange(grid, dx, dy) {
					antinodes[coordinate{dx, dy}] = true
				}
			}
		}
	}
	return antinodes
}

func isRange(grid [][]rune, i, j int) bool {
	return 0 <= i && i < len(grid) && 0 <= j && j < len(grid[0])
}
