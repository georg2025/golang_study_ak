package main

import (
	"reflect"
	"testing"
)

type TestData struct {
	args     []int
	inSlice  []int
	outSlice []int
}

func TestInsertToStart(t *testing.T) {
	testCases := []TestData{
		{inSlice: []int{}, outSlice: []int{}, args: []int{}},
		{inSlice: []int{1, 2, 3}, outSlice: []int{1, 2, 3}, args: []int{}},
		{inSlice: []int{1, 2, 3}, outSlice: []int{5, 6, 1, 2, 3}, args: []int{5, 6}},
	}
	for _, tc := range testCases {
		result := InsertToStart(tc.inSlice, tc.args...)
		if !reflect.DeepEqual(result, tc.outSlice) {
			t.Errorf("Got %v, expected %v", result, tc.outSlice)
		}
	}
}
