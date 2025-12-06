package main

import (
	"advent-of-code/utils"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	input, _ := os.ReadFile("2025/day05/input.txt")
	fmt.Println("2025 Day 05 Solution")
	fmt.Printf("Part 1: %v\n", part1(input))
	fmt.Printf("Part 2: %v\n", part2(input))
}

func part1(input []byte) int {
	sections := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	ingredients := []int{}
	for _, s := range strings.Split(sections[1], "\n") {
		ingredients = append(ingredients, utils.StrToInt(s))
	}
	count := 0
	for _, ingredient := range ingredients {
		for _, s := range strings.Split(sections[0], "\n") {
			split := strings.Split(s, "-")
			start, end := utils.StrToInt(split[0]), utils.StrToInt(split[1])
			if start <= ingredient && ingredient <= end {
				count++
				break
			}
		}
	}
	return count
}

func part2(input []byte) int {
	sections := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	intervals := Intervals{}
	for _, s := range strings.Split(sections[0], "\n") {
		split := strings.Split(s, "-")
		start, end := utils.StrToInt(split[0]), utils.StrToInt(split[1])
		intervals.Add(start, end)
	}
	intervals.SortAsc()
	merged := intervals.GetMerged()
	count := 0
	for _, interval := range merged {
		count += interval.End - interval.Start + 1
	}
	return count
}

type Interval struct {
	Start, End int
}

type Intervals []Interval

func (intervals *Intervals) Add(start, end int) {
	*intervals = append(*intervals, Interval{start, end})
}

func (intervals *Intervals) SortAsc() {
	sort.Slice(*intervals, func(i, j int) bool {
		return (*intervals)[i].Start < (*intervals)[j].Start
	})
}

func (intervals *Intervals) GetMerged() []Interval {
	merged := []Interval{}
	for _, interval := range *intervals {
		n := len(merged)
		if n == 0 || merged[n-1].End < interval.Start {
			merged = append(merged, interval)
		} else {
			start, end := merged[n-1].Start, max(merged[n-1].End, interval.End)
			merged[n-1] = Interval{start, end}
		}
	}
	return merged
}
