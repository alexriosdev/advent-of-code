package main

import (
	"advent-of-code/utils"
	"fmt"
	"sort"
	"strings"
)

func main() {
	lines, _ := utils.ReadLines("input.txt")
	fmt.Println("2024 Day 01 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines))
	fmt.Printf("Part 2: %v\n", part2(lines))
}

func part1(lines []string) int {
	left, right := []int{}, []int{}
	for _, line := range lines {
		split := strings.Fields(line)
		left = append(left, utils.StrToInt(split[0]))
		right = append(right, utils.StrToInt(split[1]))
	}
	sort.Ints(left)
	sort.Ints(right)
	sum := 0
	for i := range left {
		sum += utils.Abs(right[i] - left[i])
	}
	return sum
}

func part2(lines []string) int {
	left := []int{}
	freq := map[int]int{}
	for _, line := range lines {
		split := strings.Fields(line)
		left = append(left, utils.StrToInt(split[0]))
		freq[utils.StrToInt(split[1])]++
	}
	score := 0
	for _, num := range left {
		score += num * freq[num]
	}
	return score
}
