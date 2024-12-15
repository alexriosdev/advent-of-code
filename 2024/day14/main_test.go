package main

import (
	"advent-of-code/utils"
	"testing"
)

func TestPart1(t *testing.T) {
	lines1, _ := utils.ReadLines("input_test.txt")
	lines2, _ := utils.ReadLines("input.txt")
	tests := []struct {
		input          []string
		m, n, expected int
	}{
		{lines1, 11, 7, 12},
		{lines2, 103, 101, 232589280},
	}
	for _, test := range tests {
		result := part1(test.input, test.m, test.n)
		if result != test.expected {
			t.Errorf("Result %v not equal to expected %v", result, test.expected)
		}
	}
}

func TestPart2(t *testing.T) {
	lines2, _ := utils.ReadLines("input.txt")
	tests := []struct {
		input          []string
		useDisplay     bool
		m, n, expected int
	}{
		{lines2, false, 103, 101, 7569},
	}
	for _, test := range tests {
		result := part2(test.input, test.m, test.n, test.useDisplay)
		if result != test.expected {
			t.Errorf("Result %v not equal to expected %v", result, test.expected)
		}
	}
}

func BenchmarkPart1(b *testing.B) {
	lines, _ := utils.ReadLines("input.txt")
	for n := 0; n < b.N; n++ {
		part1(lines, 103, 101)
	}
}

func BenchmarkPart2(b *testing.B) {
	lines, _ := utils.ReadLines("input.txt")
	for n := 0; n < b.N; n++ {
		part2(lines, 103, 101, false)
	}
}
