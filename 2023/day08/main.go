package main

import (
	"advent-of-code/utils"
	"fmt"
	"strings"
)

type node struct {
	left, right string
}

func main() {
	lines, _ := utils.ReadLines("input.txt")
	fmt.Println("2023 Day 08 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines))
}

func part1(lines []string) int {
	inst := []rune(lines[0])
	replacer := strings.NewReplacer("=", "", "(", "", ",", "", ")", "")
	nodeMap := map[string]node{}
	for i := 1; i < len(lines); i++ {
		split := strings.Fields(replacer.Replace(lines[i]))
		nodeMap[split[0]] = node{left: split[1], right: split[2]}
	}
	curr := "AAA"
	count := 0
	for curr != "ZZZ" {
		c := inst[count%len(inst)]
		node := nodeMap[curr]
		if c == 'L' {
			curr = node.left
		} else {
			curr = node.right
		}
		count++
	}
	return count
}
