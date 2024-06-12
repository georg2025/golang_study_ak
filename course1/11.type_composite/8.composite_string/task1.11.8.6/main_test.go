package main

import (
	"testing"
)

func TestCountVowels(t *testing.T) {
	result := CountVowels("Hello worldie")
	expected := 5
	if result != expected {
		t.Errorf("Got %d, expected %d", result, expected)
	}
	result = CountVowels("")
	expected = 0
	if result != expected {
		t.Errorf("Got %d, expected %d", result, expected)
	}
	result = CountVowels("Привееет мир")
	expected = 5
	if result != expected {
		t.Errorf("Got %d, expected %d", result, expected)
	}
}
