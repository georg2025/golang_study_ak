package main

import (
	"unicode"
	"unicode/utf8"
)

func main() {
}

func countRussianLetters(s string) map[rune]int {
	russianLetters := []rune("абвгдеёжзийклмнопрстуфхцчшщъыьэюя")
	lettersMap := make(map[rune]int)
	for _, i := range russianLetters {
		lettersMap[i] = 0
	}
	for len(s) > 0 {
		r, size := utf8.DecodeRuneInString(s)
		_, ok := lettersMap[unicode.ToLower(r)]
		if ok {
			lettersMap[unicode.ToLower(r)]++
		}
		s = s[size:]
	}
	return lettersMap
}
