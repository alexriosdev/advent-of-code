package main

import (
	"advent-of-code/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	lines, _ := utils.ReadLines("input.txt")
	fmt.Println("2024 Day 07 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines))
	fmt.Printf("Part 2: %v\n", part2(lines))
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

func part2(lines []string) int {
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
		if findTotalWaysConcat(operation.operands, 0, 0, operation.result) > 0 {
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
	return findTotalWays(nums, i+1, curr+nums[i], target) +
		findTotalWays(nums, i+1, curr*nums[i], target)
}

func findTotalWaysConcat(nums []int, i, curr, target int) int {
	if i == len(nums) && curr == target {
		return 1
	}
	if i >= len(nums) {
		return 0
	}
	return findTotalWaysConcat(nums, i+1, curr+nums[i], target) +
		findTotalWaysConcat(nums, i+1, curr*nums[i], target) +
		findTotalWaysConcat(nums, i+1, numConcat(curr, nums[i]), target)
}

func numConcat(a, b int) int {
	n := len(strconv.Itoa(b))
	a *= int(math.Pow10(n))
	return a + b
}
