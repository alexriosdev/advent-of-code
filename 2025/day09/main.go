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
	return len(lines)
}

type Pos struct {
	X, Y int
}

func (a *Pos) Rectangle(b Pos) int {
	return (utils.Abs(a.X-b.X) + 1) * (utils.Abs(a.Y-b.Y) + 1)
}

type Positions []Pos

func (p *Positions) Add(pos Pos) {
	*p = append(*p, pos)
}

func (p *Positions) MaxRectangle() int {
	maxRect := 0
	visited := map[Pos]bool{}
	for i := 0; i < len(*p); i++ {
		for j := i + 1; j < len(*p); j++ {
			PosA := (*p)[i]
			PosB := (*p)[j]
			if visited[PosB] {
				continue
			}
			visited[PosA] = true
			maxRect = max(maxRect, PosA.Rectangle(PosB))
		}
	}
	return maxRect
}
