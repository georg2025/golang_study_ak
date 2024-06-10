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

func TestPop(t *testing.T) {
	testCases := []TestData{
		{inSlice: []int{}, outSlice: []int{}, outInt: 0},
		{inSlice: []int{1}, outSlice: []int{}, outInt: 1},
		{inSlice: []int{1, 2, 3, 4, 5}, outSlice: []int{2, 3, 4, 5}, outInt: 1},
		{inSlice: []int{11, 2, 1, 4}, outSlice: []int{2, 1, 4}, outInt: 11},
	}
	for _, tc := range testCases {
		number, result := Pop(tc.inSlice)
		if !reflect.DeepEqual(result, tc.outSlice) {
			t.Errorf("Got %v, expected %v", result, tc.outSlice)
		}
		if number != tc.outInt {
			t.Errorf("Got %d, expected %d", number, tc.outInt)
		}
	}
}
