package main

import "strings"

func main() {
}

func createUniqueText(text string) string {
	words := strings.Split(text, " ")
	helpNumber := 1
	wordCounter := make(map[string]int)
	for _, j := range words {
		_, ok := wordCounter[j]
		if !ok {
			wordCounter[j] = helpNumber
			helpNumber += 1
		}
	}
	words = make([]string, helpNumber)
	for key, value := range wordCounter {
		words[value-1] = key
	}
	var sb strings.Builder
	for _, word := range words {
		sb.WriteString(word)
		sb.WriteString(" ")
	}
	return strings.Trim(sb.String(), " ")
}
