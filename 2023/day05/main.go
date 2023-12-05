package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("2023 Day 05 Solution")
	fmt.Printf("Part 1: %v\n", part1(input))
}

func part1(input []byte) int {
	split := strings.Split(string(input), "\n\n")
	seeds := []int{}
	for _, s := range strings.Fields(strings.Split(split[0], ":")[1]) {
		seeds = append(seeds, strToInt(s))
	}

	conversionMaps := [][]rangeMap{}
	for i := 1; i < len(split); i++ {
		conversionMap := []rangeMap{}
		for _, s := range strings.Split(strings.Split(split[i], ":\n")[1], "\n") {
			numbers := strings.Fields(s)
			if len(numbers) == 0 {
				continue
			}
			conversionMap = append(conversionMap, rangeMap{
				dest:   strToInt(numbers[0]),
				source: strToInt(numbers[1]),
				len:    strToInt(numbers[2]),
			})
		}
		conversionMaps = append(conversionMaps, conversionMap)
	}

	for i := range seeds {
		for _, conversionMap := range conversionMaps {
			for _, r := range conversionMap {
				if r.source <= seeds[i] && seeds[i] < r.source+r.len {
					seeds[i] = r.dest + (seeds[i] - r.source)
					break
				}
			}
		}
	}

	result := math.MaxInt
	for _, seed := range seeds {
		result = getMin(result, seed)
	}
	return result
}

type rangeMap struct {
	dest, source, len int
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func strToInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}
