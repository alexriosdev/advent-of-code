package main

import (
	"advent-of-code/utils"
	"fmt"
)

func main() {
	lines1, _ := utils.ReadLines("2025/day01/input.txt")
	lines2, _ := utils.ReadLines("2025/day01/input.txt")
	fmt.Println("2025 Day 01 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines1))
	fmt.Printf("Part 2: %v\n", part2(lines2))
}

func part1(lines []string) int {
	dial := &Dial{}
	dial.Init(99)
	dial.Rotate('R', 50)
	for _, line := range lines {
		dir, dist := rune(line[0]), utils.StrToInt(line[1:])
		dial.Rotate(dir, dist)
		dial.CountZeroes()
	}
	return dial.Zeroes
}

func part2(lines []string) int {
	dial := &Dial{}
	dial.Init(99)
	dial.Rotate('R', 50)
	for _, line := range lines {
		dir, dist := rune(line[0]), utils.StrToInt(line[1:])
		dial.RotateWithCount(dir, dist)
	}
	return dial.Zeroes
}

type Node struct {
	Val        int
	Prev, Next *Node
}

type Dial struct {
	Head        *Node
	Len, Zeroes int
}

func (d *Dial) Init(size int) *Dial {
	for i := size; i >= 0; i-- {
		d.Insert(i)
	}
	return d
}

func (d *Dial) Insert(val int) {
	curr := &Node{Val: val}
	d.Len++

	if d.Head == nil {
		curr.Next = curr
		curr.Prev = curr
		d.Head = curr
		return
	}

	prev := d.Head.Prev
	curr.Next = d.Head
	curr.Prev = prev

	d.Head.Prev = curr
	prev.Next = curr

	d.Head = curr
}

func (d *Dial) Rotate(dir rune, dist int) {
	for i := 0; i < dist; i++ {
		if dir == 'R' {
			d.Head = d.Head.Next
		} else {
			d.Head = d.Head.Prev
		}
	}
}

func (d *Dial) RotateWithCount(dir rune, dist int) {
	for i := 0; i < dist; i++ {
		if dir == 'R' {
			d.Head = d.Head.Next
		} else {
			d.Head = d.Head.Prev
		}
		d.CountZeroes()
	}
}

func (d *Dial) CountZeroes() {
	if d.Head.Val == 0 {
		d.Zeroes++
	}
}
