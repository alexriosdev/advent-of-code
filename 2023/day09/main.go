package main

import (
	"advent-of-code/utils"
	"fmt"
	"strings"
)

func main() {
	lines, _ := utils.ReadLines("input.txt")
	fmt.Println("2023 Day 09 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines))
	fmt.Printf("Part 2: %v\n", part2(lines))
}

func part1(lines []string) int {
	result := 0
	for _, line := range lines {
		nums := lineToNumbers(line)
		predictions := [][]int{}
		predictions = append(predictions, nums)
		computePredictions(&nums, &predictions)
		result += computeLast(&predictions)
	}
	return result
}

func part2(lines []string) int {
	result := 0
	for _, line := range lines {
		nums := lineToNumbers(line)
		predictions := [][]int{}
		predictions = append(predictions, nums)
		computePredictions(&nums, &predictions)
		result += computeFirst(&predictions)
	}
	return result
}

func computeFirst(predictions *[][]int) int {
	first := 0
	for i := len(*predictions) - 1; i >= 0; i-- {
		first = (*predictions)[i][0] - first
	}
	return first
}

func computeLast(predictions *[][]int) int {
	last := 0
	for i := len(*predictions) - 1; i >= 0; i-- {
		n := len((*predictions)[i]) - 1
		last += (*predictions)[i][n]
	}
	return last
}

func computePredictions(nums *[]int, predictions *[][]int) {
	if isZeroes(*nums) {
		return
	}
	temp := []int{}
	for i := 1; i < len(*nums); i++ {
		temp = append(temp, (*nums)[i]-(*nums)[i-1])
	}
	*predictions = append(*predictions, temp)
	computePredictions(&temp, predictions)
}

func isZeroes(nums []int) bool {
	for _, num := range nums {
		if num != 0 {
			return false
		}
	}
	return true
}

func lineToNumbers(line string) []int {
	nums := []int{}
	for _, s := range strings.Fields(line) {
		nums = append(nums, utils.StrToInt(s))
	}
	return nums
}
