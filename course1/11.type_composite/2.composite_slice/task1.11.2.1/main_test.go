package main

import (
	"reflect"
	"testing"
)

func TestGetSubSlice(t *testing.T) {
	result := getSubSlice([]int{5, 2, 6, 8, 4, 8, 0, 4, 3}, 2, 5)
	expected := []int{6, 8, 4}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Result is %v, expected %v", result, expected)
	}
}
