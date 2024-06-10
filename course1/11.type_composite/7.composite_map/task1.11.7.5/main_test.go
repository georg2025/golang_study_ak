package main

import (
	"testing"
)

func TestFilterSentence(t *testing.T) {
	filter := make(map[string]bool)
	filter["foo"] = true
	filter["hoo"] = true
	result := filterSentence("foo moo poo hoo soo hoo moo", filter)
	expected := "moo poo soo moo"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
