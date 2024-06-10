package main

import (
	"testing"
)

func TestInsertAfterIDX(t *testing.T) {
	testSlice := make([]int, 3, 6)
	result := cap(RemoveExtraMemory(testSlice))
	if result != 3 {
		t.Errorf("Got %d, expected 3", result)
	}
	testSlice = make([]int, 3)
	result = cap(RemoveExtraMemory(testSlice))
	if result != 3 {
		t.Errorf("Got %d, expected 3", result)
	}
}
