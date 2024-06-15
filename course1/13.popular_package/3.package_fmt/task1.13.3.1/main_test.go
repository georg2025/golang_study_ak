package main

import (
	"testing"
)

func TestGenerateMathString(t *testing.T) {
	result := generateMathString([]int{22, 2, 5, 2}, "+")
	expected := "22 + 2 + 5 + 2 = 31"
	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
	result = generateMathString([]int{22, 2, 5, 2}, "-")
	expected = "22 - 2 - 5 - 2 = 13"
	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
	result = generateMathString([]int{22, 2, 5, 2}, "*")
	expected = "22 * 2 * 5 * 2 = 440"
	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
	result = generateMathString([]int{40, 2, 5, 2}, "/")
	expected = "40 / 2 / 5 / 2 = 2"
	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
	result = generateMathString([]int{}, "-")
	expected = ""
	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
	result = generateMathString([]int{5}, "-")
	expected = "5 = 5"
	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
	result = generateMathString([]int{4, 5, 6}, "^")
	expected = ""
	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}
