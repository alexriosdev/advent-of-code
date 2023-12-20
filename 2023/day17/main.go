package main

import (
	"advent-of-code/utils"
	"fmt"

	"github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/sets/hashset"
	pqutil "github.com/emirpasic/gods/utils"
)

var UP = coordinate{-1, 0}
var RIGHT = coordinate{0, 1}
var DOWN = coordinate{1, 0}
var LEFT = coordinate{0, -1}
var ORIGIN = coordinate{0, 0}

type coordinate struct {
	y, x int
}

type state struct {
	pos, dir coordinate
	dist     int
}

type heatState struct {
	state    state
	heatLoss int
}

func byHeatLoss(a, b interface{}) int {
	priorityA := a.(heatState).heatLoss
	priorityB := b.(heatState).heatLoss
	return pqutil.IntComparator(priorityA, priorityB)
}

func main() {
	lines, _ := utils.ReadLines("input.txt")
	fmt.Println("2023 Day 17 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines))
}

func part1(lines []string) int {
	grid := linesToGrid(lines)
	start, end := ORIGIN, coordinate{len(grid) - 1, len(grid[0]) - 1}
	return getMinHeatLoss(&grid, start, end, 3)
}

func getMinHeatLoss(grid *[][]rune, start, end coordinate, maxDist int) int {
	visited := hashset.New()
	pq := priorityqueue.NewWith(byHeatLoss)
	pq.Enqueue(heatState{state{start, start, 0}, 0})
	for !pq.Empty() {
		val, _ := pq.Dequeue()
		curr := val.(heatState).state
		heatLoss := val.(heatState).heatLoss
		if curr.pos == end {
			return heatLoss
		}
		if visited.Contains(curr) {
			continue
		}
		visited.Add(curr)
		if curr.dist < maxDist && curr.dir != ORIGIN {
			next := coordinate{curr.pos.y + curr.dir.y, curr.pos.x + curr.dir.x}
			if isWithinBounds(start, next, end) {
				pq.Enqueue(heatState{state{next, curr.dir, curr.dist + 1}, heatLoss + int((*grid)[next.y][next.x]) - '0'})
			}
		}
		for _, dir := range []coordinate{UP, RIGHT, DOWN, LEFT} {
			if dir != curr.dir && dir != reverseDir(curr.dir) {
				next := coordinate{curr.pos.y + dir.y, curr.pos.x + dir.x}
				if isWithinBounds(start, next, end) {
					pq.Enqueue(heatState{state{next, dir, 1}, heatLoss + int((*grid)[next.y][next.x]) - '0'})
				}
			}
		}
	}
	return -1
}

func isWithinBounds(start, next, end coordinate) bool {
	return (start.y <= next.y && next.y <= end.y) && (start.x <= next.x && next.x <= end.x)
}

func reverseDir(c coordinate) coordinate {
	return coordinate{-c.y, -c.x}
}

func linesToGrid(lines []string) [][]rune {
	grid := [][]rune{}
	for _, line := range lines {
		grid = append(grid, []rune(line))
	}
	return grid
}
