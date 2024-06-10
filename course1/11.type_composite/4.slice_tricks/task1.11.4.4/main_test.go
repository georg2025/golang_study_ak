package main

import (
	"reflect"
	"testing"
)

type TestData struct {
	idx      int
	inSlice  []int
	outSlice []int
}

func TestRemoveIDX(t *testing.T) {
	testCases := []TestData{
		{inSlice: []int{}, outSlice: []int{}, idx: 2},
		{inSlice: []int{1, 2, 3}, outSlice: []int{1, 2, 3}, idx: 5},
		{inSlice: []int{1, 2, 3}, outSlice: []int{1, 2, 3}, idx: -1},
		{inSlice: []int{1, 2, 3}, outSlice: []int{2, 3}, idx: 0},
		{inSlice: []int{1, 2, 3, 4, 5}, outSlice: []int{1, 2, 4, 5}, idx: 2},
		{inSlice: []int{1, 2, 3, 4, 5}, outSlice: []int{1, 2, 3, 4}, idx: 4},
	}
	for _, tc := range testCases {
		result := RemoveIDX(tc.inSlice, tc.idx)
		if !reflect.DeepEqual(result, tc.outSlice) {
			t.Errorf("Got %v, expected %v", result, tc.outSlice)
		}
	}
}
