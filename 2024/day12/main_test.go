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
		{lines4, 1473620},
	}
	for _, test := range tests {
		result := part1(test.input)
		if result != test.expected {
			t.Errorf("Result %v not equal to expected %v", result, test.expected)
		}
	}
}

func TestPart2(t *testing.T) {
	lines1, _ := utils.ReadLines("input_test1.txt")
	lines2, _ := utils.ReadLines("input_test2.txt")
	lines3, _ := utils.ReadLines("input_test4.txt")
	lines4, _ := utils.ReadLines("input_test5.txt")
	lines5, _ := utils.ReadLines("input.txt")
	tests := []struct {
		input    []string
		expected int
	}{
		{lines1, 80},
		{lines2, 436},
		{lines3, 236},
		{lines4, 368},
		{lines5, 902620},
	}
	for _, test := range tests {
		result := part2(test.input)
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

func BenchmarkPart2(b *testing.B) {
	lines, _ := utils.ReadLines("input.txt")
	for n := 0; n < b.N; n++ {
		part2(lines)
	}
}
