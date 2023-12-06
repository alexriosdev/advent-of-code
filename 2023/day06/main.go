package main

import (
	"advent-of-code/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines, _ := utils.ReadLines("input.txt")
	fmt.Println("2023 Day 06 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines))
}

func part1(lines []string) int {
	result := 1
	times := strings.Fields(strings.Split(lines[0], ":")[1])
	dists := strings.Fields(strings.Split(lines[1], ":")[1])
	for i := range times {
		count := 0
		time, dist, availableDist := strToInt(times[i]), strToInt(dists[i]), strToInt(times[i])
		for currTime := 0; currTime < time; currTime++ {
			if currTime*availableDist > dist {
				count++
			}
			availableDist--
		}
		if count != 0 {
			result *= count
		}
	}
	return result
}

func strToInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}
