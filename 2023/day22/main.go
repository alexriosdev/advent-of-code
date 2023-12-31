package main

import (
	"advent-of-code/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/emirpasic/gods/sets/hashset"
)

type brick struct {
	x1, y1, z1, x2, y2, z2 int
}

func main() {
	lines, _ := utils.ReadLines("input.txt")
	fmt.Println("2023 Day 22 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines))
}

func part1(lines []string) int {
	replacer := strings.NewReplacer(",", " ", "~", " ")
	bricks := []brick{}
	for _, line := range lines {
		split := strings.Fields(replacer.Replace(line))
		brick := brick{
			x1: strToInt(split[0]),
			y1: strToInt(split[1]),
			z1: strToInt(split[2]),
			x2: strToInt(split[3]),
			y2: strToInt(split[4]),
			z2: strToInt(split[5]),
		}
		bricks = append(bricks, brick)
	}
	sort.Slice(bricks, func(i, j int) bool {
		return bricks[i].z1 < bricks[j].z1
	})
	for i, a := range bricks {
		maxZ := 1
		for _, b := range bricks[:i] {
			if isIntersect(a, b) {
				maxZ = max(maxZ, b.z2+1)
			}
		}
		bricks[i].z2 = maxZ + bricks[i].z2 - bricks[i].z1
		bricks[i].z1 = maxZ
	}
	sort.Slice(bricks, func(i, j int) bool {
		return bricks[i].z1 < bricks[j].z1
	})
	aSupportsB := make([]*hashset.Set, len(bricks))
	bSupportsA := make([]*hashset.Set, len(bricks))
	for i := 0; i < len(bricks); i++ {
		aSupportsB[i] = hashset.New()
		bSupportsA[i] = hashset.New()
	}
	for i, a := range bricks {
		for j, b := range bricks[:i] {
			if isIntersect(a, b) && a.z1 == b.z2+1 {
				aSupportsB[j].Add(i)
				bSupportsA[i].Add(j)
			}
		}
	}
	result := 0
	for i := 0; i < len(bricks); i++ {
		if isDisintegrate(i, aSupportsB, bSupportsA) {
			result++
		}
	}
	return result
}

func isIntersect(a, b brick) bool {
	return max(a.x1, b.x1) <= min(a.x2, b.x2) && max(a.y1, b.y1) <= min(a.y2, b.y2)
}

func isDisintegrate(i int, aSupportsB, bSupportsA []*hashset.Set) bool {
	for _, j := range aSupportsB[i].Values() {
		if bSupportsA[j.(int)].Size() < 2 {
			return false
		}
	}
	return true
}

func strToInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
