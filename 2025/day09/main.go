package main

import (
	"advent-of-code/utils"
	"fmt"
	"strings"
)

func main() {
	lines, _ := utils.ReadLines("2025/day09/input.txt")
	fmt.Println("2025 Day 09 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines))
	fmt.Printf("Part 2: %v\n", part2(lines))
}

func part1(lines []string) int {
	positions := Positions{}
	for _, line := range lines {
		split := strings.Split(line, ",")
		pos := Pos{utils.StrToInt(split[0]), utils.StrToInt(split[1])}
		positions.Add(pos)
	}
	return positions.MaxRectangle()
}

func part2(lines []string) int {
	positions := Positions{}
	for _, line := range lines {
		split := strings.Split(line, ",")
		pos := Pos{utils.StrToInt(split[0]), utils.StrToInt(split[1])}
		positions.Add(pos)
	}
	return positions.MaxRectangleInBounds()
}

type Pos struct {
	X, Y int
}

func Rectangle(a, b Pos) int {
	return (utils.Abs(a.X-b.X) + 1) * (utils.Abs(a.Y-b.Y) + 1)
}

func IsOverlap(a, b, c, d Pos) bool {
	minX, maxX := min(a.X, b.X), max(a.X, b.X)
	minY, maxY := min(a.Y, b.Y), max(a.Y, b.Y)
	xOverlap := maxX > min(c.X, d.X) && minX < max(c.X, d.X)
	yOverlap := maxY > min(c.Y, d.Y) && minY < max(c.Y, d.Y)
	return xOverlap && yOverlap
}

type Positions []Pos

func (p *Positions) Add(pos Pos) {
	*p = append(*p, pos)
}

func (p *Positions) MaxRectangle() int {
	maxRect := 0
	for i := 0; i < len(*p); i++ {
		for j := i + 1; j < len(*p); j++ {
			a, b := (*p)[i], (*p)[j]
			maxRect = max(maxRect, Rectangle(a, b))
		}
	}
	return maxRect
}

func (p *Positions) MaxRectangleInBounds() int {
	maxRect := 0
	for i := 0; i < len(*p); i++ {
		for j := i + 1; j < len(*p); j++ {
			a, b := (*p)[i], (*p)[j]
			if (*p).IsInBounds(a, b) {
				maxRect = max(maxRect, Rectangle(a, b))
			}
		}
	}
	return maxRect
}

func (p *Positions) IsInBounds(a, b Pos) bool {
	for i := 0; i < len(*p); i++ {
		j := (i + 1) % len(*p)
		c, d := (*p)[i], (*p)[j]
		if IsOverlap(a, b, c, d) {
			return false
		}
	}
	return true
}
