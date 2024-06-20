package main

import (
	"bytes"
	"os"
	"testing"
)

func TestMergeMaps(t *testing.T) {
	map1 := map[string]int{"gold": 3, "bronze": 2}
	map2 := map[string]int{"silver": 6, "copper": 11}
	result := mergeMaps(map1, map2)
	if result["gold"] != 3 {
		t.Errorf("Expected 3, got %d", result["gold"])
	}
	if result["copper"] != 11 {
		t.Errorf("Expected 11, got %d", result["copper"])
	}
}

func TestMainFunc(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = old

	var stdout bytes.Buffer
	stdout.ReadFrom(r)

	expected := "apple: 3\nbanana: 2\norange: 5\ngrape: 4\n"
	if stdout.String() != expected {
		t.Errorf("got %s, want %s", stdout.String(), expected)
	}
}
