package main

import (
	"advent-of-code/utils"
	"fmt"
	"unicode"
)

func main() {
	lines, _ := utils.ReadLines("input.txt")
	fmt.Println("2023 Day 01 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines))
	fmt.Printf("Part 2: %v\n", part2(lines))
}

func part1(lines []string) int {
	result := 0
	for _, s := range lines {
		first, last := 0, 0
		for _, c := range s {
			if unicode.IsDigit(c) {
				first, last = updateDigits(first, last, int(c-'0'))
			}
		}
		result += (first * 10) + last
	}
	return result
}

func part2(lines []string) int {
	result := 0
	numberMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	for _, s := range lines {
		first, last := 0, 0
		for i, c := range s {
			n := len(s)
			if unicode.IsDigit(c) {
				first, last = updateDigits(first, last, int(c-'0'))
				continue
			}
			if i < n-2 && numberMap[s[i:i+3]] != 0 {
				first, last = updateDigits(first, last, numberMap[s[i:i+3]])
				continue
			}
			if i < n-3 && numberMap[s[i:i+4]] != 0 {
				first, last = updateDigits(first, last, numberMap[s[i:i+4]])
				continue
			}
			if i < n-4 && numberMap[s[i:i+5]] != 0 {
				first, last = updateDigits(first, last, numberMap[s[i:i+5]])
				continue
			}
		}
		result += (first * 10) + last
	}
	return result
}

func updateDigits(first, last, val int) (int, int) {
	if first == 0 {
		first = val
	}
	last = val
	return first, last
}
