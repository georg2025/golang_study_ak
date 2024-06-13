package main

import (
	"testing"
)

func TestOperate(t *testing.T) {
	result := Operate(Concat, "Hello, ", "world!")
	expected := "Hello, world!"
	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
	resultSum := Operate(Sum, 1, 2, 3)
	expectedSum := 6
	if resultSum != expectedSum {
		t.Errorf("Got %d, expected %d", resultSum, expectedSum)
	}
}
