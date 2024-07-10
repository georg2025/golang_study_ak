package main

import (
	"testing"
)

func TestFactorialRecursive(t *testing.T) {
	result := factorialRecursive(5)
	expected := 120

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}

	result = factorialRecursive(0)
	expected = 1

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestFactorialIterative(t *testing.T) {
	result := factorialIterative(5)
	expected := 120

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}

	result = factorialIterative(0)
	expected = 1

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestCompareWhichFactorialIsFaster(t *testing.T) {
	result := compareWhichFactorialIsFaster(5)
	expected := "Iterative func is faster"

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}
