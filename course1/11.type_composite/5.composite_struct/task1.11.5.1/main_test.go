package main

import (
	"testing"
)

func TestPreparePrint(t *testing.T) {
	result := preparePrint([]User{{Name: "Egor", Age: 22}})
	expected := "Имя: Egor, Возраст: 22\n"
	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}
