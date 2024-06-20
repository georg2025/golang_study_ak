package main

import (
	"reflect"
	"testing"
)

func TestAverage(t *testing.T) {
	testData1 := generateSlice(6)
	testData2 := generateSlice(6)
	if reflect.DeepEqual(testData1, testData2) {
		result := average(testData1)
		if result != 3.5 {
			t.Errorf("Expected 3.5, got %v", result)
		}
	}
	testData3 := []float64{}
	result := average(testData3)
	if result != 0.0 {
		t.Errorf("Expected 0.0, got %v", result)
	}
}

func generateSlice(size int) []float64 {
	slice := make([]float64, size)
	addNumber := 1.0
	for i := 0; i < size; i++ {
		slice[i] = addNumber
		addNumber += 1.0
	}
	return slice
}
