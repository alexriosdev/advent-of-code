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
	fmt.Printf("Part 2: %v\n", part2(lines))
}

func part1(lines []string) int {
	pattern := regexp.MustCompile(`mul\(\d+,\d+\)`)
	replacer := strings.NewReplacer("mul(", "", ",", " ", ")", "")
	result := 0
	for _, match := range pattern.FindAllString(strings.Join(lines, ""), -1) {
		split := strings.Fields(replacer.Replace(match))
		result += strToInt(split[0]) * strToInt(split[1])
	}
	return result
}

func part2(lines []string) int {
	pattern := regexp.MustCompile(`mul\(\d+,\d+\)`)
	replacer := strings.NewReplacer("mul(", "", ",", " ", ")", "")
	enabledPattern := regexp.MustCompile(`(do\(\))(.*?)(don't\(\)|$)`)
	firstPattern := regexp.MustCompile(`(mul\(\d+,\d+\))(.*?)(do\(\)|don't\(\))`)
	join := strings.Join(lines, "")
	first := firstPattern.FindString(join)
	matches := pattern.FindAllString(first, -1)
	for _, enabled := range enabledPattern.FindAllString(join, -1) {
		matches = append(matches, pattern.FindAllString(enabled, -1)...)
	}
	result := 0
	for _, match := range matches {
		split := strings.Fields(replacer.Replace(match))
		result += strToInt(split[0]) * strToInt(split[1])
	}
	return result
}

func strToInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}
