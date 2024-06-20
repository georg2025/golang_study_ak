package main

import (
	"strings"
)

func main() {
}

func CountWordsIntext(txt string, words []string) map[string]int {
	wordsCountMap := make(map[string]int)
	wordsSlice := strings.Fields(txt)

	for _, i := range wordsSlice {
		i = strings.Trim(i, ",")
		i = strings.Trim(i, ".")
		i = strings.Trim(i, "!")
		wordsCountMap[strings.ToLower(i)]++
	}

	answer := make(map[string]int)
	for _, i := range words {
		answer[strings.ToLower(i)] = wordsCountMap[strings.ToLower(i)]
	}

	return answer
}
