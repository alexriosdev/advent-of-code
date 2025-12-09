package main

import (
	"advent-of-code/utils"
	"fmt"
	"math"
	"sort"
	"strings"
)

func main() {
	lines, _ := utils.ReadLines("2025/day08/input.txt")
	fmt.Println("2025 Day 08 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines, 1_000))
	fmt.Printf("Part 2: %v\n", part2(lines))
}

func part1(lines []string, n int) int {
	boxes := Boxes{}
	for _, line := range lines {
		split := strings.Split(line, ",")
		x, y, z := utils.StrToInt(split[0]), utils.StrToInt(split[1]), utils.StrToInt(split[2])
		boxes.Add(Box{x, y, z})
	}
	pairs := boxes.Pairs()
	dsu := NewDSU(len(pairs))
	freq := map[int]int{}
	for k := 0; k < n; k++ {
		i, j := pairs[k].Key[0], pairs[k].Key[1]
		dsu.Union(i, j)
	}
	for i := 0; i < len(pairs); i++ {
		freq[dsu.Find(i)]++
	}
	values := []int{}
	for _, v := range freq {
		values = append(values, v)
	}
	sort.Ints(values)
	values = values[len(values)-3:]
	return values[0] * values[1] * values[2]
}

func part2(lines []string) int {
	boxes := Boxes{}
	for _, line := range lines {
		split := strings.Split(line, ",")
		x, y, z := utils.StrToInt(split[0]), utils.StrToInt(split[1]), utils.StrToInt(split[2])
		boxes.Add(Box{x, y, z})
	}
	pairs := boxes.Pairs()
	dsu := NewDSU(len(pairs))
	connections := 0
	for _, pair := range pairs {
		i, j := pair.Key[0], pair.Key[1]
		if dsu.Find(i) != dsu.Find(j) {
			connections++
		}
		if connections == len(boxes)-1 {
			return boxes[i].X * boxes[j].X
		}
		dsu.Union(i, j)
	}
	return -1
}

type Pair struct {
	Key   [2]int
	Value float64
}

type Box struct {
	X, Y, Z int
}

func (a *Box) Dist(b *Box) float64 {
	dx := a.X - b.X
	dy := a.Y - b.Y
	dz := a.Z - b.Z
	squareSum := dx*dx + dy*dy + dz*dz
	return math.Sqrt(float64(squareSum))
}

type Boxes []Box

func (b *Boxes) Add(box Box) {
	*b = append(*b, box)
}

func (b *Boxes) Pairs() []Pair {
	pairs := []Pair{}
	visited := map[Pair]bool{}
	for i := 0; i < len(*b); i++ {
		for j := i + 1; j < len(*b); j++ {
			pairA := Pair{[2]int{i, j}, (*b)[i].Dist(&(*b)[j])}
			pairB := Pair{[2]int{j, i}, (*b)[j].Dist(&(*b)[i])}
			if visited[pairB] {
				continue
			}
			visited[pairA] = true
			pairs = append(pairs, pairA)
		}
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Value < pairs[j].Value
	})
	return pairs
}

type DSU struct {
	Parent []int
}

func NewDSU(n int) *DSU {
	dsu := &DSU{Parent: make([]int, n)}
	for i := 0; i < n; i++ {
		dsu.Parent[i] = i
	}
	return dsu
}

func (d *DSU) Find(i int) int {
	if d.Parent[i] == i {
		return i
	}
	d.Parent[i] = d.Find(d.Parent[i])
	return d.Parent[i]
}

func (d *DSU) Union(i, j int) bool {
	rootI := d.Find(i)
	rootJ := d.Find(j)
	if rootI == rootJ {
		return false
	}
	d.Parent[rootI] = rootJ
	return true
}
