package main

import (
	"testing"
)

func TestReplaceSymbols(t *testing.T) {
	result := ReplaceSymbols("Hello World", 'l', 'z')
	expected := "Hezzo Worzd"
	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
	result = ReplaceSymbols("", 'l', 'z')
	expected = ""
	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}
