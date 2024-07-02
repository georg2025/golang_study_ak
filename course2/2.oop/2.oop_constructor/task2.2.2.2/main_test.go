package main

import (
	"testing"
)

func TestCreatingUserWithAllOptions(t *testing.T) {
	user := NewUser(1, WithUsername("testuser"), WithEmail("testuser@example.com"), WithRole("admin"))

	if user.ID != 1 {
		t.Errorf("expected ID to be 1, got %d", user.ID)
	}
	if user.Username != "testuser" {
		t.Errorf("expected Username to be 'testuser', got %s", user.Username)
	}
	if user.Email != "testuser@example.com" {
		t.Errorf("expected Email to be 'testuser@example.com', got %s", user.Email)
	}
	if user.Role != "admin" {
		t.Errorf("expected Role to be 'admin', got %s", user.Role)
	}
}

func TestCreatingUserWithEmptyUsername(t *testing.T) {
	user := NewUser(2, WithUsername(""), WithEmail("emptyuser@example.com"), WithRole("user"))

	if user.ID != 2 {
		t.Errorf("expected ID to be 2, got %d", user.ID)
	}
	if user.Username != "" {
		t.Errorf("expected Username to be empty, got %s", user.Username)
	}
	if user.Email != "emptyuser@example.com" {
		t.Errorf("expected Email to be 'emptyuser@example.com', got %s", user.Email)
	}
	if user.Role != "user" {
		t.Errorf("expected Role to be 'user', got %s", user.Role)
	}
}
