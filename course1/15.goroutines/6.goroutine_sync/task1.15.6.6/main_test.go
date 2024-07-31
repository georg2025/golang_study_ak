package main

import (
	"sync"
	"testing"
)

func TestMain(t *testing.T) {
	var data sync.Map
	cache := Cache{data: data}
	cache.Set("Ivan", 44)
	result, ok := cache.Get("Ivan")

	if !ok {
		t.Errorf("Didnt find info")
	}

	if result != 44 {
		t.Errorf("Got %d, expected 44", result)
	}
}
