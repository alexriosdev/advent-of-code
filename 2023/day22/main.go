package main

import (
	"advent-of-code/utils"
	"fmt"
	"sort"
	"strings"

	"github.com/emirpasic/gods/queues/linkedlistqueue"
	"github.com/emirpasic/gods/sets/hashset"
)

type brick struct {
	x1, y1, z1, x2, y2, z2 int
}

func main() {
	lines, _ := utils.ReadLines("input.txt")
	fmt.Println("2023 Day 22 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines))
	fmt.Printf("Part 2: %v\n", part2(lines))
}

func part1(lines []string) int {
	replacer := strings.NewReplacer(",", " ", "~", " ")
	bricks := []brick{}
	for _, line := range lines {
		split := strings.Fields(replacer.Replace(line))
		brick := brick{
			x1: utils.StrToInt(split[0]),
			y1: utils.StrToInt(split[1]),
			z1: utils.StrToInt(split[2]),
			x2: utils.StrToInt(split[3]),
			y2: utils.StrToInt(split[4]),
			z2: utils.StrToInt(split[5]),
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

func part2(lines []string) int {
	replacer := strings.NewReplacer(",", " ", "~", " ")
	bricks := []brick{}
	for _, line := range lines {
		split := strings.Fields(replacer.Replace(line))
		brick := brick{
			x1: utils.StrToInt(split[0]),
			y1: utils.StrToInt(split[1]),
			z1: utils.StrToInt(split[2]),
			x2: utils.StrToInt(split[3]),
			y2: utils.StrToInt(split[4]),
			z2: utils.StrToInt(split[5]),
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
		queue := linkedlistqueue.New()
		fallenBricks := hashset.New()
		fallenBricks.Add(i)
		for _, j := range aSupportsB[i].Values() {
			if bSupportsA[j.(int)].Size() < 2 {
				queue.Enqueue(j.(int))
				fallenBricks.Add(j.(int))
			}
		}
		for !queue.Empty() {
			val, _ := queue.Dequeue()
			j := val.(int)
			for _, k := range getDifference(aSupportsB[j], fallenBricks).Values() {
				if isSubset(fallenBricks, bSupportsA[k.(int)]) {
					queue.Enqueue(k.(int))
					fallenBricks.Add(k.(int))
				}
			}
		}
		result += fallenBricks.Size() - 1
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

func isSubset(a, b *hashset.Set) bool {
	return a.Contains(b.Values()...)
}

func getDifference(a, b *hashset.Set) *hashset.Set {
	c := hashset.New()
	c.Add(a.Values()...)
	c.Remove(b.Values()...)
	return c
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
