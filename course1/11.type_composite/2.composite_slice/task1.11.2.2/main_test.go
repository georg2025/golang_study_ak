package main

import "testing"

func TestMaxDifference(t *testing.T) {
	result := MaxDifference([]int{1, 5, 2, 55, 1, 99, 6, 44})
	expected := 98
	if result != expected {
		t.Errorf("Got %d, expected %d", result, expected)
	}
	result = MaxDifference([]int{})
	expected = 0
	if result != expected {
		t.Errorf("Got %d, expected %d", result, expected)
	}
	result = MaxDifference([]int{1, 5, 2, 55, 1, 6, -44})
	expected = 99
	if result != expected {
		t.Errorf("Got %d, expected %d", result, expected)
	}
}
