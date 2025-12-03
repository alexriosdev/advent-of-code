package main

import (
	"advent-of-code/utils"
	"testing"
)

func TestPart1(t *testing.T) {
	line1, _ := utils.ReadLine("input_test.txt")
	line2, _ := utils.ReadLine("input.txt")
	tests := []struct {
		input    string
		expected int
	}{
		{line1, 1227775554},
		{line2, 15873079081},
	}
	for _, test := range tests {
		result := part1(test.input)
		if result != test.expected {
			t.Errorf("Result %v not equal to expected %v", result, test.expected)
		}
	}
}

func TestPart2(t *testing.T) {
	line1, _ := utils.ReadLine("input_test.txt")
	line2, _ := utils.ReadLine("input.txt")
	tests := []struct {
		input    string
		expected int
	}{
		{line1, 4174379265},
		{line2, 22617871034},
	}
	for _, test := range tests {
		result := part2(test.input)
		if result != test.expected {
			t.Errorf("Result %v not equal to expected %v", result, test.expected)
		}
	}
}

func BenchmarkPart1(b *testing.B) {
	line, _ := utils.ReadLine("input.txt")
	for n := 0; n < b.N; n++ {
		part1(line)
	}
}

func BenchmarkPart2(b *testing.B) {
	line, _ := utils.ReadLine("input.txt")
	for n := 0; n < b.N; n++ {
		part2(line)
	}
}
