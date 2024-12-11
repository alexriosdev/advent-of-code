package main

import (
	"advent-of-code/utils"
	"testing"
)

func TestPart1(t *testing.T) {
	lines1, _ := utils.ReadLines("input_test.txt")
	lines2, _ := utils.ReadLines("input.txt")
	tests := []struct {
		input    []string
		n        int
		expected int
	}{
		{lines1, 6, 22},
		{lines1, 25, 55312},
		{lines2, 25, 189092},
	}
	for _, test := range tests {
		result := part1(test.input, test.n)
		if result != test.expected {
			t.Errorf("Result %v not equal to expected %v", result, test.expected)
		}
	}
}

func TestPart2(t *testing.T) {
	lines2, _ := utils.ReadLines("input.txt")
	tests := []struct {
		input    []string
		n        int
		expected int
	}{
		{lines2, 75, 224869647102559},
	}
	for _, test := range tests {
		result := part1(test.input, test.n)
		if result != test.expected {
			t.Errorf("Result %v not equal to expected %v", result, test.expected)
		}
	}
}

func BenchmarkPart1(b *testing.B) {
	lines, _ := utils.ReadLines("input.txt")
	for n := 0; n < b.N; n++ {
		part1(lines, 25)
	}
}

func BenchmarkPart2(b *testing.B) {
	lines, _ := utils.ReadLines("input.txt")
	for n := 0; n < b.N; n++ {
		part1(lines, 75)
	}
}
