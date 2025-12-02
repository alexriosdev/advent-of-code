package main

import (
	"advent-of-code/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	lines, _ := utils.ReadLines("2024/day11/input.txt")
	fmt.Println("2024 Day 11 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines, 25))
	fmt.Printf("Part 2: %v\n", part1(lines, 75))
}

func part1(lines []string, n int) int {
	stones := getStones(lines)
	for i := 0; i < n; i++ {
		rearrangeStones(&stones)
	}
	result := 0
	for _, count := range stones {
		result += count
	}
	return result
}

func getStones(lines []string) map[int]int {
	stones := map[int]int{}
	for _, line := range lines {
		for _, s := range strings.Fields(line) {
			stones[utils.StrToInt(s)] = 1
		}
	}
	return stones
}

func rearrangeStones(stones *map[int]int) {
	newStones := map[int]int{}
	for stone, count := range *stones {
		switch {
		case stone == 0:
			newStones[1] += count
		case getDigits(stone)%2 == 0:
			s := strconv.Itoa(stone)
			a, b := utils.StrToInt(s[:len(s)/2]), utils.StrToInt(s[len(s)/2:])
			newStones[a] += count
			newStones[b] += count
		default:
			newStones[stone*2024] += count
		}
	}
	*stones = newStones
}

func getDigits(num int) int {
	return int(math.Floor(math.Log10(float64(num))) + 1)
}
