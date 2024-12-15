package main

import (
	"advent-of-code/utils"
	"fmt"
	"strings"
)

func main() {
	lines, _ := utils.ReadLines("input.txt")
	fmt.Println("2024 Day 14 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines, 103, 101))
}

func part1(lines []string, m, n int) int {
	robots := getRobots(lines)
	quadrants := [4]int{}
	for i := range robots {
		robots[i].p.x = (robots[i].p.x + robots[i].v.x*100) % n
		robots[i].p.y = (robots[i].p.y + robots[i].v.y*100) % m
		if robots[i].p.x < 0 {
			abs := utils.Abs(robots[i].p.x)
			robots[i].p.x = n - abs
		}
		if robots[i].p.y < 0 {
			abs := utils.Abs(robots[i].p.y)
			robots[i].p.y = m - abs
		}
		if robots[i].p.x < n/2 && robots[i].p.y < m/2 {
			quadrants[0]++
		}
		if robots[i].p.x > n/2 && robots[i].p.y < m/2 {
			quadrants[1]++
		}
		if robots[i].p.x < n/2 && robots[i].p.y > m/2 {
			quadrants[2]++
		}
		if robots[i].p.x > n/2 && robots[i].p.y > m/2 {
			quadrants[3]++
		}
	}
	result := 1
	for _, quadrant := range quadrants {
		result *= quadrant
	}
	return result
}

type coordinate struct {
	x, y int
}

type robot struct {
	p, v coordinate
}

func getRobots(lines []string) []robot {
	replacer := strings.NewReplacer("p=", "", "v=", "", ",", " ")
	robots := []robot{}
	for _, line := range lines {
		split := strings.Fields(replacer.Replace(line))
		robot := robot{
			coordinate{utils.StrToInt(split[0]), utils.StrToInt(split[1])},
			coordinate{utils.StrToInt(split[2]), utils.StrToInt(split[3])},
		}
		robots = append(robots, robot)
	}
	return robots
}
