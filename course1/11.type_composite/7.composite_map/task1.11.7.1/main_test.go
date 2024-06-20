package main

import (
	"testing"
)

func TestCountWordOccurrence(t *testing.T) {
	result := countWordOccurrences("hello we are we so hello")
	if result["hello"] != 2 {
		t.Errorf("Expected 2, got %d", result["hello"])
	}
	if result["are"] != 1 {
		t.Errorf("Expected 1, got %d", result["are"])
	}
}
