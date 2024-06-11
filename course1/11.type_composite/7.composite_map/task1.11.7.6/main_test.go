package main

import (
	"testing"
)

func TestFilterWords(t *testing.T) {
	censorMap := map[string]string{
		"foo": "moo",
		"poo": "doo",
	}
	result := filterWords("Foo koo zoo! Poo foo hoo!", censorMap)
	expected := "Moo koo zoo! Doo moo hoo! "
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
	result = filterWords("Foo!", censorMap)
	expected = "Moo! "
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
