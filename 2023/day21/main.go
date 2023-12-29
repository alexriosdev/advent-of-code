package main

import (
	"advent-of-code/utils"
	"fmt"

	"github.com/emirpasic/gods/queues/linkedlistqueue"
)

var UP = coordinate{-1, 0}
var RIGHT = coordinate{0, 1}
var DOWN = coordinate{1, 0}
var LEFT = coordinate{0, -1}

type coordinate struct {
	y, x int
}

type state struct {
	pos   coordinate
	steps int
}

func main() {
	lines, _ := utils.ReadLines("input.txt")
	fmt.Println("2023 Day 21 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines, 64))
}

func part1(lines []string, steps int) int {
	grid := linesToGrid(lines)
	var start coordinate
	for i, row := range grid {
		for j, c := range row {
			if c == 'S' {
				start = coordinate{i, j}
				break
			}
		}
	}
	return getGardenPlotCount(grid, start, steps)
}

func getGardenPlotCount(grid [][]rune, start coordinate, steps int) int {
	count := 0
	visited := map[coordinate]bool{}
	visited[start] = true
	queue := linkedlistqueue.New()
	queue.Enqueue(state{start, steps})
	for !queue.Empty() {
		val, _ := queue.Dequeue()
		curr := val.(state)
		if curr.steps%2 == 0 {
			count++
		}
		for _, dir := range []coordinate{UP, RIGHT, DOWN, LEFT} {
			next := coordinate{curr.pos.y + dir.y, curr.pos.x + dir.x}
			if (0 <= next.y && next.y < len(grid)) && (0 <= next.x && next.x < len(grid[0])) && grid[next.y][next.x] != '#' && !visited[next] && curr.steps-1 >= 0 {
				visited[next] = true
				queue.Enqueue(state{next, curr.steps - 1})
			}
		}
	}
	return count
}

func linesToGrid(lines []string) [][]rune {
	grid := [][]rune{}
	for _, line := range lines {
		grid = append(grid, []rune(line))
	}
	return grid
}
