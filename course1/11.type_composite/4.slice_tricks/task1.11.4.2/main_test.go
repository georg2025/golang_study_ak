package main

import (
	"reflect"
	"testing"
)

type TestData struct {
	idx        int
	addedElems []int
	inSlice    []int
	outSlice   []int
}

func TestInsertAfterIDX(t *testing.T) {
	testCases := []TestData{
		{inSlice: []int{}, outSlice: []int{}, idx: 0, addedElems: []int{3, 2, 5}},
		{inSlice: []int{1, 2, 3}, outSlice: []int{}, idx: -2, addedElems: []int{3, 2, 5}},
		{inSlice: []int{1, 2, 3}, outSlice: []int{1, 2, 3, 4, 5}, idx: 2, addedElems: []int{4, 5}},
		{inSlice: []int{1, 2, 3}, outSlice: []int{1, 2, 4, 5}, idx: 1, addedElems: []int{4, 5}},
	}
	for _, tc := range testCases {
		result := InsertAfterIDX(tc.inSlice, tc.idx, tc.addedElems...)
		if !reflect.DeepEqual(result, tc.outSlice) {
			t.Errorf("Got %v, expected %v", result, tc.outSlice)
		}
	}
}
