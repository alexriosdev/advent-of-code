package main

import (
	"advent-of-code/utils"
	"fmt"
)

func main() {
	lines, _ := utils.ReadLines("input_test.txt")
	fmt.Println("2024 Day 09 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines))
}

func part1(lines []string) int {
	blocks := []int{}
	for _, line := range lines {
		for i, c := range line {
			num := -1
			if i%2 == 0 {
				num = i / 2
			}
			blocks = append(blocks, repeatNum(num, int(c-'0'))...)
		}
	}
	left, right := 0, len(blocks)-1
	for left < right {
		for blocks[left] != -1 && left < right {
			left++
		}
		for blocks[right] == -1 && left < right {
			right--
		}
		blocks[left], blocks[right] = blocks[right], blocks[left]
	}
	result := 0
	for i, block := range blocks {
		if block == -1 {
			break
		}
		result += i * block
	}
	return result
}

func repeatNum(num, n int) []int {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = num
	}
	return nums
}
