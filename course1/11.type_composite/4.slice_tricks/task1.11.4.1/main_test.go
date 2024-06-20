package main

import (
	"reflect"
	"testing"
)

type TestData struct {
	start    int
	end      int
	inSlice  []int
	outSlice []int
}

func TestCut(t *testing.T) {
	testCases := []TestData{
		{inSlice: []int{}, outSlice: []int{}, start: 4, end: 6},
		{inSlice: []int{1, 2, 3}, outSlice: []int{}, start: 0, end: 6},
		{inSlice: []int{1, 2, 3}, outSlice: []int{}, start: -2, end: 2},
		{inSlice: []int{1, 2, 3}, outSlice: []int{}, start: 4, end: 2},
		{inSlice: []int{1, 2, 3, 4, 5, 6, 7, 8}, outSlice: []int{3, 4, 5, 6, 7, 8}, start: 2, end: 7},
	}
	for _, tc := range testCases {
		result := Cut(tc.inSlice, tc.start, tc.end)
		if !reflect.DeepEqual(result, tc.outSlice) {
			t.Errorf("Got %v, expected %v", result, tc.outSlice)
		}
	}
}
