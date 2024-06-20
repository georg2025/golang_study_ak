package main

import (
	"testing"
)

func TestGetType(t *testing.T) {
	result := getType("Привет, мир!")
	expected := "string"
	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
	result = getType(22)
	expected = "int"
	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
	result = getType(nil)
	expected = "Пустой интерфейс"
	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
	result = getType([]int{1, 2, 3})
	expected = "[]int"
	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
	result = getType(22.4)
	expected = "cant figure out"
	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}
