package main

import (
	"advent-of-code/utils"
	"testing"
)

func TestPart1(t *testing.T) {
	lines1, _ := utils.ReadLines("input_test1.txt")
	lines2, _ := utils.ReadLines("input.txt")
	tests := []struct {
		input    []string
		expected int
	}{
		{lines1, 142},
		{lines2, 54644},
	}
	for _, test := range tests {
		result := part1(test.input)
		if result != test.expected {
			t.Errorf("Result %v not equal to expected %v", result, test.expected)
		}
	}
}

func TestPart2(t *testing.T) {
	lines1, _ := utils.ReadLines("input_test2.txt")
	lines2, _ := utils.ReadLines("input.txt")
	tests := []struct {
		input    []string
		expected int
	}{
		{lines1, 281},
		{lines2, 53348},
	}
	for _, test := range tests {
		result := part2(test.input)
		if result != test.expected {
			t.Errorf("Result %v not equal to expected %v", result, test.expected)
		}
	}
}
