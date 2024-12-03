package main

import (
	"advent-of-code/utils"
	"fmt"
	"strings"
)

func main() {
	lines, _ := utils.ReadLines("input.txt")
	fmt.Println("2024 Day 02 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines))
	fmt.Printf("Part 2: %v\n", part2(lines))
}

func part1(lines []string) int {
	reports := [][]int{}
	for _, line := range lines {
		levels := []int{}
		for _, str := range strings.Fields(line) {
			levels = append(levels, utils.StrToInt(str))
		}
		reports = append(reports, levels)
	}
	count := 0
	for _, levels := range reports {
		if isIncreasing(levels) || isDecreasing(levels) {
			count++
		}
	}
	return count
}

func part2(lines []string) int {
	reports := [][]int{}
	for _, line := range lines {
		levels := []int{}
		for _, str := range strings.Fields(line) {
			levels = append(levels, utils.StrToInt(str))
		}
		reports = append(reports, levels)
	}
	count := 0
	for _, levels := range reports {
		if isIncreasing(levels) || isDecreasing(levels) {
			count++
			continue
		}
		for i := range levels {
			newLevels := removeIndex(levels, i)
			if isIncreasing(newLevels) || isDecreasing(newLevels) {
				count++
				break
			}
		}
	}
	return count
}

func isIncreasing(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		diff := utils.Abs(nums[i+1] - nums[i])
		if nums[i+1] <= nums[i] || !isRange(1, 3, diff) {
			return false
		}
	}
	return true
}

func isDecreasing(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		diff := utils.Abs(nums[i+1] - nums[i])
		if nums[i+1] >= nums[i] || !isRange(1, 3, diff) {
			return false
		}
	}
	return true
}

func isRange(a, b, num int) bool {
	return a <= num && num <= b
}

func removeIndex(nums []int, idx int) []int {
	result := []int{}
	for i, num := range nums {
		if i == idx {
			continue
		}
		result = append(result, num)
	}
	return result
}
