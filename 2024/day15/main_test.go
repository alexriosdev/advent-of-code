package main

import (
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	input1, _ := os.ReadFile("input_test1.txt")
	input2, _ := os.ReadFile("input_test2.txt")
	input3, _ := os.ReadFile("input.txt")
	tests := []struct {
		input    []byte
		expected int
	}{
		{input1, 2028},
		{input2, 10092},
		{input3, 1516281},
	}
	for _, test := range tests {
		result := part1(test.input)
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
