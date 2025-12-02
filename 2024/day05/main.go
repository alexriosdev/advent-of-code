package main

import (
	"advent-of-code/utils"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	input, _ := os.ReadFile("2024/day05/input.txt")
	fmt.Println("2024 Day 05 Solution")
	fmt.Printf("Part 1: %v\n", part1(input))
	fmt.Printf("Part 2: %v\n", part2(input))
}

func part1(input []byte) int {
	sections := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	orderingMap := getOrderingMap(sections[0])
	allPages := getAllPages(sections[1])
	nums := []int{}
	for _, pages := range allPages {
		count, n := 0, len(pages)
		for i := 0; i < n-1; i++ {
			ordering := orderingMap[pages[i]]
			if !slices.Contains(ordering, pages[i+1]) {
				break
			}
			count++
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

func part2(input []byte) int {
	sections := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	orderingMap := getOrderingMap(sections[0])
	allPages := getAllPages(sections[1])
	incorrectPages := [][]int{}
	for _, pages := range allPages {
		for i := 0; i < len(pages)-1; i++ {
			ordering := orderingMap[pages[i]]
			if !slices.Contains(ordering, pages[i+1]) {
				incorrectPages = append(incorrectPages, pages)
				break
			}
		}
	}
	nums := []int{}
	for _, pages := range incorrectPages {
		isComplete := false
		for !isComplete {
			if isOrdered(&pages, &orderingMap) {
				isComplete = true
			}
		}
		nums = append(nums, pages[len(pages)/2])
	}
	result := 0
	for _, num := range nums {
		result += num
	}
	return result
}

func getOrderingMap(section string) map[int][]int {
	orderingMap := map[int][]int{}
	for _, line := range strings.Fields(section) {
		split := strings.Split(line, "|")
		key, val := utils.StrToInt(split[0]), utils.StrToInt(split[1])
		orderingMap[key] = append(orderingMap[key], val)
	}
	return orderingMap
}

func getAllPages(section string) [][]int {
	allPages := [][]int{}
	for _, line := range strings.Fields(section) {
		pages := []int{}
		for _, s := range strings.Split(line, ",") {
			pages = append(pages, utils.StrToInt(s))
		}
		allPages = append(allPages, pages)
	}
	return allPages
}

func isOrdered(pages *[]int, orderingMap *map[int][]int) bool {
	for start, ordering := range *orderingMap {
		for _, end := range ordering {
			i, j := slices.Index(*pages, start), slices.Index(*pages, end)
			if i != -1 && j != -1 && i >= j {
				shift := append((*pages)[:j], (*pages)[j+1:]...)
				shift = append(shift, end)
				*pages = shift
				return false
			}
		}
	}
	return true
}
