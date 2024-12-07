package main

import (
	"advent-of-code/utils"
	"fmt"
	"strings"
)

func main() {
	lines, _ := utils.ReadLines("input.txt")
	fmt.Println("2024 Day 07 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines))
}

func part1(lines []string) int {
	operations := []operation{}
	for _, line := range lines {
		split := strings.Split(line, ":")
		operation := operation{utils.StrToInt(split[0]), []int{}}
		for _, s := range strings.Fields(split[1]) {
			operation.operands = append(operation.operands, utils.StrToInt(s))
		}
		operations = append(operations, operation)
	}
	sum := 0
	for _, operation := range operations {
		if findTotalWays(operation.operands, 0, 0, operation.result) > 0 {
			sum += operation.result
		}
	}
	return sum
}

type operation struct {
	result   int
	operands []int
}

func findTotalWays(nums []int, i, curr, target int) int {
	if i == len(nums) && curr == target {
		return 1
	}
	if i >= len(nums) {
		return 0
	}
	return findTotalWays(nums, i+1, curr+nums[i], target) + findTotalWays(nums, i+1, curr*nums[i], target)
}
