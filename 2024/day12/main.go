package main

import (
	"advent-of-code/utils"
	"fmt"

	"github.com/emirpasic/gods/queues/linkedlistqueue"
	"github.com/emirpasic/gods/sets"
	"github.com/emirpasic/gods/sets/hashset"
)

func main() {
	lines, _ := utils.ReadLines("input.txt")
	fmt.Println("2024 Day 12 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines))
}

func part1(lines []string) int {
	grid := convertToGrid(lines)
	visited := hashset.New()
	regions := []region{}
	for i, row := range grid {
		for j, c := range row {
			if !visited.Contains(coordinate{i, j}) {
				region := bfs(grid, coordinate{i, j}, c)
				regions = append(regions, region)
				visited.Add(region.plots.Values()...)
			}
		}
	}
	result := 0
	for _, region := range regions {
		result += region.area * region.perimeter
	}
	return result
}

type coordinate struct {
	x, y int
}

type region struct {
	plots           sets.Set
	area, perimeter int
}

func bfs(grid [][]rune, start coordinate, c rune) region {
	plots, visited := hashset.New(), hashset.New()
	queue := linkedlistqueue.New()
	queue.Enqueue(start)
	visited.Add(start)
	perimeter := 0
	for !queue.Empty() {
		size := queue.Size()
		for i := 0; i < size; i++ {
			val, _ := queue.Dequeue()
			curr := val.(coordinate)
			plots.Add(curr)
			neighbors := getNeighbors(grid, curr, c)
			perimeter += 4 - len(neighbors)
			for _, next := range neighbors {
				if !visited.Contains(next) {
					queue.Enqueue(next)
					visited.Add(next)
				}
			}
		}
	}
	return region{plots, plots.Size(), perimeter}
}

func getNeighbors(grid [][]rune, curr coordinate, c rune) []coordinate {
	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	neighbors := []coordinate{}
	for _, dir := range dirs {
		nextX := curr.x + dir[0]
		nextY := curr.y + dir[1]
		next := coordinate{nextX, nextY}
		if isRange(grid, nextX, nextY) && grid[next.x][next.y] == c {
			neighbors = append(neighbors, next)
		}
	}
	return neighbors
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
