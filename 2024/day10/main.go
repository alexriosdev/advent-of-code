package main

import (
	"advent-of-code/utils"
	"fmt"

	"github.com/emirpasic/gods/queues/linkedlistqueue"
	"github.com/emirpasic/gods/sets/hashset"
)

func main() {
	lines, _ := utils.ReadLines("input.txt")
	fmt.Println("2024 Day 10 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines))
	fmt.Printf("Part 2: %v\n", part2(lines))
}

func part1(lines []string) int {
	grid := convertToGrid(lines)
	result := 0
	for i, row := range grid {
		for j, c := range row {
			if c == '0' {
				score, _ := bfs(grid, coordinate{i, j})
				result += score
			}
		}
	}
	return result
}

func part2(lines []string) int {
	grid := convertToGrid(lines)
	result := 0
	for i, row := range grid {
		for j, c := range row {
			if c == '0' {
				_, rating := bfs(grid, coordinate{i, j})
				result += rating
			}
		}
	}
	return result
}

type coordinate struct {
	x, y int
}

func bfs(grid [][]rune, start coordinate) (int, int) {
	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	queue := linkedlistqueue.New()
	queue.Enqueue(start)
	score, rating := hashset.New(), 0
	for !queue.Empty() {
		size := queue.Size()
		for i := 0; i < size; i++ {
			val, _ := queue.Dequeue()
			curr := val.(coordinate)
			if grid[curr.x][curr.y] == '9' {
				score.Add(coordinate{curr.x, curr.y})
				rating++
			}
			for _, dir := range dirs {
				nextX := curr.x + dir[0]
				nextY := curr.y + dir[1]
				next := coordinate{nextX, nextY}
				if isRange(grid, nextX, nextY) && grid[nextX][nextY] == grid[curr.x][curr.y]+1 {
					queue.Enqueue(next)
				}
			}
		}
	}
	return score.Size(), rating
}

func convertToGrid(lines []string) [][]rune {
	grid := [][]rune{}
	for _, line := range lines {
		grid = append(grid, []rune(line))
	}
	return grid
}

func isRange(grid [][]rune, i, j int) bool {
	return 0 <= i && i < len(grid) && 0 <= j && j < len(grid[0])
}
