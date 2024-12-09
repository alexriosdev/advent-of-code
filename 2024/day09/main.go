package main

import (
	"advent-of-code/utils"
	"fmt"
	"slices"
)

func main() {
	lines, _ := utils.ReadLines("input_test.txt")
	fmt.Println("2024 Day 09 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines))
	fmt.Printf("Part 2: %v\n", part2(lines))
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

func part2(lines []string) int {
	blocks := getBlocks(lines)
	performSwap(&blocks)
	result, idx := 0, 0
	for _, block := range blocks {
		if block.val == -1 {
			idx += block.size
			continue
		}
		for i := 0; i < block.size; i++ {
			result += idx * block.val
			idx++
		}
	}
	return result
}

type block struct {
	val, size int
}

func getBlocks(lines []string) []block {
	blocks := []block{}
	for _, line := range lines {
		for i, c := range line {
			num := -1
			if i%2 == 0 {
				num = i / 2
			}
			blocks = append(blocks, block{num, int(c - '0')})
		}
	}
	return blocks
}

func performSwap(blocks *[]block) {
	for i := len(*blocks) - 1; i > 0; i-- {
		if (*blocks)[i].val == -1 {
			continue
		}
		for j := 0; j < i; j++ {
			if (*blocks)[j].val == -1 && (*blocks)[j].size >= (*blocks)[i].size {
				n := (*blocks)[j].size - (*blocks)[i].size
				(*blocks)[j] = block{(*blocks)[i].val, (*blocks)[i].size}
				(*blocks)[i].val = -1
				*blocks = slices.Insert(*blocks, j+1, block{-1, n})
				break
			}
		}
	}
}

func repeatNum(num, n int) []int {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = num
	}
	return nums
}
