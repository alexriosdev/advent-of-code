package main

import (
	"advent-of-code/utils"
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("2024 Day 05 Solution")
	fmt.Printf("Part 1: %v\n", part1(input))
}

func part1(input []byte) int {
	sections := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	orderingMap := map[int][]int{}
	for _, line := range strings.Fields(sections[0]) {
		split := strings.Split(line, "|")
		key, val := utils.StrToInt(split[0]), utils.StrToInt(split[1])
		orderingMap[key] = append(orderingMap[key], val)
	}
	allPages := [][]int{}
	for _, line := range strings.Fields(sections[1]) {
		pages := []int{}
		for _, s := range strings.Split(line, ",") {
			pages = append(pages, utils.StrToInt(s))
		}
		allPages = append(allPages, pages)
	}
	nums := []int{}
	for _, pages := range allPages {
		count, n := 0, len(pages)
		for i := 0; i < n-1; i++ {
			ordering := orderingMap[pages[i]]
			if sliceContains(ordering, pages[i+1]) {
				count++
			} else {
				break
			}
		}
		if count == n-1 {
			nums = append(nums, pages[n/2])
		}
	}
	result := 0
	for _, num := range nums {
		result += num
	}
	return result
}

func sliceContains(nums []int, target int) bool {
	for _, num := range nums {
		if num == target {
			return true
		}
	}
	return false
}
