package main

import (
	"advent-of-code/utils"
	"fmt"
	"math"
	"strings"
)

func main() {
	lines, _ := utils.ReadLines("2023/day04/input.txt")
	fmt.Println("2023 Day 04 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines))
	fmt.Printf("Part 2: %v\n", part2(lines))
}

func part1(lines []string) int {
	result := 0
	for _, line := range lines {
		line = strings.Split(line, ":")[1]
		numbers := strings.Split(line, "|")
		winningSet := map[string]bool{}
		for _, num := range strings.Fields(numbers[0]) {
			winningSet[num] = true
		}
		count := 0
		for _, num := range strings.Fields(numbers[1]) {
			if winningSet[num] {
				count++
			}
		}
		if count != 0 {
			result += powInt(2, count-1)
		}
	}
	return result
}

func part2(lines []string) int {
	result := 0
	cards := make([]int, len(lines))
	for i, line := range lines {
		line = strings.Split(line, ":")[1]
		numbers := strings.Split(line, "|")
		winningSet := map[string]bool{}
		for _, num := range strings.Fields(numbers[0]) {
			winningSet[num] = true
		}
		cards[i]++
		count := 0
		for _, num := range strings.Fields(numbers[1]) {
			if winningSet[num] {
				count++
				cards[i+count] += cards[i]
			}
		}
		result += cards[i]
	}
	return result
}

func powInt(base, exp int) int {
	return int(math.Pow(float64(base), float64(exp)))
}
