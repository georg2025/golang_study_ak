package main

import (
	"unicode/utf8"
)

func main() {
}

func countUniqueUTF8Chars(s string) int {
	if len(s) == 0 {
		return 0
	}
	uniqueRunes := make(map[rune]struct{})
	for len(s) > 0 {
		r, size := utf8.DecodeRuneInString(s)
		_, ok := uniqueRunes[r]
		if !ok {
			uniqueRunes[r] = struct{}{}
		}
		s = s[size:]
	}
	return len(uniqueRunes)
}
