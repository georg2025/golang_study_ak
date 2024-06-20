package main

import (
	"testing"
)

func TestCountUniqueUTF8Chars(t *testing.T) {
	result := countUniqueUTF8Chars("Oh my god oh oh")
	expected := 8
	if result != expected {
		t.Errorf("Got %d, expected %d", result, expected)
	}
	result = countUniqueUTF8Chars("")
	expected = 0
	if result != expected {
		t.Errorf("Got %d, expected %d", result, expected)
	}
	result = countUniqueUTF8Chars("Ничего не знаю")
	expected = 11
	if result != expected {
		t.Errorf("Got %d, expected %d", result, expected)
	}
}
