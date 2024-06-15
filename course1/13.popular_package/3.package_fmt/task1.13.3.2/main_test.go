package main

import (
	"testing"
)

func TestGetVariableType(t *testing.T) {
	result := getVariableType("Привет, мир!")
	expected := "Variable type is: string"
	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
	result = getVariableType(22)
	expected = "Variable type is: int"
	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
	result = getVariableType(nil)
	expected = "Пустой интерфейс"
	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
	result = getVariableType([]int{1, 2, 3})
	expected = "Variable type is: []int"
	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
	result = getVariableType(22.4)
	expected = "Cant figure out variable type"
	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}
