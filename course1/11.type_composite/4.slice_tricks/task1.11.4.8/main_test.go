package main

import (
	"reflect"
	"testing"
)

type TestData struct {
	outInt   int
	inSlice  []int
	outSlice []int
}

func TestShift(t *testing.T) {
	testCases := []TestData{
		{inSlice: []int{}, outSlice: []int{}, outInt: 0},
		{inSlice: []int{3}, outSlice: []int{3}, outInt: 3},
		{inSlice: []int{1, 2, 3, 4, 5}, outSlice: []int{5, 1, 2, 3, 4}, outInt: 1},
		{inSlice: []int{11, 2, 1, 4}, outSlice: []int{4, 11, 2, 1}, outInt: 11},
	}
	for _, tc := range testCases {
		number, result := Shift(tc.inSlice)
		if !reflect.DeepEqual(result, tc.outSlice) {
			t.Errorf("Got %v, expected %v", result, tc.outSlice)
		}
		if number != tc.outInt {
			t.Errorf("Got %d, expected %d", number, tc.outInt)
		}
	}
}
