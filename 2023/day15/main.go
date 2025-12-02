package main

import (
	"advent-of-code/utils"
	"fmt"
	"strings"
	"unicode"

	"github.com/emirpasic/gods/maps/linkedhashmap"
)

func main() {
	lines, _ := utils.ReadLines("2023/day15/input.txt")
	fmt.Println("2023 Day 15 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines))
	fmt.Printf("Part 2: %v\n", part2(lines))
}

func part1(lines []string) int {
	result := 0
	for _, line := range lines {
		val := 0
		for _, c := range line {
			if c == ',' {
				result += val
				val = 0
				continue
			}
			val = applyHash(val, c)
		}
		result += val
	}
	return result
}

func part2(lines []string) int {
	boxes := make([]linkedhashmap.Map, 256)
	for i := range boxes {
		boxes[i] = *linkedhashmap.New()
	}
	for _, line := range lines {
		sb := strings.Builder{}
		val := 0
		for _, c := range line {
			switch {
			case unicode.IsLetter(c):
				sb.WriteRune(c)
				val = applyHash(val, c)
			case unicode.IsNumber(c):
				label := sb.String()
				boxes[val].Put(label, int(c-'0'))
				val = 0
				sb.Reset()
			case c == '-':
				label := sb.String()
				boxes[val].Remove(label)
				sb.Reset()
				val = 0
			}
		}
	}
	result := 0
	for i, box := range boxes {
		if box.Empty() {
			continue
		}
		slot := 1
		iter := box.Iterator()
		for iter.Next() {
			result += applyFocusingPower(i, slot, iter.Value().(int))
			slot++
		}
	}
	return result
}

func applyHash(val int, c rune) int {
	return ((val + int(c)) * 17) % 256
}

func applyFocusingPower(box, slot, lens int) int {
	return (1 + box) * slot * lens
}
