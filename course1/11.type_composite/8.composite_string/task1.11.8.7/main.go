package main

import (
	"strings"
)

func main() {
}

func ReplaceSymbols(text string, oldChar, newChar rune) string {
	if len(text) == 0 {
		return ""
	}
	var sb strings.Builder
	for _, ch := range text {
		if ch == oldChar {
			ch = newChar
		}
		sb.WriteRune(ch)
	}
	return sb.String()
}
