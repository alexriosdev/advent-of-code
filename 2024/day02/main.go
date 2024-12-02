package main

import (
	"advent-of-code/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines, _ := utils.ReadLines("input.txt")
	fmt.Println("2024 Day 02 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines))
}

func part1(lines []string) int {
	reports := [][]int{}
	for _, line := range lines {
		split := strings.Fields(line)
		levels := []int{}
		for _, str := range split {
			levels = append(levels, strToInt(str))
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

func isIncreasing(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] <= nums[i] {
			return false
		}
		diff := getAbs(nums[i+1] - nums[i])
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func isDecreasing(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] >= nums[i] {
			return false
		}
		diff := getAbs(nums[i+1] - nums[i])
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
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