package main

import (
	"reflect"
	"testing"
)

func TestFilterDividers(t *testing.T) {
	result := FilterDividers([]int{}, 16)
	expected := []int{}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Got %v, expected %v", result, expected)
	}
	result = FilterDividers([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3)
	expected = []int{3, 6, 9}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Got %v, expected %v", result, expected)
	}
}
