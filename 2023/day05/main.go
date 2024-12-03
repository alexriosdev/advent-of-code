package main

import (
	"advent-of-code/utils"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("2023 Day 05 Solution")
	fmt.Printf("Part 1: %v\n", part1(input))
	fmt.Printf("Part 2: %v\n", part2(input))
}

func part1(input []byte) int {
	split := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	seeds := []int{}
	for _, s := range strings.Fields(strings.Split(split[0], ":")[1]) {
		seeds = append(seeds, utils.StrToInt(s))
	}

	conversionMaps := [][]rangeMap{}
	for i := 1; i < len(split); i++ {
		conversionMap := []rangeMap{}
		for _, s := range strings.Split(strings.Split(split[i], ":\n")[1], "\n") {
			numbers := strings.Fields(s)
			conversionMap = append(conversionMap, rangeMap{
				dest:   utils.StrToInt(numbers[0]),
				source: utils.StrToInt(numbers[1]),
				len:    utils.StrToInt(numbers[2]),
			})
		}
		conversionMaps = append(conversionMaps, conversionMap)
	}

	result := math.MaxInt
	for _, seed := range seeds {
		result = min(result, convertSeed(seed, conversionMaps))
	}
	return result
}

func part2(input []byte) int {
	split := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	seeds := []int{}
	for _, s := range strings.Fields(strings.Split(split[0], ":")[1]) {
		seeds = append(seeds, utils.StrToInt(s))
	}

	conversionMaps := [][]rangeMap{}
	for i := 1; i < len(split); i++ {
		conversionMap := []rangeMap{}
		for _, s := range strings.Split(strings.Split(split[i], ":\n")[1], "\n") {
			numbers := strings.Fields(s)
			conversionMap = append(conversionMap, rangeMap{
				dest:   utils.StrToInt(numbers[0]),
				source: utils.StrToInt(numbers[1]),
				len:    utils.StrToInt(numbers[2]),
			})
		}
		conversionMaps = append(conversionMaps, conversionMap)
	}

	// TODO: This approach is suboptimal and will timeout in tests.
	// It takes about 8mins to finish. Must find better solution.
	result := math.MaxInt
	for i := range seeds {
		if i%2 == 0 {
			for seed := seeds[i]; seed < seeds[i]+seeds[i+1]; seed++ {
				result = min(result, convertSeed(seed, conversionMaps))
			}
		}
	}
	return result
}

type rangeMap struct {
	dest, source, len int
}

func convertSeed(seed int, conversionMaps [][]rangeMap) int {
	for _, conversionMap := range conversionMaps {
		for _, r := range conversionMap {
			if r.source <= seed && seed < r.source+r.len {
				seed = r.dest + (seed - r.source)
				break
			}
		}
	}
	return seed
}
