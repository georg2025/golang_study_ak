package main

import (
	"testing"
)

func TestCreateUniqueText(t *testing.T) {
	result := createUniqueText("bar bar bar foo foo baz")
	expected := "bar foo baz"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
