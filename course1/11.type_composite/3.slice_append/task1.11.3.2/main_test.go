package main

import (
	"reflect"
	"testing"
)

func TestAppendInt(t *testing.T) {
	initSlice := []int{5, 2}
	appendInt(&initSlice, 2, 5)
	expected := []int{5, 2, 2, 5}
	if !reflect.DeepEqual(initSlice, expected) {
		t.Errorf("Result is %v, expected %v", initSlice, expected)
	}
}
