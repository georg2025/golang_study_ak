package main

import (
	"strings"
)

func main() {
}

func CountWordsIntext(txt string, words []string) map[string]int {
	helpMap := make(map[string]int)
	wordsSlice := strings.Fields(txt)
	for _, i := range wordsSlice {
		i = strings.Trim(i, ",")
		i = strings.Trim(i, ".")
		i = strings.Trim(i, "!")
		helpMap[strings.ToLower(i)]++
	}
	answer := make(map[string]int)
	for _, i := range words {
		answer[strings.ToLower(i)] = helpMap[strings.ToLower(i)]
	}
	return answer
}
