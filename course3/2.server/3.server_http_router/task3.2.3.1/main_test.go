package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRouterHandlesEmptyOrMalformedRequests(t *testing.T) {
	router := makeRouter()

	req, _ := http.NewRequest("GET", "/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code, "Should return 200 (ok)")
	assert.Equal(t, "Hello world", resp.Body.String(), "Answer should be Hello world")

	req, _ = http.NewRequest("GET", "/2", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code, "Should return 200 (ok)")
	assert.Equal(t, "Hello world 2", resp.Body.String(), "Answer should be Hello world 2")

	req, _ = http.NewRequest("GET", "/3", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code, "Should return 200 (ok)")
	assert.Equal(t, "Hello world 3", resp.Body.String(), "Answer should be Hello world 3")

	req, _ = http.NewRequest("GET", "/", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusNotFound, resp.Code, "Should return 404 for unknown routes")
}
