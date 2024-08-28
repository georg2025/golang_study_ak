package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnswer(t *testing.T) {
	req := httptest.NewRequest("GET", "localhost:0000", nil)
	w := httptest.NewRecorder()
	answer(w, req)

	resp := w.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Should return 200 (ok)")

	expected := "Hello from API"
	result := w.Body.String()
	assert.Equal(t, expected, result, "Wrong answer")
}

func TestShouldProxy(t *testing.T) {
	req, err := http.NewRequest("GET", "/somepath", nil)

	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	result := shouldProxy(req)
	if !result {
		t.Errorf("Expected true, got %v", result)
	}
}
