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
		{input1, 405},
		{input2, 30487},
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
		{input1, 400},
		{input2, 31954},
	}
	for _, test := range tests {
		result := part2(test.input)
		if result != test.expected {
			t.Errorf("Result %v not equal to expected %v", result, test.expected)
		}
	}
}

func BenchmarkPart1(b *testing.B) {
	input, _ := os.ReadFile("input_test.txt")
	for n := 0; n < b.N; n++ {
		part1(input)
	}
}

func BenchmarkPart2(b *testing.B) {
	input, _ := os.ReadFile("input_test.txt")
	for n := 0; n < b.N; n++ {
		part1(input)
	}
}
