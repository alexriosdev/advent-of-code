package main

import (
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	input1, _ := os.ReadFile("input_test.txt")
	input2, _ := os.ReadFile("input.txt")
	tests := []struct {
		input    []byte
		expected int
	}{
		{input1, 143},
		{input2, 5108},
	}
	for _, test := range tests {
		result := part1(test.input)
		if result != test.expected {
			t.Errorf("Result %v not equal to expected %v", result, test.expected)
		}
	}
}

func TestPart2(t *testing.T) {
	input1, _ := os.ReadFile("input_test.txt")
	input2, _ := os.ReadFile("input.txt")
	tests := []struct {
		input    []byte
		expected int
	}{
		{input1, 123},
		{input2, 7380},
	}
	for _, test := range tests {
		result := part2(test.input)
		if result != test.expected {
			t.Errorf("Result %v not equal to expected %v", result, test.expected)
		}
	}
}

func BenchmarkPart1(b *testing.B) {
	input, _ := os.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		part1(input)
	}
}

func BenchmarkPart2(b *testing.B) {
	input, _ := os.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		part2(input)
	}
}
