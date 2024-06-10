package main

import (
	"reflect"
	"testing"
)

func TestAppendInt(t *testing.T) {
	result := appendInt([]int{5, 2}, 2, 5)
	expected := []int{5, 2, 2, 5}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Result is %v, expected %v", result, expected)
	}
}
