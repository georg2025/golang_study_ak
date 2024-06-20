package main

import (
	"fmt"
	"strings"
)

func main() {
}

func filterSentence(sentence string, filter map[string]bool) string {
	words := strings.Split(sentence, " ")
	var sb strings.Builder

	for _, word := range words {
		_, ok := filter[word]
		if !ok {
			sb.WriteString(word)
			sb.WriteString(" ")
		}
	}
	answer := strings.Trim(sb.String(), " ")
	fmt.Println(answer)
	return answer
}
