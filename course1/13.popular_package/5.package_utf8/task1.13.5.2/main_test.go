package main

import (
	"testing"
)

func TestCountRussianLetters(t *testing.T) {
	result := countRussianLetters("ВввввАААаДДд")
	if result['в'] != 5 {
		t.Errorf("Got %d, expected 5", result['в'])
	}
	if result['а'] != 4 {
		t.Errorf("Got %d, expected 4", result['а'])
	}
	if result['д'] != 3 {
		t.Errorf("Got %d, expected 3", result['д'])
	}
}
