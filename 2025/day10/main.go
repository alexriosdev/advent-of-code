package main

import (
	"advent-of-code/utils"
	"fmt"
	"math"
	"slices"
	"strings"
)

func main() {
	lines, _ := utils.ReadLines("2025/day10/input.txt")
	fmt.Println("2025 Day 10 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines))
	fmt.Printf("Part 2: %v\n", part2(lines))
}

func part1(lines []string) int {
	lineReplacer := strings.NewReplacer("] (", "]x(", ") {", ")x{")
	machines := Machines{}
	for _, line := range lines {
		split := strings.Split(lineReplacer.Replace(line), "x")
		machine := NewMachine(split)
		machines.Add(machine)
	}
	sum := 0
	for _, machine := range machines {
		machine.PerformCombinations()
		sum += machine.MinCombinations
	}
	return sum
}

func part2(lines []string) int {
	return len(lines)
}

type Machine struct {
	Lights          Lights
	Buttons         Buttons
	MinCombinations int
}

func NewMachine(s []string) Machine {
	return Machine{
		Lights:          NewLights(s[0]),
		Buttons:         NewButtons(s[1]),
		MinCombinations: math.MaxInt32,
	}
}

func (m *Machine) PerformCombinations() {
	combinations := []int{}
	m.Backtrack(0, combinations)
}

func (m *Machine) Backtrack(start int, combinations []int) {
	if m.Lights.MatchConfig() {
		m.MinCombinations = min(m.MinCombinations, len(combinations))
		return
	}
	for i := start; i < len(m.Buttons); i++ {
		combinations = append(combinations, i)
		m.Buttons[i].Press(&m.Lights)
		m.Backtrack(i+1, combinations)
		m.Buttons[i].UndoPress(&m.Lights)
		combinations = combinations[:len(combinations)-1]
	}
}

type Machines []Machine

func (m *Machines) Add(machine Machine) {
	*m = append(*m, machine)
}

type Lights struct {
	Mutable []bool
	Final   []bool
}

func NewLights(s string) Lights {
	s = strings.NewReplacer("[", "", "]", "").Replace(s)
	n := len(s)
	values := make([]bool, n)
	for i, c := range s {
		if c == '#' {
			values[i] = true
		}
	}
	return Lights{make([]bool, n), values}
}

func (l *Lights) Toggle(i int) {
	(*l).Mutable[i] = !(*l).Mutable[i]
}

func (l *Lights) MatchConfig() bool {
	return slices.Equal(l.Mutable, l.Final)
}

type Button []int

func NewButton(split []string) Button {
	button := make(Button, len(split))
	for i, s := range split {
		button[i] = utils.StrToInt(s)
	}
	return button
}

func (b *Button) Press(l *Lights) {
	for _, val := range *b {
		l.Toggle(val)
	}
}

func (b *Button) UndoPress(l *Lights) {
	b.Press(l)
}

type Buttons []Button

func NewButtons(s string) Buttons {
	fields := strings.Fields(s)
	buttons := make([]Button, len(fields))
	for i, field := range fields {
		split := strings.Split(strings.NewReplacer("(", "", ")", "").Replace(field), ",")
		buttons[i] = NewButton(split)
	}
	return buttons
}
