package main

import (
	"unicode"
)

func main() {
}

func CountVowels(str string) int {
	if len(str) == 0 {
		return 0
	}
	vowelCount := 0
	for _, i := range str {
		if isVowel(i) {
			vowelCount++
		}
	}
	return vowelCount
}

func isVowel(ch rune) bool {
	ch = unicode.ToLower(ch)
	switch ch {
	case 'a', 'e', 'i', 'o', 'u', 'а', 'у', 'о', 'и', 'э', 'ы', 'я', 'ю', 'е', 'ё':
		return true
	}
	return false
}
