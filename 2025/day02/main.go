package main

import (
	"advent-of-code/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	line, _ := utils.ReadLine("2025/day02/input.txt")
	fmt.Println("2025 Day 02 Solution")
	fmt.Printf("Part 1: %v\n", part1(line))
	fmt.Printf("Part 2: %v\n", part2(line))
}

func part1(line string) int {
	sum := 0
	for _, s := range strings.Split(line, ",") {
		split := strings.Split(s, "-")
		start, end := utils.StrToInt(split[0]), utils.StrToInt(split[1])
		for num := start; num <= end; num++ {
			if isExactlyTwice(num) {
				sum += num
			}
		}
	}
	return sum
}

func part2(line string) int {
	sum := 0
	for _, s := range strings.Split(line, ",") {
		split := strings.Split(s, "-")
		start, end := utils.StrToInt(split[0]), utils.StrToInt(split[1])
		for num := start; num <= end; num++ {
			if isAtLeastTwice(num) {
				sum += num
			}
		}
	}
	return sum
}

func isExactlyTwice(num int) bool {
	n := countDigits(num)
	if n%2 != 0 {
		return false
	}
	a, b := splitNumber(num, n)
	return a == b
}

func isAtLeastTwice(num int) bool {
	n := countDigits(num)
	if n >= 10 {
		joined := strconv.Itoa(num) + strconv.Itoa(num)
		joined = joined[1 : len(joined)-1]
		return strings.Contains(joined, strconv.Itoa(num))
	}
	joined := joinNumbers(num, num)
	joined = removeFirstAndLastDigit(joined)
	return containsNumber(joined, num)
}

func countDigits(num int) int {
	if num == 0 {
		return 1
	}
	num = utils.Abs(num)
	return int(math.Floor(math.Log10(float64(num)))) + 1
}

func removeFirstAndLastDigit(num int) int {
	if num < 100 {
		return 0
	}
	num /= 10
	power := getPower(num) / 10
	return num % power
}

func containsNumber(a, b int) bool {
	if a == b {
		return true
	}
	if a < b {
		return false
	}
	power := getPower(b)
	for i := a; i >= b; i /= 10 {
		if (i % power) == b {
			return true
		}
	}
	return false
}

func splitNumber(num, n int) (int, int) {
	divisor := int(math.Pow(10, float64(n/2)))
	return num % divisor, num / divisor
}

func joinNumbers(a, b int) int {
	power := getPower(b)
	return a*power + b
}

func getPower(num int) int {
	power := 1
	for i := num; i > 0; i /= 10 {
		power *= 10
	}
	return power
}
