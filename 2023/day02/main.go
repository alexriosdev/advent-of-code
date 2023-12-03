package main

import (
	"advent-of-code/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines, _ := utils.ReadLines("input.txt")
	fmt.Println("2023 Day 02 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines))
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
				split := strings.Split(strings.TrimSpace(cubeSet), " ")
				amount, _ := strconv.Atoi(split[0])
				color := split[1]
				if amount > cubeMap[color] && cubeMap[color] != 0 {
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
