package main

import (
	"strings"
)

func main() {
}

func ReplaceSymbols(text string, badchar, goodchar rune) string {
	if len(text) == 0 {
		return ""
	}
	var sb strings.Builder
	for _, ch := range text {
		if ch == badchar {
			ch = goodchar
		}
		sb.WriteRune(ch)
	}
	return sb.String()
}
