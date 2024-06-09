package main

import "testing"

type testData struct {
	slice  []int
	result int
}

func TestFindSingleNumber(t *testing.T) {
	testCases := []testData{
		{slice: []int{}, result: 0},
		{slice: []int{4}, result: 4},
		{slice: []int{2, 5, 3, 1, 3, 5, 2}, result: 1},
		{slice: []int{-8, -8, 4, 4, 5, 5, -6}, result: -6},
		{slice: []int{0, 0, 0, 0, 7}, result: 7},
	}
	for _, tc := range testCases {
		result := findSingleNumber(tc.slice)
		if result != tc.result {
			t.Errorf("Expected %d, got %d", tc.result, result)
		}
	}
}

func TestBitwiseXOR(t *testing.T) {
	result := bitwiseXOR(2, 2)
	expected := 0
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
