package main

import (
	"advent-of-code/utils"
	"fmt"
	"strconv"
	"strings"
)

var UP 		= coordinate{-1, 0}
var RIGHT 	= coordinate{0, 1}
var DOWN 	= coordinate{1, 0}
var LEFT 	= coordinate{0, -1}
var ORIGIN	= coordinate{0, 0}

type coordinate struct {
	y, x int
}

func main() {
	lines, _ := utils.ReadLines("input.txt")
	fmt.Println("2023 Day 18 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines))
	fmt.Printf("Part 2: %v\n", part2(lines))
}

func part1(lines []string) int {
	points := []coordinate{ORIGIN}
	totalDist := 0
	for _, line := range lines {
		split := strings.Fields(line)
		dir, dist := split[0], utils.StrToInt(split[1])
		point := points[len(points)-1]
		switch dir {
		case "U":
			point = coordinate{point.y + UP.y*dist, point.x + UP.x*dist}
		case "R":
			point = coordinate{point.y + RIGHT.y*dist, point.x + RIGHT.x*dist}
		case "D":
			point = coordinate{point.y + DOWN.y*dist, point.x + DOWN.x*dist}
		case "L":
			point = coordinate{point.y + LEFT.y*dist, point.x + LEFT.x*dist}
		}
		points = append(points, point)
		totalDist += dist
	}
	area := applyShoelaceFormula(points)
	return applyPicksTheorem(totalDist, area)
}

func part2(lines []string) int {
	points := []coordinate{ORIGIN}
	totalDist := 0
	replacer := strings.NewReplacer("(", "", "#", "", ")", "")
	for _, line := range lines {
		color := replacer.Replace(strings.Fields(line)[2])
		dir, dist := color[5], hexToInt(color[:5])
		point := points[len(points)-1]
		switch dir {
		case '0':
			point = coordinate{point.y + UP.y*dist, point.x + UP.x*dist}
		case '1':
			point = coordinate{point.y + RIGHT.y*dist, point.x + RIGHT.x*dist}
		case '2':
			point = coordinate{point.y + DOWN.y*dist, point.x + DOWN.x*dist}
		case '3':
			point = coordinate{point.y + LEFT.y*dist, point.x + LEFT.x*dist}
		}
		points = append(points, point)
		totalDist += dist
	}
	area := applyShoelaceFormula(points)
	return applyPicksTheorem(totalDist, area)
}

func applyShoelaceFormula(points []coordinate) int {
	sum := 0
	j := len(points) - 1
	for i := 0; i < len(points); i++ {
		sum += (points[i].x * points[j].y) - (points[j].x * points[i].y)
		j = i
	}
	return utils.Abs(sum) / 2
}

func applyPicksTheorem(totalDist, area int) int {
	return (totalDist / 2) + area + 1
}

func hexToInt(s string) int {
	num, _ := strconv.ParseInt(s, 16, 64)
	return int(num)
}
