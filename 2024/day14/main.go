package main

import (
	"advent-of-code/utils"
	"fmt"
	"strings"
)

func main() {
	lines, _ := utils.ReadLines("2024/day14/input.txt")
	fmt.Println("2024 Day 14 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines, 103, 101))
	fmt.Printf("Part 2: %v\n", part2(lines, 103, 101, true))
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

func part2(lines []string, m, n int, useDisplay bool) int {
	robots := getRobots(lines)
	grid := buildGrid(m, n)
	for sec := 1; sec <= m*n; sec++ {
		for i := range robots {
			robots[i].p.x = (robots[i].p.x + robots[i].v.x) % n
			robots[i].p.y = (robots[i].p.y + robots[i].v.y) % m
			if robots[i].p.x < 0 {
				abs := utils.Abs(robots[i].p.x)
				robots[i].p.x = n - abs
			}
			if robots[i].p.y < 0 {
				abs := utils.Abs(robots[i].p.y)
				robots[i].p.y = m - abs
			}
			grid[robots[i].p.y][robots[i].p.x] = '#'
		}
		if useDisplay {
			displayGrid(&grid, sec)
		}
		if findOutline(&grid) {
			return sec
		}
		clearGrid(&grid)
	}
	return -1
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

func buildGrid(m, n int) [][]rune {
	grid := make([][]rune, m)
	for i := range grid {
		grid[i] = make([]rune, n)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}
	return grid
}

func displayGrid(grid *[][]rune, second int) {
	fmt.Println(second)
	for _, row := range *grid {
		fmt.Println(string(row))
	}
	fmt.Println()
	// milliseconds := 50
	// time.Sleep(time.Duration(milliseconds) * time.Millisecond)
}

func clearGrid(grid *[][]rune) {
	for i := range *grid {
		for j := range (*grid)[i] {
			(*grid)[i][j] = '.'
		}
	}
}

func findOutline(grid *[][]rune) bool {
	outlineSize := 30
	top, bottom := 0, 0
	left, right := 0, 0
	for i := 0; i < len(*grid)-outlineSize; i++ {
		if top >= outlineSize {
			break
		}
		for j := 0; j < len((*grid)[0])-outlineSize; j++ {
			if (*grid)[i][j] != '#' {
				top = 0
				continue
			} else {
				top++
			}
			if top >= outlineSize {
				for k := i; k < len(*grid)-outlineSize; k++ {
					if bottom >= outlineSize {
						break
					}
					for l := 0; l < len((*grid)[0])-outlineSize; l++ {
						if (*grid)[k][l] != '#' {
							bottom = 0
							continue
						} else {
							bottom++
						}
						if bottom >= outlineSize {
							break
						}
					}
				}
				break
			}
		}
	}
	for i := 0; i < len((*grid)[0])-outlineSize; i++ {
		if left >= outlineSize {
			break
		}
		for j := 0; j < len(*grid)-outlineSize; j++ {
			if (*grid)[j][i] != '#' {
				left = 0
				continue
			} else {
				left++
			}
			if left >= outlineSize {
				for k := i; k < len((*grid)[0])-outlineSize; k++ {
					if right >= outlineSize {
						break
					}
					for l := 0; l < len(*grid)-outlineSize; l++ {
						if (*grid)[l][k] != '#' {
							right = 0
							continue
						} else {
							right++
						}
						if right >= outlineSize {
							break
						}
					}
				}
				break
			}
		}
	}
	return top == outlineSize && top == bottom && left == outlineSize && left == right
}
