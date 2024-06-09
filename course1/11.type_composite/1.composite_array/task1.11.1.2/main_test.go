package main

import "testing"

func TestSortDescInt(t *testing.T) {
	result := sortDescInt([8]int{1, 2, 6, 5, 1, 99, 7, 11})
	if result != [8]int{99, 11, 7, 6, 5, 2, 1, 1} {
		t.Errorf("Wrong result in sortDescInt function")
	}
}

func TestSortAscInt(t *testing.T) {
	result := sortAscInt([8]int{1, 2, 6, 5, 1, 99, 7, 11})
	if result != [8]int{1, 1, 2, 5, 6, 7, 11, 99} {
		t.Errorf("Wrong result in sortAscInt function")
	}
}

func TestSortDescFloat(t *testing.T) {
	result := sortDescFloat([8]float64{1.0, 2.0, 6.0, 3.0, 12.0, 31.0, 21.0, 77.0})
	if result != [8]float64{77.0, 31.0, 21.0, 12.0, 6.0, 3.0, 2.0, 1.0} {
		t.Errorf("Wrong result in sortDescFloat function")
	}
}

func TestSortAscFloat(t *testing.T) {
	result := sortAscFloat([8]float64{1.0, 2.0, 6.0, 3.0, 12.0, 31.0, 21.0, 77.0})
	if result != [8]float64{1.0, 2.0, 3.0, 6.0, 12.0, 21.0, 31.0, 77.0} {
		t.Errorf("Wrong result in sortAscFloat function")
	}
}
