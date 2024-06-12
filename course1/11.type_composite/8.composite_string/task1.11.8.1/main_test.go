package main

import (
	"testing"
)

func TestCountBytes(t *testing.T) {
	result := countBytes("Привет, мир!")
	expected := 21
	if result != expected {
		t.Errorf("Got %d, expected %d", result, expected)
	}
}

func TestCountSymbols(t *testing.T) {
	result := countSymbols("Привет, мир!")
	expected := 12
	if result != expected {
		t.Errorf("Got %d, expected %d", result, expected)
	}
}
