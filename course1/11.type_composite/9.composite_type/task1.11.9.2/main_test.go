package main

import (
	"testing"
)

func TestMainFunction(t *testing.T) {
	Samsunger := SamsungTV{
		status: false,
		model:  "Samsunger",
	}
	LGer := LgTV{
		status: false,
		model:  "LGer",
	}
	result := Samsunger.GetStatus()
	expected := false
	if result != expected {
		t.Errorf("Got %t, expected %t", result, expected)
	}
	result = LGer.GetStatus()
	expected = false
	if result != expected {
		t.Errorf("Got %t, expected %t", result, expected)
	}
	err := Samsunger.switchOn()
	if err != nil {
		t.Errorf("Error with switching on: %e", err)
	}
	err = LGer.switchOn()
	if err != nil {
		t.Errorf("Error with switching on: %e", err)
	}
	result = Samsunger.GetStatus()
	expected = true
	if result != expected {
		t.Errorf("Got %t, expected %t", result, expected)
	}
	result = LGer.GetStatus()
	expected = true
	if result != expected {
		t.Errorf("Got %t, expected %t", result, expected)
	}
	resultModel := Samsunger.GetModel()
	expectedModel := "Samsunger"
	if result != expected {
		t.Errorf("Got %s, expected %s", resultModel, expectedModel)
	}
	resultModel = LGer.GetModel()
	expectedModel = "Lger"
	if result != expected {
		t.Errorf("Got %s, expected %s", resultModel, expectedModel)
	}
	err = Samsunger.switchOn()
	if err == nil {
		t.Errorf("This should trigger error")
	}
	err = Samsunger.switchOn()
	if err == nil {
		t.Errorf("This should trigger error")
	}
	err = Samsunger.switchOFF()
	if err != nil {
		t.Errorf("Error with switching on: %e", err)
	}
	err = LGer.switchOFF()
	if err != nil {
		t.Errorf("Error with switching on: %e", err)
	}
	err = Samsunger.switchOFF()
	if err == nil {
		t.Errorf("This should trigger error")
	}
	err = Samsunger.switchOFF()
	if err == nil {
		t.Errorf("This should trigger error")
	}
}
