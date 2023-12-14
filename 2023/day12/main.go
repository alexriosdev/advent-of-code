package main

import (
	"advent-of-code/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines, _ := utils.ReadLines("input.txt")
	fmt.Println("2023 Day 12 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines))
}

func part1(lines []string) int {
	sum := 0
	for _, line := range lines {
		split := strings.Split(line, " ")
		runes, nums := []rune(split[0]), strToNumbers(split[1])
		cache := map[[3]int]int{}
		sum += arrangementCount(cache, runes, nums, 0, 0, 0)
	}
	return sum
}

func arrangementCount(cache map[[3]int]int, runes []rune, nums []int, i, j, num int) int {
	if i >= len(runes) {
		if j >= len(nums) || (j == len(nums)-1 && nums[j] == num) {
			return 1
		}
		return 0
	}
	switch runes[i] {
	case '.':
		if num == 0 {
			return arrangementCount(cache, runes, nums, i+1, j, num)
		}
		if j >= len(nums) || num != nums[j] {
			return 0
		}
		return arrangementCount(cache, runes, nums, i+1, j+1, 0)
	case '#':
		if j >= len(nums) || num+1 > nums[j] {
			return 0
		}
		return arrangementCount(cache, runes, nums, i+1, j, num+1)
	case '?':
		key := [3]int{i, j, num}
		if val, ok := cache[key]; ok {
			return val
		}
		count := 0
		if num == 0 {
			count += arrangementCount(cache, runes, nums, i+1, j, num)
		}
		if j < len(nums) && num < nums[j] {
			count += arrangementCount(cache, runes, nums, i+1, j, num+1)
		}
		if j < len(nums) && num == nums[j] {
			count += arrangementCount(cache, runes, nums, i+1, j+1, 0)
		}
		cache[key] = count
		return count
	}
	return 0
}

func strToNumbers(s string) []int {
	nums := []int{}
	for _, num := range strings.Split(s, ",") {
		nums = append(nums, strToInt(num))
	}
	return nums
}

func strToInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}
