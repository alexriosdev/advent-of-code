package main

import (
	"advent-of-code/utils"
	"testing"
)

func TestPart1(t *testing.T) {
	lines1, _ := utils.ReadLines("input_test1.txt")
	lines2, _ := utils.ReadLines("input_test2.txt")
	lines3, _ := utils.ReadLines("input_test3.txt")
	lines4, _ := utils.ReadLines("input.txt")
	tests := []struct {
		input    []string
		expected int
	}{
		{lines1, 140},
		{lines2, 772},
		{lines3, 1930},
		{lines4, 0},
	}
	for _, test := range tests {
		result := part1(test.input)
		if result != test.expected {
			t.Errorf("Result %v not equal to expected %v", result, test.expected)
		}
	}
}

func BenchmarkPart1(b *testing.B) {
	lines, _ := utils.ReadLines("input.txt")
	for n := 0; n < b.N; n++ {
		part1(lines)
	}
}
