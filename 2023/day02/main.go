package main

import (
	"advent-of-code/utils"
	"fmt"
	"strings"
)

func main() {
	lines, _ := utils.ReadLines("2023/day02/input.txt")
	fmt.Println("2023 Day 02 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines))
	fmt.Printf("Part 2: %v\n", part2(lines))
}

func part1(lines []string) int {
	cubeMap := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	result := 0
	for i, line := range lines {
		isValid := true
		newLine := strings.Split(line, ":")[1]
		gameSets := strings.Split(newLine, ";")
		for _, gameSet := range gameSets {
			if !isValid {
				break
			}
			cubeSets := strings.Split(gameSet, ",")
			for _, cubeSet := range cubeSets {
				split := strings.Fields(cubeSet)
				amount := utils.StrToInt(split[0])
				color := split[1]
				if amount > cubeMap[color] {
					isValid = false
					break
				}
			}
		}
		if isValid {
			result += i + 1
		}
	}
	return result
}

func part2(lines []string) int {
	cubeMap := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	result := 0
	for _, line := range lines {
		newLine := strings.Split(line, ":")[1]
		gameSets := strings.Split(newLine, ";")
		for _, gameSet := range gameSets {
			cubeSets := strings.Split(gameSet, ",")
			for _, cubeSet := range cubeSets {
				split := strings.Fields(cubeSet)
				amount := utils.StrToInt(split[0])
				color := split[1]
				if amount > cubeMap[color] {
					cubeMap[color] = amount
				}
			}
		}
		power := 1
		for k, v := range cubeMap {
			power *= v
			cubeMap[k] = 0
		}
		result += power
	}
	return result
}
