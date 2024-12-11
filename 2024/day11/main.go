package main

import (
	"advent-of-code/utils"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func main() {
	lines, _ := utils.ReadLines("input.txt")
	fmt.Println("2024 Day 11 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines, 25))
}

func part1(lines []string, n int) int {
	stones := getStones(lines)
	for i := 0; i < n; i++ {
		rearrangeStones(&stones)
	}
	return len(stones)
}

func getStones(lines []string) []int {
	stones := []int{}
	for _, line := range lines {
		for _, s := range strings.Fields(line) {
			stones = append(stones, utils.StrToInt(s))
		}
	}
	return stones
}

func rearrangeStones(stones *[]int) {
	for i := 0; i < len(*stones); i++ {
		switch {
		case (*stones)[i] == 0:
			(*stones)[i] = 1
		case getDigits((*stones)[i])%2 == 0:
			s := strconv.Itoa((*stones)[i])
			a, b := utils.StrToInt(s[:len(s)/2]), utils.StrToInt(s[len(s)/2:])
			(*stones)[i] = a
			*stones = slices.Insert(*stones, i+1, b)
			i++
		default:
			(*stones)[i] *= 2024
		}
	}
}

func getDigits(num int) int {
	return int(math.Floor(math.Log10(float64(num))) + 1)
}
