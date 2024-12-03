package main

import (
	"advent-of-code/utils"
	"fmt"
	"strings"
)

func main() {
	lines, _ := utils.ReadLines("input.txt")
	fmt.Println("2023 Day 06 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines))
	fmt.Printf("Part 2: %v\n", part2(lines))
}

func part1(lines []string) int {
	result := 1
	times := strings.Fields(strings.Split(lines[0], ":")[1])
	dists := strings.Fields(strings.Split(lines[1], ":")[1])
	for i := range times {
		count := 0
		time, dist, availableDist := utils.StrToInt(times[i]), utils.StrToInt(dists[i]), utils.StrToInt(times[i])
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

func part2(lines []string) int {
	result := 1
	count := 0
	time, dist, availableDist := buildNumber(lines[0]), buildNumber(lines[1]), buildNumber(lines[0])
	for currTime := 0; currTime < time; currTime++ {
		if currTime*availableDist > dist {
			count++
		}
		availableDist--
	}
	if count != 0 {
		result *= count
	}
	return result
}

func buildNumber(line string) int {
	sb := strings.Builder{}
	for _, s := range strings.Fields(strings.Split(line, ":")[1]) {
		sb.WriteString(s)
	}
	return utils.StrToInt(sb.String())
}
