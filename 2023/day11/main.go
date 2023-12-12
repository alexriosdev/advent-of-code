package main

import (
	"advent-of-code/utils"
	"fmt"
	"slices"
	"sort"
)

type coordinate struct {
	y, x int
}

func (a *coordinate) getDistance(b *coordinate) int {
	return getAbs(a.x-b.x) + getAbs(a.y-b.y)
}

func main() {
	lines, _ := utils.ReadLines("input.txt")
	fmt.Println("2023 Day 11 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines))
	fmt.Printf("Part 2: %v\n", part2(lines, 1000000))
}

func part1(lines []string) int {
	grid := linesToGrid(lines)
	galaxies := getGalaxies(grid)
	emptyRows := getEmptyRows(grid)
	emptyCols := getEmptyCols(grid)
	return getDistanceSum(2, galaxies, emptyRows, emptyCols)
}

func part2(lines []string, factor int) int {
	grid := linesToGrid(lines)
	galaxies := getGalaxies(grid)
	emptyRows := getEmptyRows(grid)
	emptyCols := getEmptyCols(grid)
	return getDistanceSum(factor, galaxies, emptyRows, emptyCols)
}

func getDistanceSum(factor int, galaxies []coordinate, emptyRows, emptyCols []int) int {
	sum := 0
	pairs := map[[2]int]bool{}
	for i := range galaxies {
		for j := range galaxies {
			curr := [2]int{i, j}
			sort.Ints(curr[:])
			if i == j || pairs[curr] {
				continue
			}
			a, b := galaxies[i], galaxies[j]
			dist := a.getDistance(&b)
			rangeRows := makeRange(getMin(a.y, b.y), getMax(a.y, b.y))
			for _, row := range emptyRows {
				if slices.Contains(rangeRows, row) {
					dist += factor - 1
				}
			}
			rangeCols := makeRange(getMin(a.x, b.x), getMax(a.x, b.x))
			for _, col := range emptyCols {
				if slices.Contains(rangeCols, col) {
					dist += factor - 1
				}
			}
			sum += dist
			pairs[curr] = true
		}
	}
	return sum
}

func getGalaxies(grid [][]rune) []coordinate {
	galaxies := []coordinate{}
	for i, row := range grid {
		for j, c := range row {
			if c == '#' {
				galaxies = append(galaxies, coordinate{y: i, x: j})
			}
		}
	}
	return galaxies
}

func getEmptyRows(grid [][]rune) []int {
	emptyRows := []int{}
	for i, row := range grid {
		isEmpty := true
		for _, c := range row {
			if c == '#' {
				isEmpty = false
				break
			}
		}
		if isEmpty {
			emptyRows = append(emptyRows, i)
		}
	}
	return emptyRows
}

func getEmptyCols(grid [][]rune) []int {
	emptyCols := []int{}
	for i := range grid[0] {
		isEmpty := true
		for _, col := range grid {
			if col[i] == '#' {
				isEmpty = false
				break
			}
		}
		if isEmpty {
			emptyCols = append(emptyCols, i)
		}
	}
	return emptyCols
}

func linesToGrid(lines []string) [][]rune {
	grid := [][]rune{}
	for _, line := range lines {
		grid = append(grid, []rune(line))
	}
	return grid
}

func makeRange(min, max int) []int {
	nums := make([]int, max-min+1)
	for i := range nums {
		nums[i] = min + i
	}
	return nums
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getAbs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
