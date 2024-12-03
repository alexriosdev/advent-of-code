package main

import (
	"advent-of-code/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	lines, _ := utils.ReadLines("input.txt")
	fmt.Println("2024 Day 03 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines))
}

func part1(lines []string) int {
	result := 0
	pattern := regexp.MustCompile(`mul\(\d+,\d+\)`)
	replacer := strings.NewReplacer("mul(", "", ",", " ", ")", "")
	for _, line := range lines {
		matches := pattern.FindAllString(line, -1)
		for _, match := range matches {
			split := strings.Fields(replacer.Replace(match))
			result += strToInt(split[0]) * strToInt(split[1])
		}
	}
	return result
}

func strToInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}
