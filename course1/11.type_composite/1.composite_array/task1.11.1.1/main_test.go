package main

import "testing"

func TestSum(t *testing.T) {
	result := sum([8]int{1, 2, 3, 2, 1, 3, 2, 1})
	if result != 15 {
		t.Errorf("Result is %v, expected 15", result)
	}
}

func TestAverage(t *testing.T) {
	result := average([8]int{1, 2, 3, 2, 1, 3, 2, 1})
	if result != 1.875 {
		t.Errorf("Result is %v, expected 1.875", result)
	}
}

func TestAverageFloat(t *testing.T) {
	result := averageFloat([8]float64{1.0, 2.0, 3.0, 2.0, 1.0, 3.0, 2.0, 1.0})
	if result != 1.875 {
		t.Errorf("Result is %v, expected 1.875", result)
	}
}

func TestReverse(t *testing.T) {
	result := reverse([8]int{1, 2, 3, 2, 1, 3, 2, 1})
	if result != [8]int{1, 2, 3, 1, 2, 3, 2, 1} {
		t.Errorf("Result is wrong, check function")
	}
}
