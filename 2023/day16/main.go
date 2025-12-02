package main

import (
	"advent-of-code/utils"
	"fmt"
	"slices"

	"github.com/emirpasic/gods/queues/linkedlistqueue"
	"github.com/emirpasic/gods/sets/hashset"
)

var UP = direction{-1, 0}
var RIGHT = direction{0, 1}
var DOWN = direction{1, 0}
var LEFT = direction{0, -1}

type coordinate struct {
	y, x int
}

type direction coordinate

type beam struct {
	pos coordinate
	dir direction
}

func main() {
	lines, _ := utils.ReadLines("2023/day16/input.txt")
	fmt.Println("2023 Day 16 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines))
	fmt.Printf("Part 2: %v\n", part2(lines))
}

func part1(lines []string) int {
	grid := linesToGrid(lines)
	return getEnergizedTiles(&grid, beam{coordinate{0, 0}, RIGHT})
}

func part2(lines []string) int {
	grid := linesToGrid(lines)
	rows, cols := len(grid), len(grid[0])
	candidates := []int{}
	for i := 0; i < rows; i++ {
		candidates = append(candidates, getEnergizedTiles(&grid, beam{coordinate{i, 0}, RIGHT}))
		candidates = append(candidates, getEnergizedTiles(&grid, beam{coordinate{i, cols - 1}, LEFT}))
	}
	for j := 0; j < cols; j++ {
		candidates = append(candidates, getEnergizedTiles(&grid, beam{coordinate{0, j}, DOWN}))
		candidates = append(candidates, getEnergizedTiles(&grid, beam{coordinate{rows - 1, j}, UP}))
	}
	return slices.Max(candidates)
}

func getEnergizedTiles(grid *[][]rune, start beam) int {
	rows, cols := len(*grid), len((*grid)[0])
	energized, visited := hashset.New(), hashset.New()
	queue := linkedlistqueue.New()
	queue.Enqueue(start)
	for !queue.Empty() {
		val, _ := queue.Dequeue()
		curr := val.(beam)
		if (0 <= curr.pos.y && curr.pos.y < rows) && (0 <= curr.pos.x && curr.pos.x < cols) {
			energized.Add(curr.pos)
			if visited.Contains(curr) {
				continue
			}
			visited.Add(curr)
			switch (*grid)[curr.pos.y][curr.pos.x] {
			case '.':
				queue.Enqueue(continueBeam(curr))
			case '/':
				queue.Enqueue(mirrorBeam(curr, []direction{RIGHT, UP, LEFT, DOWN}))
			case '\\':
				queue.Enqueue(mirrorBeam(curr, []direction{LEFT, DOWN, RIGHT, UP}))
			case '|':
				for _, b := range splitBeam(curr, []direction{UP, DOWN}) {
					queue.Enqueue(b)
				}
			case '-':
				for _, b := range splitBeam(curr, []direction{LEFT, RIGHT}) {
					queue.Enqueue(b)
				}
			}
		}
	}
	return energized.Size()
}

func continueBeam(b beam) beam {
	return beam{
		pos: coordinate{b.pos.y + b.dir.y, b.pos.x + b.dir.x},
		dir: b.dir,
	}
}

func mirrorBeam(b beam, dirs []direction) beam {
	switch b.dir {
	case UP:
		b.dir = dirs[0]
	case RIGHT:
		b.dir = dirs[1]
	case DOWN:
		b.dir = dirs[2]
	case LEFT:
		b.dir = dirs[3]
	}
	return continueBeam(b)
}

func splitBeam(b beam, dirs []direction) []beam {
	if slices.Contains(dirs, b.dir) {
		return []beam{continueBeam(b)}
	}
	beams := make([]beam, len(dirs))
	for i, dir := range dirs {
		beams[i] = beam{pos: b.pos, dir: dir}
	}
	return beams
}

func linesToGrid(lines []string) [][]rune {
	grid := [][]rune{}
	for _, line := range lines {
		grid = append(grid, []rune(line))
	}
	return grid
}
