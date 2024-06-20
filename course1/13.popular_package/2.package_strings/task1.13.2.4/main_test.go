package main

import (
	"testing"
)

func TestGenerateActivationKey(t *testing.T) {
	result := GenerateActivationKey(255)
	expected := "4qsw-4qsw-4qsw-4qsw"
	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}
