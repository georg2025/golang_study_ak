package main

import (
	"fmt"
	"strings"
)

func main() {
}

func countWordOccurrences(text string) map[string]int {
	words := strings.Split(text, " ")
	answerMap := make(map[string]int)
	for _, i := range words {
		answerMap[i]++
	}
	for i, j := range answerMap {
		fmt.Printf("Word %s: %d\n", i, j)
	}
	return answerMap
}
