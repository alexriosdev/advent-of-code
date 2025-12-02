package main

import (
	"advent-of-code/utils"
	"fmt"
	"math"
	"strings"
)

func main() {
	line, _ := utils.ReadLine("2025/day02/input.txt")
	fmt.Println("2025 Day 02 Solution")
	fmt.Printf("Part 1: %v\n", part1(line))
}

func part1(line string) int {
	sum := 0
	for _, s := range strings.Split(line, ",") {
		split := strings.Split(s, "-")
		start, end := utils.StrToInt(split[0]), utils.StrToInt(split[1])
		for num := start; num <= end; num++ {
			n := countDigits(num)
			if n%2 != 0 {
				continue
			}
			a, b := splitNumber(num, n)
			if a-b == 0 {
				sum += num
			}
		}
	}
	return sum
}

func countDigits(num int) int {
	if num == 0 {
		return 1
	}
	num = utils.Abs(num)
	return int(math.Floor(math.Log10(float64(num)))) + 1
}

func splitNumber(num, n int) (int, int) {
	divisor := int(math.Pow(10, float64(n/2)))
	return num % divisor, num / divisor
}
