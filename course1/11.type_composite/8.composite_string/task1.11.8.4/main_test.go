package main

import (
	"testing"
)

func TestConcatStrings(t *testing.T) {
	result := concatStrings("Hello", " ", "World")
	expected := "Hello World"
	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
	result = concatStrings()
	expected = ""
	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
	result = concatStrings("Hello")
	expected = "Hello"
	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}
