package main

import (
	"advent-of-code/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	lines, _ := utils.ReadLines("input.txt")
	fmt.Println("2023 Day 05 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines))
}

func part1(lines []string) int {
	seeds := []int{}
	for _, s := range strings.Fields(strings.Split(lines[0], ":")[1]) {
		seed, _ := strconv.Atoi(s)
		seeds = append(seeds, seed)
	}

	indexes := make([]index, 7)
	for i, line := range lines {
		switch {
		case line == "seed-to-soil map:":
			indexes[0].start = i + 1
		case line == "soil-to-fertilizer map:":
			indexes[1].start = i + 1
		case line == "fertilizer-to-water map:":
			indexes[2].start = i + 1
		case line == "water-to-light map:":
			indexes[3].start = i + 1
		case line == "light-to-temperature map:":
			indexes[4].start = i + 1
		case line == "temperature-to-humidity map:":
			indexes[5].start = i + 1
		case line == "humidity-to-location map:":
			indexes[6].start = i + 1
		}
	}

	for i := 0; i < len(indexes)-1; i++ {
		indexes[i].end = indexes[i+1].start - 1
	}
	indexes[6].end = len(lines)

	conversionMaps := make([][]rangeMap, 7)
	for i, index := range indexes {
		conversionMap := []rangeMap{}
		for j := index.start; j < index.end; j++ {
			numbers := strings.Fields(lines[j])
			dest, _ := strconv.Atoi(numbers[0])
			source, _ := strconv.Atoi(numbers[1])
			len, _ := strconv.Atoi(numbers[2])
			conversionMap = append(conversionMap, rangeMap{
				dest:   dest,
				source: source,
				len:    len,
			})
		}
		conversionMaps[i] = conversionMap
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

type index struct {
	start, end int
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
