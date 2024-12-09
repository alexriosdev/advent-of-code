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
		expected int
	}{
		{lines1, 1928},
		{lines2, 6384282079460},
	}
	for _, test := range tests {
		result := part1(test.input)
		if result != test.expected {
			t.Errorf("Result %v not equal to expected %v", result, test.expected)
		}
	}
}

func BenchmarkPart1(b *testing.B) {
	line, _ := utils.ReadLines("input.txt")
	for n := 0; n < b.N; n++ {
		part1(line)
	}
}
