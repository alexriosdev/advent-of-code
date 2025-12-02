package main

import (
	"advent-of-code/utils"
	"fmt"
	"unicode"
)

func main() {
	lines, _ := utils.ReadLines("2023/day03/input.txt")
	fmt.Println("2023 Day 03 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines))
	fmt.Printf("Part 2: %v\n", part2(lines))
}

func part1(lines []string) int {
	result := 0
	grid := convertToGrid(lines)
	for i, row := range grid {
		num := 0
		isValid := false
		for j := range row {
			c := grid[i][j]
			if unicode.IsDigit(c) {
				num = (num * 10) + int(c-'0')
				if isValid == false {
					isValid = checkNeighbors(grid, i, j)
				}
				continue
			}
			if isValid && num > 0 {
				result += num
				isValid = false
			}
			num = 0
		}
		if isValid && num > 0 {
			result += num
		}
	}
	return result
}

func part2(lines []string) int {
	result := 0
	grid := convertToGrid(lines)
	for i, row := range grid {
		for j := range row {
			c := grid[i][j]
			if c != '*' {
				continue
			}
			gearSet := checkNeighborsGearSet(grid, i, j)
			if len(gearSet) == 2 {
				ratio := 1
				for k := range gearSet {
					ratio *= k
				}
				result += ratio
			}
		}
	}
	return result
}

func convertToGrid(lines []string) [][]rune {
	grid := [][]rune{}
	for _, line := range lines {
		grid = append(grid, []rune(line))
	}
	return grid
}

func checkNeighbors(grid [][]rune, i, j int) bool {
	isSymbol := false
	dirs := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}}
	for _, dir := range dirs {
		r := i + dir[0]
		c := j + dir[1]
		if c >= 0 && c < len(grid[i]) && r >= 0 && r < len(grid) {
			char := grid[r][c]
			if char != '.' && !unicode.IsDigit(char) {
				isSymbol = true
			}
		}
	}
	return isSymbol
}

func checkNeighborsGearSet(grid [][]rune, i, j int) map[int]bool {
	gearSet := map[int]bool{}
	dirs := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}}
	for _, dir := range dirs {
		r := i + dir[0]
		c := j + dir[1]
		if c >= 0 && c < len(grid[i]) && r >= 0 && r < len(grid) {
			char := grid[r][c]
			if unicode.IsDigit(char) {
				num := runesToInt(grid[r], c)
				gearSet[num] = true
			}
		}
	}
	return gearSet
}

func runesToInt(runes []rune, i int) int {
	for i > 0 && unicode.IsDigit(runes[i-1]) {
		i--
	}
	num := 0
	for ; i < len(runes) && unicode.IsDigit(runes[i]); i++ {
		num = (num * 10) + int(runes[i]-'0')
	}
	return num
}
