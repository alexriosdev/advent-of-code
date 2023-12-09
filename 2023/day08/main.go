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
	fmt.Printf("Part 2: %v\n", part2(lines))
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

func part2(lines []string) int {
	inst := []rune(lines[0])
	replacer := strings.NewReplacer("=", "", "(", "", ",", "", ")", "")
	startNodes := []string{}
	nodeMap := map[string]node{}
	for i := 1; i < len(lines); i++ {
		split := strings.Fields(replacer.Replace(lines[i]))
		nodeMap[split[0]] = node{left: split[1], right: split[2]}
		if split[0][2] == 'A' {
			startNodes = append(startNodes, split[0])
		}
	}
	steps := []int{}
	for _, curr := range startNodes {
		count := 0
		for curr[2] != 'Z' {
			c := inst[count%len(inst)]
			node := nodeMap[curr]
			if c == 'L' {
				curr = node.left
			} else {
				curr = node.right
			}
			count++
		}
		steps = append(steps, count)
	}
	return lcm(steps, 0)
}

func lcm(nums []int, i int) int {
	if i == len(nums)-1 {
		return nums[i]
	}
	a, b := nums[i], lcm(nums, i+1)
	return (a * b) / gcm(a, b)
}

func gcm(a, b int) int {
	if b == 0 {
		return a
	}
	return gcm(b, a%b)
}
