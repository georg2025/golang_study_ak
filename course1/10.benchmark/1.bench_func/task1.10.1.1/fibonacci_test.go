package main

import "testing"

type testData struct {
	input    int
	expected int
}

func TestFibonacci(t *testing.T) {
	testCases := []testData{
		{input: 0, expected: 0},
		{input: 5, expected: 5},
		{input: 6, expected: 8},
		{input: 7, expected: 13},
		{input: 8, expected: 21},
		{input: 9, expected: 34},
		{input: 10, expected: 55},
	}
	for _, tc := range testCases {
		result := Fibonacci(tc.input)
		if result != tc.expected {
			t.Errorf("For input %d result is %d, expected %d", tc.input, result, tc.expected)
		}
	}
}

func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < 1000; i++ {
		Fibonacci(i)
	}
}
