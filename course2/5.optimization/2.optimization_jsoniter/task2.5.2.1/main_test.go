package main

import (
	"testing"
)

func BenchJson(b *testing.B) {
	users := GenerateUSER(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		JSON(users)
	}
}

func BenchJsonIter(b *testing.B) {
	users := GenerateUSER(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		JSONiter(users)
	}
}
