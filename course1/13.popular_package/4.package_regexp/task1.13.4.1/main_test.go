package main

import (
	"testing"
)

func TestIsValidEmail(t *testing.T) {
	result := isValidEmail("valid@email.com")
	expected := true
	if result != expected {
		t.Errorf("Wrong result for valid@email.com")
	}
	result = isValidEmail("invalid.email.com")
	expected = false
	if result != expected {
		t.Errorf("Wrong result for invalid.email.com")
	}
	result = isValidEmail("valid@email.gov.com")
	expected = true
	if result != expected {
		t.Errorf("Wrong result for valid@email.gov.com")
	}
}
