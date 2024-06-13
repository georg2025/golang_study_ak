package main

import (
	"testing"
)

func TestCountWordsInText(t *testing.T) {
	result := CountWordsIntext("Hello, hello, my dear My!", []string{"hello", "my"})
	if result["hello"] != 2 {
		t.Errorf("Got %d hello, expected 2", result["hello"])
	}
	if result["my"] != 2 {
		t.Errorf("Got %d my, expected 2", result["my"])
	}
}
