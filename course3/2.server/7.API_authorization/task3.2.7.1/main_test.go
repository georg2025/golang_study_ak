package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	byteString := []byte("Hello world")
	firstHash := hashPassword(byteString)
	secondHash := hashPassword(byteString)

	if firstHash != secondHash {
		t.Errorf("Should return similar hash on similar input")
	}

	newByteString := []byte("world Hello")
	thirdHash := hashPassword(newByteString)

	if thirdHash == firstHash {
		t.Errorf("Should return different hash on different input")
	}
}
