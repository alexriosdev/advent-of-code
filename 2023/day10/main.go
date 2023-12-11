package main

import (
	"advent-of-code/utils"
	"fmt"
)

const (
	UP    = 0
	RIGHT = 1
	DOWN  = 2
	LEFT  = 3
)

type coordinate struct {
	y, x int
}

type tileMap map[coordinate][4]bool

func main() {
	lines, _ := utils.ReadLines("input.txt")
	fmt.Println("2023 Day 10 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines))
}

func part1(lines []string) int {
	grid := linesToGrid(lines)
	start := coordinate{}
	graph := tileMap{}
	for i, row := range grid {
		for j, tile := range row {
			pos := coordinate{
				y: i,
				x: j,
			}
			if tile == 'S' {
				start = pos
			}
			graph[pos] = getTileValue(tile)
		}
	}
	loop := findLoop(graph, start)
	return len(loop) / 2
}

func getTileValue(tile rune) [4]bool {
	switch tile {
	case '|':
		return [4]bool{UP: true, RIGHT: false, DOWN: true, LEFT: false}
	case '-':
		return [4]bool{UP: false, RIGHT: true, DOWN: false, LEFT: true}
	case 'L':
		return [4]bool{UP: true, RIGHT: true, DOWN: false, LEFT: false}
	case 'J':
		return [4]bool{UP: true, RIGHT: false, DOWN: false, LEFT: true}
	case '7':
		return [4]bool{UP: false, RIGHT: false, DOWN: true, LEFT: true}
	case 'F':
		return [4]bool{UP: false, RIGHT: true, DOWN: true, LEFT: false}
	default:
		return [4]bool{UP: false, RIGHT: false, DOWN: false, LEFT: false}
	}
}

func findLoop(graph tileMap, pos coordinate) map[coordinate]bool {
	startCandidates := []rune{'J', '|', '-', 'L', '7', 'F'}
	for _, tile := range startCandidates {
		graph[pos] = getTileValue(tile)
		if loop, isFound := findLoopPath(graph, pos); isFound {
			return loop
		}
	}
	return nil
}

func findLoopPath(graph tileMap, pos coordinate) (map[coordinate]bool, bool) {
	var dir int
	for i, val := range graph[pos] {
		if val {
			dir = i
			break
		}
	}
	visited := map[coordinate]bool{}
	for {
		if _, ok := visited[pos]; ok {
			return visited, true
		}
		visited[pos] = true
		switch dir {
		case UP:
			pos.y--
			dir = DOWN
		case RIGHT:
			pos.x++
			dir = LEFT
		case DOWN:
			pos.y++
			dir = UP
		case LEFT:
			pos.x--
			dir = RIGHT
		}
		if !graph[pos][dir] {
			return nil, false
		}
		for i, val := range graph[pos] {
			if i != dir && val {
				dir = i
				break
			}
		}
	}
}

func linesToGrid(lines []string) [][]rune {
	grid := [][]rune{}
	for _, line := range lines {
		grid = append(grid, []rune(line))
	}
	return grid
}
