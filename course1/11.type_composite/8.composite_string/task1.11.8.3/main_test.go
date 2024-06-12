package main

import (
	"reflect"
	"testing"
)

func TestGetBytes(t *testing.T) {
	result := getBytes("Hi")
	expected := []byte{72, 105}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Got %v, expected %v", result, expected)
	}
}

func TestGetRunes(t *testing.T) {
	result := getRunes("Hi")
	expected := []rune{'H', 'i'}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Got %v, expected %v", result, expected)
	}
}
