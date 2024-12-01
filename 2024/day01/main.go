package main

import (
	"advent-of-code/utils"
	"fmt"
	"sort"
	"strconv"
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
		left = append(left, strToInt(split[0]))
		right = append(right, strToInt(split[1]))
	}
	sort.Ints(left)
	sort.Ints(right)
	sum := 0
	for i := range left {
		sum += int(getAbs(right[i] - left[i]))
	}
	return sum
}

func part2(lines []string) int {
	left, right := []int{}, []int{}
	for _, line := range lines {
		split := strings.Fields(line)
		left = append(left, strToInt(split[0]))
		right = append(right, strToInt(split[1]))
	}
	freq := map[int]int{}
	for _, num := range left {
		freq[num] = 0
	}
	for _, num := range right {
		freq[num]++
	}
	score := 0
	for _, num := range left {
		score += num * freq[num]
	}
	return score
}

func strToInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

func getAbs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
