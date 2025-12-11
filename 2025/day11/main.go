package main

import (
	"advent-of-code/utils"
	"fmt"
	"strings"
)

func main() {
	lines, _ := utils.ReadLines("2025/day11/input.txt")
	fmt.Println("2025 Day 11 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines))
	fmt.Printf("Part 2: %v\n", part2(lines))
}

func part1(lines []string) int {
	dict := map[string][]string{}
	for _, line := range lines {
		fields := strings.Fields(strings.NewReplacer(":", "").Replace(line))
		key := fields[0]
		dict[key] = fields[1:]
	}
	queue := []string{}
	for _, v := range dict["you"] {
		queue = append(queue, v)
	}
	count := 0
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr == "out" {
			count++
		}
		for _, v := range dict[curr] {
			queue = append(queue, v)
		}
	}
	return count
}

func part2(lines []string) int {
	return len(lines)
}
