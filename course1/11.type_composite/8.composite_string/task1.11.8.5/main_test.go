package main

import (
	"testing"
)

func TestReverseString(t *testing.T) {
	result := ReverseString("Hello")
	expected := "olleH"
	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
	result = ReverseString("H")
	expected = "H"
	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
	result = ReverseString("")
	expected = ""
	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}
