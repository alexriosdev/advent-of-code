package main

import (
	"advent-of-code/utils"
	"fmt"
	"math"
	"slices"

	"github.com/emirpasic/gods/sets/hashset"
	"github.com/emirpasic/gods/stacks/arraystack"
)

var UP		= coordinate{-1, 0}
var RIGHT	= coordinate{0, 1}
var DOWN	= coordinate{1, 0}
var LEFT	= coordinate{0, -1}

var SLOPE_MAP = map[rune][]coordinate{
	'<': {LEFT},
	'>': {RIGHT},
	'^': {UP},
	'v': {DOWN},
	'.': {UP, RIGHT, DOWN, LEFT},
}

type coordinate struct {
	y, x int
}

type state struct {
	pos   coordinate
	steps int
}

func main() {
	lines, _ := utils.ReadLines("input.txt")
	fmt.Println("2023 Day 23 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines))
}

func part1(lines []string) int {
	grid := linesToGrid(lines)
	start := coordinate{0, slices.Index(grid[0], '.')}
	end := coordinate{len(grid) - 1, slices.Index(grid[len(grid)-1], '.')}
	walkable := []coordinate{start, end}
	for i, row := range grid {
		for j, c := range row {
			if c != '#' && isWalkable(grid, i, j) {
				walkable = append(walkable, coordinate{i, j})
			}
		}
	}
	graph := getAdjacencyGraph(grid, walkable)
	return getMaxSteps(&graph, hashset.New(), start, end)
}

func getMaxSteps(graph *map[coordinate]map[coordinate]int, visited *hashset.Set, curr, end coordinate) int {
	if curr == end {
		return 0
	}
	result := math.MinInt
	visited.Add(curr)
	for next := range (*graph)[curr] {
		result = max(result, getMaxSteps(graph, visited, next, end)+(*graph)[curr][next])
	}
	visited.Remove(curr)
	return result
}

func getAdjacencyGraph(grid [][]rune, walkable []coordinate) map[coordinate]map[coordinate]int {
	graph := make(map[coordinate]map[coordinate]int)
	for _, pos := range walkable {
		graph[pos] = make(map[coordinate]int)
	}
	for _, pos := range walkable {
		stack := arraystack.New()
		stack.Push(state{pos, 0})
		visited := map[coordinate]bool{}
		visited[pos] = true
		for !stack.Empty() {
			val, _ := stack.Pop()
			curr := val.(state)
			if curr.steps != 0 && slices.Contains(walkable, curr.pos) {
				graph[pos][curr.pos] = curr.steps
				continue
			}
			c := grid[curr.pos.y][curr.pos.x]
			for _, dir := range SLOPE_MAP[c] {
				next := coordinate{curr.pos.y + dir.y, curr.pos.x + dir.x}
				if isWithinBounds(grid, next) && grid[next.y][next.x] != '#' && !visited[next] {
					stack.Push(state{next, curr.steps + 1})
					visited[next] = true
				}
			}
		}
	}
	return graph
}

func isWalkable(grid [][]rune, i, j int) bool {
	count := 0
	for _, dir := range []coordinate{UP, RIGHT, DOWN, LEFT} {
		next := coordinate{i + dir.y, j + dir.x}
		if isWithinBounds(grid, next) && grid[next.y][next.x] != '#' {
			count++
		}
	}
	return count > 2
}

func isWithinBounds(grid [][]rune, pos coordinate) bool {
	return (0 <= pos.y && pos.y < len(grid)) && (0 <= pos.x && pos.x < len(grid[0]))
}

func linesToGrid(lines []string) [][]rune {
	grid := [][]rune{}
	for _, line := range lines {
		grid = append(grid, []rune(line))
	}
	return grid
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
