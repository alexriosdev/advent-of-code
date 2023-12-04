package main

import (
	"advent-of-code/utils"
	"fmt"
	"math"
	"strings"
)

func main() {
	lines, _ := utils.ReadLines("input.txt")
	fmt.Println("2023 Day 04 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines))
	fmt.Printf("Part 2: %v\n", part2(lines))
}

func part1(lines []string) int {
	result := 0
	for _, line := range lines {
		line = strings.Split(line, ":")[1]
		cards := strings.Split(line, "|")
		winningSet := map[string]bool{}
		for _, num := range getNumbers(cards[0]) {
			winningSet[num] = true
		}
		count := 0
		for _, num := range getNumbers(cards[1]) {
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
	cardFreq := make([]int, len(lines))
	for i := range cardFreq {
		cardFreq[i] = 1
	}
	for i, line := range lines {
		line = strings.Split(line, ":")[1]
		cards := strings.Split(line, "|")
		winningSet := map[string]bool{}
		for _, num := range getNumbers(cards[0]) {
			winningSet[num] = true
		}
		count := 0
		for _, num := range getNumbers(cards[1]) {
			if winningSet[num] {
				count++
				cardFreq[i+count] += cardFreq[i]
			}
		}
	}
	result := 0
	for _, val := range cardFreq {
		result += val
	}
	return result
}

func getNumbers(cards string) []string {
	nums := []string{}
	for _, num := range strings.Fields(cards) {
		nums = append(nums, num)
	}
	return nums
}

func powInt(base, exp int) int {
	return int(math.Pow(float64(base), float64(exp)))
}
