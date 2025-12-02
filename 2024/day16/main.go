package main

import (
	"advent-of-code/utils"
	"fmt"
	"math"

	"github.com/emirpasic/gods/queues/priorityqueue"
	pqutil "github.com/emirpasic/gods/utils"
)

var UP = coordinate{-1, 0}
var RIGHT = coordinate{0, 1}
var DOWN = coordinate{1, 0}
var LEFT = coordinate{0, -1}

func main() {
	lines, _ := utils.ReadLines("2024/day16/input.txt")
	fmt.Println("2024 Day 16 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines))
}

func part1(lines []string) int {
	grid := linesToGrid(lines)
	start, end := grid.find('S'), grid.find('E')
	return getMinDistance(grid, start, end)
}

func getMinDistance(g grid, start, end coordinate) int {
	dirs := []coordinate{UP, RIGHT, DOWN, LEFT}
	minDist := math.MaxInt
	dist := make([][]int, len(g))
	for i := range dist {
		dist[i] = make([]int, len(g[0]))
		for j := range dist[i] {
			dist[i][j] = math.MaxInt
		}
	}
	dist[start.x][start.y] = 0
	pq := priorityqueue.NewWith(byDist)
	pq.Enqueue(state{start, LEFT, 0})
	for !pq.Empty() {
		val, _ := pq.Dequeue()
		curr := val.(state)
		if curr.pos == end && curr.dist < minDist {
			minDist = curr.dist
			continue
		}
		for _, dir := range dirs {
			nextPos := coordinate{curr.pos.x + dir.x, curr.pos.y + dir.y}
			nextDir := curr.dir
			nextDist := curr.dist
			if !isRange(g, nextPos.x, nextPos.y) || g[nextPos.x][nextPos.y] == '#' {
				continue
			}
			if nextDir == dir {
				nextDist++
				if nextDist < dist[nextPos.x][nextPos.y] {
					dist[nextPos.x][nextPos.y] = nextDist
					pq.Enqueue(state{nextPos, nextDir, nextDist})
				}
			} else {
				nextDirClockwise := nextDir
				nextDirCounter := nextDir
				nextDirClockwise.Rotate90DegreesClockwise()
				nextDirCounter.Rotate90DegreesCounterClockwise()
				nextDist += 1001
				if nextDist < dist[nextPos.x][nextPos.y] {
					dist[nextPos.x][nextPos.y] = nextDist
					pq.Enqueue(state{nextPos, nextDirClockwise, nextDist})
					pq.Enqueue(state{nextPos, nextDirCounter, nextDist})
				}
			}
		}
	}
	return minDist
}

type coordinate struct {
	x, y int
}

func (c *coordinate) Rotate90DegreesClockwise() {
	c.x, c.y = c.y, -c.x
}

func (c *coordinate) Rotate90DegreesCounterClockwise() {
	c.x, c.y = -c.y, c.x
}

type state struct {
	pos, dir coordinate
	dist     int
}

func byDist(a, b interface{}) int {
	distA := a.(state).dist
	distB := b.(state).dist
	return -pqutil.IntComparator(distA, distB)
}

type grid [][]rune

func (g grid) find(val rune) coordinate {
	for i, row := range g {
		for j, c := range row {
			if c == val {
				return coordinate{i, j}
			}
		}
	}
	return coordinate{-1, -1}
}

func isRange(grid [][]rune, i, j int) bool {
	return 0 <= i && i < len(grid) && 0 <= j && j < len(grid[0])
}

func linesToGrid(lines []string) grid {
	grid := grid{}
	for _, line := range lines {
		grid = append(grid, []rune(line))
	}
	return grid
}
