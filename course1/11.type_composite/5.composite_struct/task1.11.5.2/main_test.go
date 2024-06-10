package main

import (
	"testing"
)

func TestPreparePrint(t *testing.T) {
	result := preparePrint([]Animal{{Type: "cat", Name: "Murka", Age: 6}})
	expected := "Тип: cat, Имя: Murka, Возраст: 6\n"
	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}
